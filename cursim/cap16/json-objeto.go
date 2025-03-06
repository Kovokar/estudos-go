package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Veiculo struct {
	ID    int32
	Name  string
	Rodas int
}

func main() {

	v := []Veiculo{}
	json_str := `[{"ID":1,"Name":"Moto","Rodas":2},{"ID":2,"Name":"Carro","Rodas":4},{"ID":3,"Name":"onibus","Rodas":6}]`

	fmt.Println("COM UNMARSHAL")
	b := []byte(json_str)

	err := json.Unmarshal(b, &v)

	if err != nil {
		fmt.Println("SATANENGA")
		panic(err)
	}
	fmt.Println(v)

	// COM DECODER
	fmt.Println("COM DECODER")
	dcdr := json.NewDecoder(strings.NewReader(json_str))
	err2 := dcdr.Decode(&v)

	if err2 != nil {
		fmt.Println("caoooo")
	}

	fmt.Println(v)

}
