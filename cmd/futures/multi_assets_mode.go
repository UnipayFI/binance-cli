package futures

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/spf13/cobra"
)

var (
	multiAssetsModeCmd = &cobra.Command{
		Use:   "multi-assets-mode",
		Short: "Show and set multi-assets mode",
	}

	showMultiAssetsModeCmd = &cobra.Command{
		Use:   "show",
		Short: "Show multi-assets mode",
		Long: `Get user's Multi-Assets mode (Multi-Assets Mode or Single-Asset Mode) on Every symbol.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Current-Multi-Assets-Mode`,
		Run: showMultiAssetsMode,
	}

	setMultiAssetsModeCmd = &cobra.Command{
		Use:   "set",
		Short: "Set multi-assets mode",
		Long: `Change user's Multi-Assets mode (Multi-Assets Mode or Single-Asset Mode) on Every symbol.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Change-Multi-Assets-Mode`,
		Run: setMultiAssetsMode,
	}
)

func InitMultiAssetsModeCmds() []*cobra.Command {
	setMultiAssetsModeCmd.Flags().BoolP("multiAssetsMargin", "e", false, "enable multi-assets mode")
	setMultiAssetsModeCmd.MarkFlagRequired("multiAssetsMargin")

	multiAssetsModeCmd.AddCommand(showMultiAssetsModeCmd, setMultiAssetsModeCmd)
	return []*cobra.Command{multiAssetsModeCmd}
}

func showMultiAssetsMode(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	multiAssetsMode, err := client.GetMultiAssetsMode()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("multi-assets mode is:", multiAssetsMode)
}

func setMultiAssetsMode(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	multiAssetsMode, _ := cmd.Flags().GetBool("multiAssetsMargin")
	err := client.SetMultiAssetsMode(multiAssetsMode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("multi-assets mode is set to:", multiAssetsMode)
}
