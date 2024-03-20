package function

import (
	"encoding/json"
	"log"
	"os"
)

// Structs for function parameters and domain

type FunctionParameters struct {
    Cx3, Cy3, Cx2y, Cxy2, Cx2, Cy2, Cxy, Cx, Cy, C, Min_x, Max_x, Min_y, Max_y float32
}

type Domain struct {
    Min_x, Max_x, Min_y, Max_y float32
}

type FuncInfo struct {
    Func func(float32, float32) float32
    Dom Domain
    Min float32
}

// Reads function info from a JSON file
func ExtractFuncInfo (DataSize string) *[]FuncInfo {
	dataPathFile := "./data/" + DataSize + ".json"
	effectsFile, err := os.Open(dataPathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer effectsFile.Close()

	var funcsInfo []FuncInfo

	decoder := json.NewDecoder(effectsFile)
    _, err = decoder.Token()
    if err != nil {
        log.Fatal(err)
    }

    for decoder.More() {
        var p FunctionParameters
        err := decoder.Decode(&p)
        if err != nil {
            log.Fatal(err)
        }
        funcsInfo = append(funcsInfo, NewFuncInfo(&p))
    }

	return &funcsInfo
}

// Helper functions to create new instances of structs

func NewFunctionParameters () *FunctionParameters {
	return &FunctionParameters{Cx3: 0, Cy3: 0, Cx2y: 0, Cxy2: 0, Cx2: 0, Cy2: 0, Cxy: 0, Cx: 0, Cy: 0, C: 0, Min_x: 0, Max_x: 0, Min_y: 0, Max_y: 0}
}

func NewFunction (p *FunctionParameters) func(float32, float32) float32 {
	my_polynomial := func(x float32, y float32) float32 {
		return p.Cx3*x*x*x + p.Cy3*y*y*y + p.Cx2y*x*x*y + p.Cxy2*x*y*y + p.Cx2*x*x + p.Cy2*y*y + p.Cxy*x*y + p.Cx*x + p.Cy*y + p.C
	}
	return my_polynomial
}

func NewDomain (p *FunctionParameters) Domain {
	return Domain{p.Min_x, p.Max_x, p.Min_y, p.Max_y}
}

// Returns a new FuncInfo struct
func NewFuncInfo (p *FunctionParameters) FuncInfo {
	return FuncInfo{NewFunction(p), NewDomain(p), 0}
}
