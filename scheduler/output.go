package scheduler

import (
	"encoding/json"
	"proj/function"
	"log"
	"os"
	"fmt"
)

// Writes the minimum values of the given functions to a JSON file
func ResultsJSON (funcsInfo *[]function.FuncInfo, dataSize string) {
	var mins []string

	// Convert minimum values to strings
	for _, funcObj := range *funcsInfo {
		min := fmt.Sprintf("%.2f", funcObj.Min)
		mins = append(mins, min)
	}
	
	// Encode minimum values to JSON	
	jsonData, err := json.MarshalIndent(mins, "", "    ")
	if err != nil {
		log.Fatalf("Error encoding to JSON: %v", err)
	}

	// Write JSON data to file
	out_path := "./data/results_" + dataSize + ".json"
    err = os.WriteFile(out_path, jsonData, 0644)
    if err != nil {
        log.Fatalf("Error writing to file: %v", err)
    }
}
