package spot

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/spot"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	assetCmd = &cobra.Command{
		Use:   "asset",
		Short: "Show account assets",
	}

	assetListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list assets",
		Long: `Get current account assets.

Docs Link: https://developers.binance.com/docs/binance-spot-api-docs/testnet/rest-api/account-endpoints#account-information-user_data`,
		Run: listAssets,
	}
)

func InitAssetCmds() []*cobra.Command {
	assetCmd.AddCommand(assetListCmd)
	assetListCmd.Flags().BoolP("omitZeroBalances", "z", true, "omit zero balances")
	return []*cobra.Command{assetCmd}
}

func listAssets(cmd *cobra.Command, args []string) {
	omitZeroBalances, _ := cmd.Flags().GetBool("omitZeroBalances")
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	assets, err := client.GetAssetList(omitZeroBalances)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&assets)
}
