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
		Short: "collection",
	}

	autoCollectionCmd = &cobra.Command{
		Use:   "auto-collection",
		Short: "auto-collection",
		Run:   autoCollection,
	}

	assetCollectionCmd = &cobra.Command{
		Use:   "asset-collection",
		Short: "asset-collection",
		Run:   assetCollection,
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
