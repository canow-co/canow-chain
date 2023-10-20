//go:build upgrade_integration

package integration

import (
	"fmt"
	"path/filepath"

	clihelpers "github.com/canow-co/canow-chain/tests/integration/helpers"
	cli "github.com/canow-co/canow-chain/tests/upgrade/integration/cli"
	didcli "github.com/canow-co/cheqd-node/x/did/client/cli"
	didtypes "github.com/canow-co/cheqd-node/x/did/types"
	resourcetypes "github.com/canow-co/cheqd-node/x/resource/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Upgrade - Post", Ordered, func() {
	var feeParams didtypes.FeeParams
	var resourceFeeParams resourcetypes.FeeParams

	BeforeAll(func() {
		// Query fee params
		res, err := cli.QueryParams(cli.Validator0, didtypes.ModuleName, string(didtypes.ParamStoreKeyFeeParams))
		Expect(err).To(BeNil())
		err = clihelpers.Codec.UnmarshalJSON([]byte(res.Value), &feeParams)
		Expect(err).To(BeNil())

		res, err = cli.QueryParams(cli.Validator0, resourcetypes.ModuleName, string(resourcetypes.ParamStoreKeyFeeParams))
		Expect(err).To(BeNil())
		err = clihelpers.Codec.UnmarshalJSON([]byte(res.Value), &resourceFeeParams)
		Expect(err).To(BeNil())
	})

	Context("After a software upgrade execution has concluded", Ordered, func() {
		It("should wait for node catching up", func() {
			By("pinging the node status until catching up is flagged as false")
			err := cli.WaitForCaughtUp(cli.Validator0, cli.CliBinaryName, cli.VotingPeriod*6)
			Expect(err).To(BeNil())
		})

		It("should wait for a certain number of blocks to be produced", func() {
			By("fetching the current chain height")
			currentHeight, err := cli.GetCurrentBlockHeight(cli.Validator0, cli.CliBinaryName)
			Expect(err).To(BeNil())

			By("waiting for 10 blocks to be produced on top, after the upgrade")
			err = cli.WaitForChainHeight(cli.Validator0, cli.CliBinaryName, currentHeight+10, cli.VotingPeriod*6)
			Expect(err).To(BeNil())
		})

		It("should match the expected module version map", func() {
			By("loading the expected module version map")
			var expected upgradetypes.QueryModuleVersionsResponse
			_, err := Loader(
				filepath.Join(GeneratedJSONDir, "post", "query - module-version-map", fmt.Sprintf("%s.json", cli.UpgradeName)),
				&expected,
			)
			Expect(err).To(BeNil())

			By("matching the expected module version map")
			actual, err := cli.QueryModuleVersionMap(cli.Validator0)
			Expect(err).To(BeNil())

			Expect(actual.ModuleVersions).To(Equal(expected.ModuleVersions), "module version map mismatch")
		})

		It("should load and run existing diddoc payloads - case: update", func() {
			By("matching the glob pattern for existing diddoc payloads")
			DidDocUpdatePayloads, err := RelGlob(GeneratedJSONDir, "post", "update - diddoc", "*.json")
			Expect(err).To(BeNil())

			for _, payload := range DidDocUpdatePayloads {
				var DidDocUpdatePayload didcli.DIDDocument
				var DidDocUpdateSignInput []didcli.SignInput

				testCase := GetCaseName(payload)
				By("Running: " + testCase)
				fmt.Println("Running: " + testCase)

				By("reading ")
				DidDocUpdateSignInput, err = Loader(payload, &DidDocUpdatePayload)
				Expect(err).To(BeNil())

				tax := feeParams.UpdateDid.String()
				res, err := cli.UpdateDid(DidDocUpdatePayload, DidDocUpdateSignInput, cli.Validator0, "", tax)
				Expect(err).To(BeNil())
				Expect(res.Code).To(BeEquivalentTo(0))
			}
		})

		It("should load and run existing diddoc payloads - case: deactivate", func() {
			By("matching the glob pattern for existing diddoc payloads")
			DidDocDeactivatePayloads, err := RelGlob(GeneratedJSONDir, "post", "deactivate - diddoc", "*.json")
			Expect(err).To(BeNil())

			for _, payload := range DidDocDeactivatePayloads {
				var DidDocDeactivatePayload didtypes.MsgDeactivateDidDocPayload
				var DidDocDeactivateSignInput []didcli.SignInput

				testCase := GetCaseName(payload)
				By("Running: " + testCase)
				fmt.Println("Running: " + testCase)

				By("reading ")
				DidDocDeactivateSignInput, err = Loader(payload, &DidDocDeactivatePayload)
				Expect(err).To(BeNil())

				tax := feeParams.DeactivateDid.String()
				res, err := cli.DeactivateDid(DidDocDeactivatePayload, DidDocDeactivateSignInput, cli.Validator0, tax)
				Expect(err).To(BeNil())
				Expect(res.Code).To(BeEquivalentTo(0))
			}
		})

		It("should create resources after upgrade", func() {
			By("matching the glob pattern for resource payloads to create")
			ResourcePayloads, err := RelGlob(GeneratedJSONDir, "post", "create - resource", "*.json")
			Expect(err).To(BeNil())

			for _, payload := range ResourcePayloads {
				var ResourceCreatePayload resourcetypes.MsgCreateResourcePayload
				var ResourceCreateSignInput []didcli.SignInput

				testCase := GetCaseName(payload)
				By("Running: create " + testCase)
				fmt.Println("Running: " + testCase)

				ResourceCreateSignInput, err := Loader(payload, &ResourceCreatePayload)
				Expect(err).To(BeNil())

				ResourceFile, err := CreateTestJSON(GinkgoT().TempDir(), ResourceCreatePayload.Data)
				Expect(err).To(BeNil())

				res, err := cli.CreateResource(
					ResourceCreatePayload,
					ResourceFile,
					ResourceCreateSignInput,
					cli.Validator0,
					resourceFeeParams.Json.String(),
				)

				Expect(err).To(BeNil())
				Expect(res.Code).To(Equal(uint32(0)))
			}
		})

		It("should load and run expected diddoc payloads", func() {
			By("matching the glob pattern for existing diddoc payloads")
			ExpectedDidDocRecords, err := RelGlob(GeneratedJSONDir, "post", "query - diddoc", "*.json")
			Expect(err).To(BeNil())

			for _, payload := range ExpectedDidDocRecords {
				var DidDocRecord didtypes.DidDoc

				testCase := GetCaseName(payload)
				By("Running: query " + testCase)
				fmt.Println("Running: " + testCase)

				_, err = Loader(payload, &DidDocRecord)
				Expect(err).To(BeNil())

				res, err := cli.QueryDid(DidDocRecord.Id, cli.Validator0)
				Expect(err).To(BeNil())

				if DidDocRecord.Context == nil {
					DidDocRecord.Context = []string{}
				}
				if DidDocRecord.Authentication == nil {
					DidDocRecord.Authentication = []*didtypes.VerificationRelationship{}
				}
				if DidDocRecord.AssertionMethod == nil {
					DidDocRecord.AssertionMethod = []*didtypes.VerificationRelationship{}
				}
				if DidDocRecord.CapabilityInvocation == nil {
					DidDocRecord.CapabilityInvocation = []*didtypes.VerificationRelationship{}
				}
				if DidDocRecord.CapabilityDelegation == nil {
					DidDocRecord.CapabilityDelegation = []*didtypes.VerificationRelationship{}
				}
				if DidDocRecord.KeyAgreement == nil {
					DidDocRecord.KeyAgreement = []*didtypes.VerificationRelationship{}
				}
				if DidDocRecord.Service == nil {
					DidDocRecord.Service = []*didtypes.Service{}
				}
				for _, Service := range DidDocRecord.Service {
					if Service.Accept == nil {
						Service.Accept = []string{}
					}
					if Service.RoutingKeys == nil {
						Service.RoutingKeys = []string{}
					}
				}
				if DidDocRecord.AlsoKnownAs == nil {
					DidDocRecord.AlsoKnownAs = []string{}
				}

				Expect(*res.Value.DidDoc).To(Equal(DidDocRecord))
			}
		})

		It("should load and run expected resource payloads", func() {
			By("matching the glob pattern for existing resource payloads")
			ExpectedResourceRecords, err := RelGlob(GeneratedJSONDir, "post", "query - resource", "*.json")
			Expect(err).To(BeNil())

			for _, payload := range ExpectedResourceRecords {
				var ResourceRecord resourcetypes.ResourceWithMetadata

				testCase := GetCaseName(payload)
				By("Running: query " + testCase)
				fmt.Println("Running: " + testCase)

				_, err = Loader(payload, &ResourceRecord)
				Expect(err).To(BeNil())

				res, err := cli.QueryResource(ResourceRecord.Metadata.CollectionId, ResourceRecord.Metadata.Id, cli.Validator0)

				Expect(err).To(BeNil())

				Expect(res.Resource.Resource.Data).To(Equal(ResourceRecord.Resource.Data))

				Expect(res.Resource.Metadata.Id).To(Equal(ResourceRecord.Metadata.Id))
				Expect(res.Resource.Metadata.CollectionId).To(Equal(ResourceRecord.Metadata.CollectionId))
				Expect(res.Resource.Metadata.Name).To(Equal(ResourceRecord.Metadata.Name))
				Expect(res.Resource.Metadata.Version).To(Equal(ResourceRecord.Metadata.Version))
				Expect(res.Resource.Metadata.ResourceType).To(Equal(ResourceRecord.Metadata.ResourceType))
				Expect(res.Resource.Metadata.AlsoKnownAs).To(Equal(ResourceRecord.Metadata.AlsoKnownAs))
				Expect(res.Resource.Metadata.MediaType).To(Equal(ResourceRecord.Metadata.MediaType))
				// Created is populated on successful creation. We are ignoring it here.
				// Expect(res.Resource.Metadata.Created).To(Equal(ResourceCreateRecord.Metadata.Created))
				Expect(res.Resource.Metadata.Checksum).To(Equal(ResourceRecord.Metadata.Checksum))
				Expect(res.Resource.Metadata.PreviousVersionId).To(Equal(ResourceRecord.Metadata.PreviousVersionId))
				Expect(res.Resource.Metadata.NextVersionId).To(Equal(ResourceRecord.Metadata.NextVersionId))
			}
		})
	})
})
