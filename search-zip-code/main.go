package main

// example: go run main.go http://viacep.com.br/ws/35780000/json/
// build: go build -o cep main.go
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Request failed: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read response: %v\n", err)
		}

		var data Address
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to unmarshal json: %v\n", err)
		}

		fmt.Println(data)

		file, err := os.Create("address.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create file: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s\n", data.CEP))
	}
}
