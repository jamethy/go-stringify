package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `
Usage of %s:
Stringifies json into an escaped string that can be put into json.
I found this useful for creating AWS lambda API Gateway Proxy tests.

By default, gets the string from the clipboard and replaces it with the new string.
Alternatively you can accept user input, i.e.:

	stringify -i <<EOF
		{
			"field": "one"
		}
	EOF
	"{\"field\":\"one\"}"

(it will also copy the above to your clipboard)

Arguments:
`, os.Args[0])
		flag.PrintDefaults()
	}
	useInput := flag.Bool("i", false, "set to accept heredoc input")
	flag.Parse()

	var jsonString string

	// get data from either user input or system clipboard
	if *useInput {
		reader := bufio.NewReader(os.Stdin)

		buf := make([]byte, 100)
		for {
			n, err := reader.Read(buf)
			if err != nil {
				break
			}

			buf = buf[:n]
			jsonString += string(buf)
		}
	} else { // use clipboard
		var err error
		if jsonString, err = clipboard.ReadAll(); err != nil {
			panic(err)
		}
	}

	// unmarshal data
	var data interface{}
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		panic(err)
	}

	// remarshal data twice to escape string
	res, _ := json.Marshal(data)
	res, _ = json.Marshal(string(res))

	// print and copy to clipboard
	fmt.Println(string(res))
	if err := clipboard.WriteAll(string(res)); err != nil {
		panic(err)
	}
}
