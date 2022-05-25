// Just a simple tool that retrieves all available pairs from bitstamp and
// generates a fresh channel enums file (channel.go)
package main

import (
	"context"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"unicode"

	"github.com/georlav/bitstamp"
)

type pair struct {
	Name    string
	URLName string
}

var channelNames = []string{
	"live_trades",
	"live_orders",
	"order_book",
	"detail_order_book",
	"diff_order_book",
	// "private-my_orders",
	// "private-my_trades",
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c := bitstamp.NewHTTPAPI()
	results, err := c.GetTradingPairsInfo(context.Background())
	if err != nil {
		log.Fatal("Failed to retrieve pairs")
	}

	var pairs []pair
	for i := range results {
		if unicode.IsDigit(rune(results[i].Name[0])) {
			continue
		}

		pairs = append(pairs, pair{
			Name:    strings.ReplaceAll(results[i].Name, "/", ""),
			URLName: results[i].URLSymbol,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})

	enums := ""
	enumValues := ""

	for i := range pairs {
		for j := range channelNames {
			splitted := strings.Split(channelNames[j], "_")
			for k := range splitted {
				// nolint: staticcheck
				splitted[k] = strings.Title(splitted[k])
			}
			splitted = append(splitted, pairs[i].Name)

			enumName := strings.Join(splitted, "")
			if strings.HasPrefix(enumName, "Private-") {
				enumName = strings.TrimPrefix(enumName, "Private-") + "Private"
			}
			enumName += "Channel"

			enums += enumName + "\n"
			if i == 0 && j == 0 {
				enums = strings.TrimSuffix(enums, "\n") + " Channel = iota\n"
			}

			enumValues += enumName + `: "` + channelNames[j] + "_" + pairs[i].URLName + `",` + "\n"
		}
	}

	code := `// Code generated by generatechannels tool. DO NOT EDIT
	package bitstamp

	type Channel uint32

	const (
		%s
	)

	func (p Channel) String() string {
		return getChannels()[p]
	}

	func getChannels() map[Channel]string {
		return map[Channel]string{
			%s
		}
	}
`
	b, err := format.Source([]byte(fmt.Sprintf(code, enums, enumValues)))
	if err != nil {
		log.Fatalf("Failed to format generated code, %s", err)
	}

	if err := ioutil.WriteFile("channel.go", b, 0664); err != nil {
		log.Fatalf("Failed to create channel enum file. %s", err)
	}
}
