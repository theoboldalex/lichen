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

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	utils "github.com/theoboldalex/lichen/pkg"
)

type License struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available licenses",
	Long:  `Show a list of all availabale open source liceses that can be generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(utils.LICENSES_URL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatal(fmt.Sprintf("There was a problem with the request. Expected HTTP 200 but got HTTP %d", resp.StatusCode))
		}

		var ls []License
		err = json.NewDecoder(resp.Body).Decode(&ls)
		if err != nil {
			log.Fatal("Unable to decode JSON into struct")
		}

		boldYellow := color.New(color.FgYellow).Add(color.Bold)
		boldBlue := color.New(color.FgBlue).Add(color.Bold)
		for _, l := range ls {
			boldBlue.Printf("- ")
			fmt.Printf("%s (", l.Name)
			boldYellow.Printf("%s", l.Key)
			fmt.Println(")")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
