/*
Copyright © 2020 Dênnis Dantas de Sousa <ddspog@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"github.com/ddsgok/ts-wizard/reader"
	"github.com/spf13/cobra"
)

const (
	defaultNumForFilteringNullPackets        = -1
	defaultNumForFilteringStreamTypeCPackets = -1
)

var filterNNullPackets int
var filterNStreamTypeCPackets int
var filterStreamTypeCPIDs []uint

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter <options> ts-source-filepath destination-path",
	Short: "Reduce TS size by removing some packages.",
	Long: `Reduce TS size by filtering a number of selected packages, while removing the unneded others:

You can limit the number of packages of a single type.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		reader.ReadTS(args[0], args[1], filterNNullPackets, filterNStreamTypeCPackets, filterStreamTypeCPIDs)
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	filterCmd.Flags().IntVar(&filterNNullPackets, "null-packets", defaultNumForFilteringNullPackets, "When used, will filter the transport stream to contain only a certain number of null packages.")
	filterCmd.Flags().IntVar(&filterNStreamTypeCPackets, "streamtc-packets", defaultNumForFilteringStreamTypeCPackets, "When used will filter the transport stream to contain only a certain number of stream type C packages.")
	filterCmd.Flags().UintSliceVar(&filterStreamTypeCPIDs, "streamtc-pids", []uint{}, "When used, will filter the transport stream to contain only a certain number of stream type C packages, from specified PIDs.")
}
