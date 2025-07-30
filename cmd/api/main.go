package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"koto.com/internal/core/ports"
	"koto.com/service"
)

func main() {

	s := bufio.NewScanner(os.Stdin)
	var outs []string
	for s.Scan() {

		text := s.Text()
		if text == "" {
			break
		}
		capitalGainService := service.NewCapitalGainService(ports.NewStockService())

		cps, _ := capitalGainService.GetCapitalGain(text)
		out, _ := json.Marshal(cps)
		outs = append(outs, string(out))

	}

	for _, row := range outs {
		fmt.Println(row)
	}

}
