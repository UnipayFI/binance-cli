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
		Use: "asset",
	}

	assetListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list assets",
		Long:    "",
		Run:     listAssets,
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
