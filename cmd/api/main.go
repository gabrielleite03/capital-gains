package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"koto.com/internal/core/models"
	"koto.com/internal/core/ports"
	"koto.com/internal/service"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	allOps := []*[]models.CapitalGains{}
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
				for _, ops := range allOps {
					out, _ := json.Marshal(ops)
					fmt.Println(string(out))

				}
				os.Exit(0)

			}
			// panic(err)
		}
		allOps = append(allOps, cps)

		continue

	}

}
