package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// For conversion tool visit mholt.github.io/json-to-go/

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Alive bool   `json:"Alive"`
}

func main() {
	var p []person
	b_data := `[{"First":"Raghu","Last":"Ram","Alive":true},{"First":"A","Last":"H","Alive":false}]`

	fmt.Println(b_data)

	err := json.Unmarshal([]byte(b_data), &p)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Json Coversion Format\n", p)

	err = json.NewEncoder(os.Stdout).Encode(p)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	
}