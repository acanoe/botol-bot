/*
Copyright © 2020 Hendika N.

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
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
)

var botToken string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the bot using your bot token",
	Long:  `To start your bot, pass --botToken argument with the bot token you obtained from BotFather`,
	Args: func(cmd *cobra.Command, args []string) error {
		var err string

		if len(args) < 1 {
			err = "Please specify bot token using --botToken or -k flag"
		}
		return errors.New(err)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		b, err := tb.NewBot(tb.Settings{
			Token:  botToken,
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatal(err)
			return
		}

		b.Handle("/start", func(m *tb.Message) {
			if m.Private() {
				b.Send(m.Sender, "Hello World!")
			}
		})

		b.Handle("/hello", func(m *tb.Message) {
			b.Send(m.Sender, "Hello to you too")
		})

		b.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&botToken, "botToken", "k", "", "Bot Token (you can get one from BotFather)")
	viper.BindPFlag("botToken", rootCmd.PersistentFlags().Lookup("botToken"))
}
