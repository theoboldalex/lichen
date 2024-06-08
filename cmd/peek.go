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
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
	utils "github.com/theoboldalex/lichen/pkg"
)

// peekCmd represents the peek command
var peekCmd = &cobra.Command{
	Use:   "peek",
	Short: "See a license's content",
	Long:  `Peek a license's content either dorectly in the temrinal or visist it int the browser`,
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

		prettyFlag, err := cmd.Flags().GetBool("pretty")
		if err != nil {
			log.Fatal(err)
		}

		var c *exec.Cmd

		if prettyFlag {
			switch runtime.GOOS {
			case "darwin":
				c = exec.Command("open", l.Pretty)
			case "linux":
				c = exec.Command("xdg-open", l.Pretty)
			default:
				fmt.Println("Get tae fuck")
			}

			err := c.Start()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("%v", l.Content)
		}
	},
}

func init() {
	rootCmd.AddCommand(peekCmd)
	peekCmd.Flags().BoolP("pretty", "p", false, "If passed, open the browser version of the license at choosealicense.com")
}
