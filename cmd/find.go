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
	"fmt"
	"io"

	"github.com/ddsgok/ts-wizard/reader"
	"github.com/spf13/cobra"
)

var findCorruptedPackets bool

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find <options> ts-filepath",
	Short: "Execute search operations on the transport stream file.",
	Long: `Execute search operations on the transport stream file.
Currently, it's only able to find packets with its error indicator set to true.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p, err := reader.FindOnTS(args[0], findCorruptedPackets)
		if err == nil {
			fmt.Printf("Found package %d\n", p)
		}

		if err == io.EOF && p == -1 {
			fmt.Println("Found nothing")
		}
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	findCmd.Flags().BoolVar(&findCorruptedPackets, "corrupted", false, "Find a corrupted package")
}
