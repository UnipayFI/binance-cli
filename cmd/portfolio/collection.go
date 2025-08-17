package portfolio

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio"
	"github.com/spf13/cobra"
)

var (
	collectionCmd = &cobra.Command{
		Use:   "collection",
		Short: "Collection",
	}

	autoCollectionCmd = &cobra.Command{
		Use:   "auto-collection",
		Short: "Fund collection for Portfolio Margin",
		Long: `Fund collection for Portfolio Margin.
- The BNB would not be collected from UM-PM account to the Portfolio Margin account.
- You can only use this function 500 times per hour in a rolling manner.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Fund-Auto-collection`,
		Run: autoCollection,
	}

	assetCollectionCmd = &cobra.Command{
		Use:   "asset-collection",
		Short: "Transfers specific asset from Futures Account to Margin account",
		Long: `Transfers specific asset from Futures Account to Margin account.
- The BNB transfer is not be supported

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Fund-Collection-by-Asset`,
		Run: assetCollection,
	}
)

func init() {
	assetCollectionCmd.Flags().StringP("asset", "a", "", "asset")
	assetCollectionCmd.MarkFlagRequired("asset")
}

func InitCollectionCmds() []*cobra.Command {
	return []*cobra.Command{
		collectionCmd,
		autoCollectionCmd,
		assetCollectionCmd,
	}
}

func autoCollection(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	resp, err := client.AutoCollection()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("auto collection: %v\n", resp.Msg)
}

func assetCollection(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	asset, _ := cmd.Flags().GetString("asset")
	resp, err := client.AssetCollection(asset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("asset collection: %v\n", resp.Msg)
}
