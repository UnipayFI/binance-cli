package unified

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/unified"
	"github.com/spf13/cobra"
)

var (
	autoCollectionCmd = &cobra.Command{
		Use:   "auto-collection",
		Short: "auto-collection",
		Run:   autoCollection,
	}
)

func InitAutoCollectionCmds() []*cobra.Command {
	return []*cobra.Command{
		autoCollectionCmd,
	}
}

func autoCollection(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	resp, err := client.AutoCollection()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("auto collection: %v\n", resp.Msg)
}
