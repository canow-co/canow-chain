//go:build integration

package integration

import (
	"fmt"

	"github.com/canow-co/canow-chain/tests/integration/cli"
	"github.com/canow-co/canow-chain/tests/integration/testdata"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cheqd cli - positive transfert token", func() {
	It("can transfer token and query the result", func() {
		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query account balance"))
		// Query the receiver account balance
		balance, err := cli.QueryGetBalances(testdata.BASE_ACCOUNT_2_ADDR)
		Expect(err).To(BeNil())
		receiverAccountBalance := balance.Balances[0].Amount.Int64()

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can transfer token"))
		// Transfer token using sender account name
		var amount int64 = 100
		res, err := cli.TransferToken(testdata.BASE_ACCOUNT_1_ADDR, testdata.BASE_ACCOUNT_2_ADDR, fmt.Sprintf("%d%s", amount, testdata.DEMON), cli.CliGasParams)
		Expect(err).To(BeNil())
		Expect(res.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query account balance"))
		// Query the receiver account balance
		balance, err = cli.QueryGetBalances(testdata.BASE_ACCOUNT_2_ADDR)
		Expect(err).To(BeNil())
		Expect(balance.Balances[0].Amount.Int64()).To(BeEquivalentTo(receiverAccountBalance + amount))

		receiverAccountBalance = balance.Balances[0].Amount.Int64()

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can transfer token"))
		// Transfer token using sender account address
		res, err = cli.TransferToken(testdata.BASE_ACCOUNT_1_ADDR, testdata.BASE_ACCOUNT_2_ADDR, fmt.Sprintf("%d%s", amount, testdata.DEMON), cli.CliGasParams)
		Expect(err).To(BeNil())
		Expect(res.Code).To(BeEquivalentTo(0))

		AddReportEntry("Integration", fmt.Sprintf("%sPositive: %s", cli.Green, "can query account balance"))
		// Query the receiver account balance
		balance, err = cli.QueryGetBalances(testdata.BASE_ACCOUNT_2_ADDR)
		Expect(err).To(BeNil())
		Expect(balance.Balances[0].Amount.Int64()).To(BeEquivalentTo(receiverAccountBalance + amount))
	})
})
