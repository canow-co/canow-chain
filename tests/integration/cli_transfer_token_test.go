//go:build integration

package integration

import (
	"github.com/canow-co/canow-chain/tests/integration/cli"
	"github.com/canow-co/canow-chain/tests/integration/testdata"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cheqd cli - positive transfert token", func() {
	It("can transfer token and query the result", func() {
		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.GREEN, "can query account balance"))
		// Query the receiver account balance
		balance, err := cli.QueryGetBalances(testdata.BASE_ACCOUNT_2_ADDRESS)
		Expect(err).To(BeNil())
		receiverAccountBalance := balance.Balances[0].Amount.Int64()

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.GREEN, "can transfer token"))
		// Transfer token using sender account name
		var amount int64 = 100
		res, err := cli.TransferToken(testdata.BASE_ACCOUNT_1_ADDRESS, testdata.BASE_ACCOUNT_2_ADDRESS, fmt.Sprintf("%d%s", amount, testdata.DEMON))
		Expect(err).To(BeNil())
		Expect(res.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.GREEN, "can query account balance"))
		// Query the receiver account balance
		balance, err = cli.QueryGetBalances(testdata.BASE_ACCOUNT_2_ADDRESS)
		Expect(err).To(BeNil())
		Expect(balance.Balances[0].Amount.Int64()).To(BeEquivalentTo(receiverAccountBalance + amount))

		receiverAccountBalance = balance.Balances[0].Amount.Int64()

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.GREEN, "can transfer token"))
		// Transfer token using sender account address
		res, err = cli.TransferToken(testdata.BASE_ACCOUNT_1_ADDRESS, testdata.BASE_ACCOUNT_2_ADDRESS, fmt.Sprintf("%d%s", amount, testdata.DEMON))
		Expect(err).To(BeNil())
		Expect(res.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.GREEN, "can query account balance"))
		// Query the receiver account balance
		balance, err = cli.QueryGetBalances(testdata.BASE_ACCOUNT_2_ADDRESS)
		Expect(err).To(BeNil())
		Expect(balance.Balances[0].Amount.Int64()).To(BeEquivalentTo(receiverAccountBalance + amount))
	})
})
