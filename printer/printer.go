package printer

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type TableWriter interface {
	Header() []string
	Row() [][]any
}

func PrintTable(writer TableWriter) {
	table := tablewriter.NewTable(os.Stdout, tablewriter.WithEastAsian(false))
	table.Header(writer.Header())
	table.Bulk(writer.Row())
	table.Render()
}
