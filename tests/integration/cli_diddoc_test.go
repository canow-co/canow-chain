//go:build integration

package integration

import (
	"crypto/ed25519"
	"fmt"

	"github.com/canow-co/canow-chain/tests/integration/cli"
	"github.com/canow-co/canow-chain/tests/integration/helpers"
	"github.com/canow-co/canow-chain/tests/integration/network"
	"github.com/canow-co/canow-chain/tests/integration/testdata"
	didcli "github.com/canow-co/cheqd-node/x/did/client/cli"
	testsetup "github.com/canow-co/cheqd-node/x/did/tests/setup"
	"github.com/canow-co/cheqd-node/x/did/types"
	"github.com/google/uuid"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cheqd cli - positive did", func() {
	var tmpDir string
	var feeParams types.FeeParams

	BeforeEach(func() {
		tmpDir = GinkgoT().TempDir()

		// Query fee params
		res, err := cli.QueryParams(types.ModuleName, string(types.ParamStoreKeyFeeParams))
		Expect(err).To(BeNil())
		err = helpers.Codec.UnmarshalJSON([]byte(res.Value), &feeParams)
		Expect(err).To(BeNil())
	})

	It("can create and update diddoc using signature by method referenced from Authentication verification relationship and query the result (Ed25519VerificationKey2020)", func() {
		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can create diddoc (Ed25519VerificationKey2020)"))
		// Create a new DID Doc
		did := "did:canow:" + network.DidNamespace + ":" + uuid.NewString()
		keyID := did + "#key1"

		publicKey, privateKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		publicKeyMultibase := testsetup.GenerateEd25519VerificationKey2020VerificationMaterial(publicKey)

		payload := didcli.DIDDocument{
			ID: did,
			VerificationMethod: []didcli.VerificationMethod{
				map[string]any{
					"id":                 keyID,
					"type":               "Ed25519VerificationKey2020",
					"controller":         did,
					"publicKeyMultibase": publicKeyMultibase,
				},
			},
			Authentication: []any{keyID},
		}

		signInputs := []didcli.SignInput{
			{
				VerificationMethodID: keyID,
				PrivKey:              privateKey,
			},
		}

		versionID := uuid.NewString()

		res, err := cli.CreateDidDoc(tmpDir, payload, signInputs, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.CreateDid.String()))
		Expect(err).To(BeNil())
		Expect(res.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can update diddoc (Ed25519VerificationKey2020)"))
		// Update the DID Doc
		newPublicKey, newPrivateKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		newPublicKeyMultibase := testsetup.GenerateEd25519VerificationKey2020VerificationMaterial(newPublicKey)

		payload2 := didcli.DIDDocument{
			ID: did,
			VerificationMethod: []didcli.VerificationMethod{
				map[string]any{
					"id":                 keyID,
					"type":               "Ed25519VerificationKey2020",
					"controller":         did,
					"publicKeyMultibase": newPublicKeyMultibase,
				},
			},
			Authentication: []any{keyID},
		}

		signInputs2 := []didcli.SignInput{
			{
				VerificationMethodID: keyID,
				PrivKey:              privateKey,
			},
			{
				VerificationMethodID: keyID,
				PrivKey:              newPrivateKey,
			},
		}

		versionID = uuid.NewString()

		res2, err := cli.UpdateDidDoc(tmpDir, payload2, signInputs2, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.UpdateDid.String()))
		Expect(err).To(BeNil())
		Expect(res2.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query diddoc (Ed25519VerificationKey2020)"))
		// Query the DID Doc
		resp, err := cli.QueryDidDoc(did)
		Expect(err).To(BeNil())

		didDoc := resp.Value.DidDoc
		Expect(didDoc.Id).To(BeEquivalentTo(did))
		Expect(didDoc.Authentication).To(HaveLen(1))
		Expect(didDoc.Authentication[0].VerificationMethodId).To(BeEquivalentTo(keyID))
		Expect(didDoc.VerificationMethod).To(HaveLen(1))
		Expect(didDoc.VerificationMethod[0].Id).To(BeEquivalentTo(keyID))
		Expect(didDoc.VerificationMethod[0].VerificationMethodType).To(BeEquivalentTo("Ed25519VerificationKey2020"))
		Expect(didDoc.VerificationMethod[0].Controller).To(BeEquivalentTo(did))
		Expect(didDoc.VerificationMethod[0].VerificationMaterial).To(BeEquivalentTo(newPublicKeyMultibase))

		// Check that DIDDoc is not deactivated
		Expect(resp.Value.Metadata.Deactivated).To(BeFalse())

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can deactivate diddoc (Ed25519VerificationKey2020)"))
		// Deactivate the DID Doc
		payload3 := types.MsgDeactivateDidDocPayload{
			Id: did,
		}

		signInputs3 := []didcli.SignInput{
			{
				VerificationMethodID: keyID,
				PrivKey:              newPrivateKey,
			},
		}

		versionID = uuid.NewString()

		res3, err := cli.DeactivateDidDoc(tmpDir, payload3, signInputs3, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.DeactivateDid.String()))
		Expect(err).To(BeNil())
		Expect(res3.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query deactivated diddoc (Ed25519VerificationKey2020)"))
		// Query the DID Doc

		resp2, err := cli.QueryDidDoc(did)
		Expect(err).To(BeNil())

		didDoc2 := resp2.Value.DidDoc
		Expect(didDoc2).To(BeEquivalentTo(didDoc))

		// Check that the DID Doc is deactivated
		Expect(resp2.Value.Metadata.Deactivated).To(BeTrue())
	})

	It("can create diddoc, update it and query the result (JsonWebKey2020)", func() {
		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can create diddoc (JsonWebKey2020)"))
		// Create a new DID Doc
		did := "did:canow:" + network.DidNamespace + ":" + uuid.NewString()
		keyID := did + "#key1"

		publicKey, privateKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		publicKeyJwkJSON := testsetup.GenerateJSONWebKey2020VerificationMaterial(publicKey)
		publicKeyJwk, err := testsetup.ParseJSONToMap(publicKeyJwkJSON)
		Expect(err).To(BeNil())

		payload := didcli.DIDDocument{
			ID: did,
			VerificationMethod: []didcli.VerificationMethod{
				map[string]any{
					"id":           keyID,
					"type":         "JsonWebKey2020",
					"controller":   did,
					"publicKeyJwk": publicKeyJwk,
				},
			},
			Authentication: []any{keyID},
		}

		signInputs := []didcli.SignInput{
			{
				VerificationMethodID: keyID,
				PrivKey:              privateKey,
			},
		}

		versionID := uuid.NewString()

		res, err := cli.CreateDidDoc(tmpDir, payload, signInputs, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.CreateDid.String()))
		Expect(err).To(BeNil())
		Expect(res.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can update diddoc (JsonWebKey2020)"))
		// Update the DID Doc
		newPublicKey, newPrivateKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		newPublicKeyJwkJSON := testsetup.GenerateJSONWebKey2020VerificationMaterial(newPublicKey)
		newPublicKeyJwk, err := testsetup.ParseJSONToMap(newPublicKeyJwkJSON)
		Expect(err).To(BeNil())

		payload2 := didcli.DIDDocument{
			ID: did,
			VerificationMethod: []didcli.VerificationMethod{
				map[string]any{
					"id":           keyID,
					"type":         "JsonWebKey2020",
					"controller":   did,
					"publicKeyJwk": newPublicKeyJwk,
				},
			},
			Authentication: []any{keyID},
		}

		signInputs2 := []didcli.SignInput{
			{
				VerificationMethodID: keyID,
				PrivKey:              privateKey,
			},
			{
				VerificationMethodID: keyID,
				PrivKey:              newPrivateKey,
			},
		}

		versionID = uuid.NewString()

		res2, err := cli.UpdateDidDoc(tmpDir, payload2, signInputs2, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.UpdateDid.String()))
		Expect(err).To(BeNil())
		Expect(res2.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query diddoc (JsonWebKey2020)"))
		// Query the DID Doc
		resp, err := cli.QueryDidDoc(did)
		Expect(err).To(BeNil())

		didDoc := resp.Value.DidDoc
		Expect(didDoc.Id).To(BeEquivalentTo(did))
		Expect(didDoc.Authentication).To(HaveLen(1))
		Expect(didDoc.Authentication[0].VerificationMethodId).To(BeEquivalentTo(keyID))
		Expect(didDoc.VerificationMethod).To(HaveLen(1))
		Expect(didDoc.VerificationMethod[0].Id).To(BeEquivalentTo(keyID))
		Expect(didDoc.VerificationMethod[0].VerificationMethodType).To(BeEquivalentTo("JsonWebKey2020"))
		Expect(didDoc.VerificationMethod[0].Controller).To(BeEquivalentTo(did))
		Expect(didDoc.VerificationMethod[0].VerificationMaterial).To(BeEquivalentTo(newPublicKeyJwkJSON))

		// Check that DIDDoc is not deactivated
		Expect(resp.Value.Metadata.Deactivated).To(BeFalse())

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can deactivate diddoc (JsonWebKey2020)"))
		// Deactivate the DID Doc
		payload3 := types.MsgDeactivateDidDocPayload{
			Id: did,
		}

		signInputs3 := []didcli.SignInput{
			{
				VerificationMethodID: keyID,
				PrivKey:              newPrivateKey,
			},
		}

		versionID = uuid.NewString()

		res3, err := cli.DeactivateDidDoc(tmpDir, payload3, signInputs3, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.DeactivateDid.String()))
		Expect(err).To(BeNil())
		Expect(res3.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query deactivated diddoc (JsonWebKey2020)"))
		// Query the DID Doc

		resp2, err := cli.QueryDidDoc(did)
		Expect(err).To(BeNil())

		didDoc2 := resp2.Value.DidDoc
		Expect(didDoc2).To(BeEquivalentTo(didDoc))

		// Check that the DID Doc is deactivated
		Expect(resp2.Value.Metadata.Deactivated).To(BeTrue())
	})

	It("can create diddoc, update it and query the result (Ed25519VerificationKey2018)", func() {
		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can create diddoc (Ed25519VerificationKey2018)"))
		// Create a new DID Doc
		did := "did:canow:" + network.DidNamespace + ":" + uuid.NewString()
		keyID := did + "#key1"

		publicKey, privateKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		publicKeyBase58 := testsetup.GenerateEd25519VerificationKey2018VerificationMaterial(publicKey)

		payload := didcli.DIDDocument{
			ID: did,
			VerificationMethod: []didcli.VerificationMethod{
				map[string]any{
					"id":              keyID,
					"type":            "Ed25519VerificationKey2018",
					"controller":      did,
					"publicKeyBase58": publicKeyBase58,
				},
			},
			Authentication: []any{keyID},
		}

		signInputs := []didcli.SignInput{
			{
				VerificationMethodID: keyID,
				PrivKey:              privateKey,
			},
		}

		versionID := uuid.NewString()

		res, err := cli.CreateDidDoc(tmpDir, payload, signInputs, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.CreateDid.String()))
		Expect(err).To(BeNil())
		Expect(res.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can update diddoc (Ed25519VerificationKey2018)"))
		// Update the DID Doc
		newPublicKey, newPrivateKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		newPublicKeyBase58 := testsetup.GenerateEd25519VerificationKey2018VerificationMaterial(newPublicKey)

		payload2 := didcli.DIDDocument{
			ID: did,
			VerificationMethod: []didcli.VerificationMethod{
				map[string]any{
					"id":              keyID,
					"type":            "Ed25519VerificationKey2018",
					"controller":      did,
					"publicKeyBase58": newPublicKeyBase58,
				},
			},
			Authentication: []any{keyID},
		}

		signInputs2 := []didcli.SignInput{
			{
				VerificationMethodID: keyID,
				PrivKey:              privateKey,
			},
			{
				VerificationMethodID: keyID,
				PrivKey:              newPrivateKey,
			},
		}

		versionID = uuid.NewString()

		res2, err := cli.UpdateDidDoc(tmpDir, payload2, signInputs2, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.UpdateDid.String()))
		Expect(err).To(BeNil())
		Expect(res2.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query diddoc (Ed25519VerificationKey2018)"))
		// Query the DID Doc
		resp, err := cli.QueryDidDoc(did)
		Expect(err).To(BeNil())

		didDoc := resp.Value.DidDoc
		Expect(didDoc.Id).To(BeEquivalentTo(did))
		Expect(didDoc.Authentication).To(HaveLen(1))
		Expect(didDoc.Authentication[0].VerificationMethodId).To(BeEquivalentTo(keyID))
		Expect(didDoc.Authentication[0].VerificationMethod).To(BeNil())
		Expect(didDoc.VerificationMethod).To(HaveLen(1))
		Expect(didDoc.VerificationMethod[0].Id).To(BeEquivalentTo(keyID))
		Expect(didDoc.VerificationMethod[0].VerificationMethodType).To(BeEquivalentTo("Ed25519VerificationKey2018"))
		Expect(didDoc.VerificationMethod[0].Controller).To(BeEquivalentTo(did))
		Expect(didDoc.VerificationMethod[0].VerificationMaterial).To(BeEquivalentTo(newPublicKeyBase58))
	})

	It("can create and update diddoc using signature by method embedded in Authentication verification relationship and query the result", func() {
		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can create diddoc"))
		// Create a new DID Doc
		did := "did:canow:" + network.DidNamespace + ":" + uuid.NewString()
		keyId := did + "#key1"

		pubKey, privKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		pubKeyMultibase58 := testsetup.GenerateEd25519VerificationKey2020VerificationMaterial(pubKey)
		Expect(err).To(BeNil())
		payload := didcli.DIDDocument{
			ID: did,
			Authentication: []any{
				map[string]any{
					"id":                 keyId,
					"type":               "Ed25519VerificationKey2020",
					"controller":         did,
					"publicKeyMultibase": pubKeyMultibase58,
				},
			},
		}

		signInputs := []didcli.SignInput{
			{
				VerificationMethodID: keyId,
				PrivKey:              privKey,
			},
		}

		versionID := uuid.NewString()

		res, err := cli.CreateDidDoc(tmpDir, payload, signInputs, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.CreateDid.String()))
		Expect(err).To(BeNil())
		Expect(res.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can update diddoc"))
		// Update the DID Doc
		newPubKey, newPrivKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		newPubKeyMultibase58 := testsetup.GenerateEd25519VerificationKey2020VerificationMaterial(newPubKey)
		Expect(err).To(BeNil())

		payload2 := didcli.DIDDocument{
			ID: did,
			Authentication: []any{
				map[string]any{
					"id":                 keyId,
					"type":               "Ed25519VerificationKey2020",
					"controller":         did,
					"publicKeyMultibase": newPubKeyMultibase58,
				},
			},
		}

		signInputs2 := []didcli.SignInput{
			{
				VerificationMethodID: keyId,
				PrivKey:              privKey,
			},
			{
				VerificationMethodID: keyId,
				PrivKey:              newPrivKey,
			},
		}

		res2, err := cli.UpdateDidDoc(tmpDir, payload2, signInputs2, "", testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.CreateDid.String()))
		Expect(err).To(BeNil())
		Expect(res2.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query diddoc"))
		// Query the DID Doc
		resp, err := cli.QueryDidDoc(did)
		Expect(err).To(BeNil())

		didDoc := resp.Value.DidDoc
		Expect(didDoc.Id).To(BeEquivalentTo(did))
		Expect(didDoc.Authentication).To(HaveLen(1))
		Expect(didDoc.Authentication[0].VerificationMethodId).To(BeEmpty())
		Expect(didDoc.Authentication[0].VerificationMethod).ToNot(BeNil())
		Expect(didDoc.Authentication[0].VerificationMethod.Id).To(BeEquivalentTo(keyId))
		Expect(didDoc.Authentication[0].VerificationMethod.VerificationMethodType).To(BeEquivalentTo("Ed25519VerificationKey2020"))
		Expect(didDoc.Authentication[0].VerificationMethod.Controller).To(BeEquivalentTo(did))
		Expect(didDoc.Authentication[0].VerificationMethod.VerificationMaterial).To(BeEquivalentTo(newPubKeyMultibase58))
	})

	It("can create and update diddoc using signature by method from VerificationMethod list not referenced from Authentication verification relationship", func() {
		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can create diddoc"))
		// Create a new DID Doc
		did := "did:canow:" + network.DidNamespace + ":" + uuid.NewString()
		keyId := did + "#key1"

		pubKey, privKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		pubKeyMultibase58 := testsetup.GenerateEd25519VerificationKey2020VerificationMaterial(pubKey)
		Expect(err).To(BeNil())

		payload := didcli.DIDDocument{
			ID: did,
			VerificationMethod: []didcli.VerificationMethod{
				map[string]any{
					"id":                 keyId,
					"type":               "Ed25519VerificationKey2020",
					"controller":         did,
					"publicKeyMultibase": pubKeyMultibase58,
				},
			},
		}

		signInputs := []didcli.SignInput{
			{
				VerificationMethodID: keyId,
				PrivKey:              privKey,
			},
		}

		res, err := cli.CreateDidDoc(tmpDir, payload, signInputs, "", testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.CreateDid.String()))
		Expect(err).To(BeNil())
		Expect(res.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can update diddoc"))
		// Update the DID Doc
		newPubKey, newPrivKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		newPubKeyMultibase58 := testsetup.GenerateEd25519VerificationKey2020VerificationMaterial(newPubKey)
		Expect(err).To(BeNil())

		payload2 := didcli.DIDDocument{
			ID: did,
			VerificationMethod: []didcli.VerificationMethod{
				map[string]any{
					"id":                 keyId,
					"type":               "Ed25519VerificationKey2020",
					"controller":         did,
					"publicKeyMultibase": newPubKeyMultibase58,
				},
			},
		}

		signInputs2 := []didcli.SignInput{
			{
				VerificationMethodID: keyId,
				PrivKey:              privKey,
			},
			{
				VerificationMethodID: keyId,
				PrivKey:              newPrivKey,
			},
		}

		res2, err := cli.UpdateDidDoc(tmpDir, payload2, signInputs2, "", testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.CreateDid.String()))
		Expect(err).To(BeNil())
		Expect(res2.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query diddoc"))
		// Query the DID Doc
		resp, err := cli.QueryDidDoc(did)
		Expect(err).To(BeNil())

		didDoc := resp.Value.DidDoc
		Expect(didDoc.Id).To(BeEquivalentTo(did))
		Expect(didDoc.VerificationMethod).To(HaveLen(1))
		Expect(didDoc.VerificationMethod[0].Id).To(BeEquivalentTo(keyId))
		Expect(didDoc.VerificationMethod[0].VerificationMethodType).To(BeEquivalentTo("Ed25519VerificationKey2020"))
		Expect(didDoc.VerificationMethod[0].Controller).To(BeEquivalentTo(did))
		Expect(didDoc.VerificationMethod[0].VerificationMaterial).To(BeEquivalentTo(newPubKeyMultibase58))

		// Check that DIDDoc is not deactivated
		Expect(resp.Value.Metadata.Deactivated).To(BeFalse())

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can deactivate diddoc (Ed25519VerificationKey2018)"))
		// Deactivate the DID Doc
		payload3 := types.MsgDeactivateDidDocPayload{
			Id: did,
		}

		signInputs3 := []didcli.SignInput{
			{
				VerificationMethodID: keyId,
				PrivKey:              newPrivKey,
			},
		}
		versionID := uuid.NewString()
		res3, err := cli.DeactivateDidDoc(tmpDir, payload3, signInputs3, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.DeactivateDid.String()))
		Expect(err).To(BeNil())
		Expect(res3.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query deactivated diddoc (Ed25519VerificationKey2018)"))
		// Query the DID Doc

		resp2, err := cli.QueryDidDoc(did)
		Expect(err).To(BeNil())

		didDoc2 := resp2.Value.DidDoc
		Expect(didDoc2).To(BeEquivalentTo(didDoc))

		// Check that the DID Doc is deactivated
		Expect(resp2.Value.Metadata.Deactivated).To(BeTrue())
	})

	It("can create and update diddoc with not required Service fields", func() {
		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can create diddoc"))
		// Create a new DID Doc
		did := "did:canow:" + network.DidNamespace + ":" + uuid.NewString()
		keyId := did + "#key1"

		pubKey, privKey, err := ed25519.GenerateKey(nil)
		Expect(err).To(BeNil())

		pubKeyMultibase58 := testsetup.GenerateEd25519VerificationKey2020VerificationMaterial(pubKey)
		Expect(err).To(BeNil())

		routingKeys := []string{"did:example:HPXoCUSjrSvWC54SLWQjsm#somekey"}

		payload := didcli.DIDDocument{
			ID: did,
			VerificationMethod: []didcli.VerificationMethod{
				map[string]any{
					"id":                 keyId,
					"type":               "Ed25519VerificationKey2020",
					"controller":         did,
					"publicKeyMultibase": pubKeyMultibase58,
				},
			},
			Authentication: []any{keyId},
			Service: []didcli.Service{
				{
					ID:              did + "#service-1",
					Type:            "type-1",
					ServiceEndpoint: []string{"endpoint-1"},
					Accept:          []string{"accept-1"},
					RoutingKeys:     routingKeys,
				},
			},
		}

		signInputs := []didcli.SignInput{
			{
				VerificationMethodID: keyId,
				PrivKey:              privKey,
			},
		}

		versionID := uuid.NewString()

		res, err := cli.CreateDidDoc(tmpDir, payload, signInputs, versionID, testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.CreateDid.String()))
		Expect(err).To(BeNil())
		Expect(res.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query diddoc"))
		// Query the DID Doc
		resp, err := cli.QueryDidDoc(did)
		Expect(err).To(BeNil())

		didDoc := resp.Value.DidDoc
		Expect(didDoc.Id).To(BeEquivalentTo(did))
		Expect(didDoc.Authentication).To(HaveLen(1))
		Expect(didDoc.Authentication[0].VerificationMethodId).To(BeEquivalentTo(keyId))
		Expect(didDoc.Authentication[0].VerificationMethod).To(BeNil())
		Expect(didDoc.VerificationMethod).To(HaveLen(1))
		Expect(didDoc.VerificationMethod[0].Id).To(BeEquivalentTo(keyId))
		Expect(didDoc.VerificationMethod[0].VerificationMethodType).To(BeEquivalentTo("Ed25519VerificationKey2020"))
		Expect(didDoc.VerificationMethod[0].Controller).To(BeEquivalentTo(did))
		Expect(didDoc.VerificationMethod[0].VerificationMaterial).To(BeEquivalentTo(pubKeyMultibase58))
		Expect(didDoc.Service).To(HaveLen(1))
		Expect(didDoc.Service[0].Id).To(BeEquivalentTo(did + "#service-1"))
		Expect(didDoc.Service[0].ServiceType).To(BeEquivalentTo("type-1"))
		Expect(didDoc.Service[0].ServiceEndpoint[0]).To(BeEquivalentTo("endpoint-1"))
		Expect(didDoc.Service[0].Accept).To(BeEquivalentTo([]string{"accept-1"}))
		Expect(didDoc.Service[0].RoutingKeys).To(BeEquivalentTo(routingKeys))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can update diddoc"))
		// Update the DID Doc

		payload2 := didcli.DIDDocument{
			ID: did,
			VerificationMethod: []didcli.VerificationMethod{
				map[string]any{
					"id":                 keyId,
					"type":               "Ed25519VerificationKey2020",
					"controller":         did,
					"publicKeyMultibase": pubKeyMultibase58,
				},
			},
			Authentication: []any{keyId},
		}

		res2, err := cli.UpdateDidDoc(tmpDir, payload2, signInputs, "", testdata.BASE_ACCOUNT_1, helpers.GenerateFees(feeParams.CreateDid.String()))
		Expect(err).To(BeNil())
		Expect(res2.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query diddoc"))
		// Query the DID Doc
		resp, err = cli.QueryDidDoc(did)
		Expect(err).To(BeNil())

		didDoc = resp.Value.DidDoc
		Expect(didDoc.Id).To(BeEquivalentTo(did))
		Expect(didDoc.Authentication).To(HaveLen(1))
		Expect(didDoc.Authentication[0].VerificationMethodId).To(BeEquivalentTo(keyId))
		Expect(didDoc.Authentication[0].VerificationMethod).To(BeNil())
		Expect(didDoc.VerificationMethod).To(HaveLen(1))
		Expect(didDoc.VerificationMethod[0].Id).To(BeEquivalentTo(keyId))
		Expect(didDoc.VerificationMethod[0].VerificationMethodType).To(BeEquivalentTo("Ed25519VerificationKey2020"))
		Expect(didDoc.VerificationMethod[0].Controller).To(BeEquivalentTo(did))
		Expect(didDoc.VerificationMethod[0].VerificationMaterial).To(BeEquivalentTo(pubKeyMultibase58))
		Expect(didDoc.Service).To(HaveLen(0))
	})
})
