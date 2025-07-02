package unified_pro

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/unified_pro"
	"github.com/spf13/cobra"
)

var (
	bfusdCmd = &cobra.Command{
		Use:   "bfusd",
		Short: "bfusd",
	}

	mintBfUSD = &cobra.Command{
		Use:   "mint",
		Short: "mint",
		Run:   mint,
	}

	redeemBfUSD = &cobra.Command{
		Use:   "redeem",
		Short: "redeem",
		Run:   redeem,
	}
)

func InitBFUSDCmds() []*cobra.Command {
	bfusdCmd.PersistentFlags().StringP("fromAsset", "f", "", "fromAsset")
	bfusdCmd.PersistentFlags().StringP("targetAsset", "t", "", "targetAsset")
	bfusdCmd.PersistentFlags().StringP("amount", "a", "", "amount")
	bfusdCmd.MarkFlagRequired("fromAsset")
	bfusdCmd.MarkFlagRequired("targetAsset")
	bfusdCmd.MarkFlagRequired("amount")

	bfusdCmd.AddCommand(mintBfUSD, redeemBfUSD)
	return []*cobra.Command{bfusdCmd}
}

func mint(cmd *cobra.Command, _ []string) {
	fromAsset, _ := cmd.Flags().GetString("fromAsset")
	targetAsset, _ := cmd.Flags().GetString("targetAsset")
	amount, _ := cmd.Flags().GetString("amount")

	unifiedPro := unified_pro.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	response, err := unifiedPro.Mint(fromAsset, targetAsset, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mint %v BFUSD success, rate: %v\n", response.TargetAssetQty, response.MintRate)
}

func redeem(cmd *cobra.Command, _ []string) {
	fromAsset, _ := cmd.Flags().GetString("fromAsset")
	targetAsset, _ := cmd.Flags().GetString("targetAsset")
	amount, _ := cmd.Flags().GetString("amount")

	unifiedPro := unified_pro.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	response, err := unifiedPro.Redeem(fromAsset, targetAsset, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("redeem %v BFUSD success, rate: %v\n", response.TargetAssetQty, response.RedeemRate)
}
