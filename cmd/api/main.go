package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"koto.com/internal/core/ports"
	"koto.com/internal/service"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {

		text := s.Text()
		capitalGainService := service.NewCapitalGainService(ports.NewStockService())
		if text == "exit" {
			break
		}
		if text == "quit" {
			break
		}

		cps, err := capitalGainService.GetCapitalGain(text)

		if err != nil {
			if err.Error() == "input string is empty" {
				fmt.Println("Input string is empty, please provide valid stock operations.")
				break

			}
			panic(err)
		}
		out, _ := json.Marshal(cps)
		fmt.Println(string(out))
		continue

	}

}
