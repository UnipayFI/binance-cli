package futures

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	accountCmd = &cobra.Command{
		Use:   "account",
		Short: "show account balances & account config",
	}

	balancesCmd = &cobra.Command{
		Use:   "balances",
		Short: "show account balances",
		Long: `Get current account's balances.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Futures-Account-Balance-V3`,
		Run: balances,
	}

	accountConfigCmd = &cobra.Command{
		Use:   "config",
		Short: "show account config",
		Long: `Query account configuration.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Account-Config`,
		Run: accountConfig,
	}
)

func InitAccountCmds() []*cobra.Command {
	accountCmd.AddCommand(balancesCmd, accountConfigCmd)
	return []*cobra.Command{accountCmd}
}

func balances(cmd *cobra.Command, args []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	balances, err := client.GetBalances()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(balances)
}

func accountConfig(cmd *cobra.Command, args []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	config, err := client.GetAccountConfig()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(config)
}
