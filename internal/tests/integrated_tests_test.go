package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	reflect "reflect"
	"strings"
	"testing"

	"koto.com/internal/core/infrastructure"
	"koto.com/internal/core/models"
	"koto.com/internal/core/ports"
	"koto.com/service"
)

func TestCase_1_case_2(t *testing.T) {
	sInput := readFile(infrastructure.Case1Case2InputFile)

	sOutput := readFile(infrastructure.Case1Case2OutputFile)
	cg := service.NewCapitalGainService(ports.NewStockService())
	ops, err := cg.GetCapitalGain(sInput)
	if err != nil {
		t.Errorf("Error getting capital gains: %v", err)
	}

	cGModelsOutput := getModelFromString(sOutput)

	if !reflect.DeepEqual(ops, cGModelsOutput) {
		t.Errorf("Expected output: %v, got: %v", cGModelsOutput, ops)
	}

}

func TestCase_all_files(t *testing.T) {
	filesQuantity := 9
	for i := 1; i < filesQuantity; i++ {
		file := fmt.Sprintf("%s%d%s", infrastructure.PrefixUseCase, i, infrastructure.SufixUseCase)
		fmt.Println(file)
		t.Run(fmt.Sprintf("Running tests for %s", file), func(t *testing.T) {
			sInput := readFile(file)

			sOutput := readFile(file)
			cg := service.NewCapitalGainService(ports.NewStockService())
			ops, err := cg.GetCapitalGain(sInput)
			if err != nil {
				t.Errorf("Error getting capital gains: %v", err)
			}

			cGModelsOutput := getModelFromString(sOutput)

			if !reflect.DeepEqual(ops, cGModelsOutput) {
				t.Errorf("Expected output: %v, got: %v", cGModelsOutput, ops)
			}
		})

	}

}

func readFile(f string) string {
	file, err := os.Open(infrastructure.TestCasesPath + f)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)
	return string(b)
}

func getModelFromString(row string) *[][]models.CapitalGains {
	s := strings.Split(row, "\r\n")
	cgs := &[][]models.CapitalGains{}
	for _, r := range s {
		if r != "" {
			var cgRow []models.CapitalGains
			e := json.Unmarshal([]byte(r), &cgRow)
			if e != nil {
				log.Fatalf("Error unmarshalling operation: %v", e)
			}
			*cgs = append(*cgs, cgRow)
		}
	}

	return cgs
}
