/*
Copyright © 2024 Alex Theobold <theoboldalex@gmail.com>

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
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	utils "github.com/theoboldalex/lichen/pkg"
)

type LicenseBody struct {
	Content string `json:"body"`
	Pretty  string `json:"html_url"`
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a license",
	Long:  `Generate a license from the available open source licenses`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(fmt.Sprintf("%s/%v", utils.LICENSES_URL, args[0]))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatal(fmt.Sprintf("There was a problem with the request. Expected HTTP 200 but got HTTP %d", resp.StatusCode))
		}

		var l LicenseBody
		err = json.NewDecoder(resp.Body).Decode(&l)
		if err != nil {
			log.Fatal("Unable to decode JSON into struct")
		}

		// TODO: Write this to file
		fmt.Printf("%s", l.Content)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// TODO: Add flags for file name to write and location
}
