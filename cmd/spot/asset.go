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
)

func InitAssetCmds() []*cobra.Command {
	assetCmd.AddCommand(&cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list assets",
		Run:     listAssets,
	})
	return []*cobra.Command{assetCmd}
}

func listAssets(cmd *cobra.Command, args []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	assets, err := client.GetAssetList()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&assets)
}
