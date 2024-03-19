package scheduler

import (
	"encoding/json"
	"proj3-redesigned/function"
	"log"
	"os"
	"fmt"
)

func ResultsJSON (funcsInfo *[]function.FuncInfo, dataSize string) {
	var mins []string

	for _, funcObj := range *funcsInfo {
		min := fmt.Sprintf("%.2f", funcObj.Min)
		mins = append(mins, min)
	}
	
	jsonData, err := json.MarshalIndent(mins, "", "    ")
	if err != nil {
		log.Fatalf("Error encoding to JSON: %v", err)
	}

	out_path := "./data/results_" + dataSize + ".json"
    err = os.WriteFile(out_path, jsonData, 0644)
    if err != nil {
        log.Fatalf("Error writing to file: %v", err)
    }
}
