//go:build integration

package integration

import (
	"canow-chain/tests/integration/cli"
	"canow-chain/tests/integration/testdata"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cheqd cli - negative transfert token", func() {
	It("cannot transfer token with missing cli arguments", func() {
		amount := fmt.Sprintf("%d%s", 100, testdata.DEMON)

		AddReportEntry("Integration", fmt.Sprintf("%sNegative: %s", cli.PURPLE, "cannot transfer token with missing cli arguments"))
		// Fail to transfer token with missing cli arguments
		//	 a. missing sender account address
		_, err := cli.TransferToken("", testdata.BASE_ACCOUNT_2_ADDRESS, amount)
		Expect(err).ToNot(BeNil())

		//    b. missing receiver account address
		_, err = cli.TransferToken(testdata.BASE_ACCOUNT_1_ADDRESS, "", amount)
		Expect(err).ToNot(BeNil())

		//    c. we cannot use receiver account name instead of receiver account address
		_, err = cli.TransferToken(testdata.BASE_ACCOUNT_1_ADDRESS, testdata.BASE_ACCOUNT_2, amount)
		Expect(err).ToNot(BeNil())

		//    d. missing sender and receiver account addresses
		_, err = cli.TransferToken("", "", amount)
		Expect(err).ToNot(BeNil())

		//    e. missing amount
		_, err = cli.TransferToken(testdata.BASE_ACCOUNT_1_ADDRESS, testdata.BASE_ACCOUNT_2_ADDRESS, "")
		Expect(err).ToNot(BeNil())

		AddReportEntry("Integration", fmt.Sprintf("%sNegative: %s", cli.PURPLE, "cannot transfer token with invalid amount"))
		//    a. invalid amount
		_, err = cli.TransferToken(testdata.BASE_ACCOUNT_1_ADDRESS, testdata.BASE_ACCOUNT_2_ADDRESS, fmt.Sprintf("-1%s", testdata.DEMON))
		Expect(err).ToNot(BeNil())

		_, err = cli.TransferToken(testdata.BASE_ACCOUNT_1_ADDRESS, testdata.BASE_ACCOUNT_2_ADDRESS, fmt.Sprintf("0%s", testdata.DEMON))
		Expect(err).ToNot(BeNil())
	})
})
