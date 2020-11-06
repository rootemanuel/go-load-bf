package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

type Reg struct {
	DataRef       string
	DataComp      string
	Uf            string
	CodMunicipio  string
	NomeMunicipio string
	Cpf           string
	NisFav        string
	NomeFav       string
	ValorParcela  string
}

func main() {
	log.Print("ini")
	arr := make([]Reg, 0)

	rc, _ := os.Open("/Users/h4x0rs/go/src/github.com/rootemanuel/go-load-bf/arquivo/202010_BolsaFamilia_Pagamentos.csv")
	ch := make(chan bool, 1)

	r := csv.NewReader(rc)
	r.LazyQuotes = true

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	defer close(ch)

	//root thread var wg sync.WaitGroup

	for {
		//root thread defer wg.Done()
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)

		}

		//root thread wg.Add(1)
		//root thread go func() {

		spl := strings.Split(rec[0], ";")

		row := Reg{
			DataRef:       spl[0],
			DataComp:      spl[1],
			Uf:            spl[2],
			CodMunicipio:  spl[3],
			NomeMunicipio: spl[4],
			Cpf:           spl[5],
			NisFav:        spl[6],
			NomeFav:       spl[7],
			ValorParcela:  spl[8],
		}

		arr = append(arr, row)
		//root thread 	}()
	}

	log.Print(len(arr))
	//root thread wg.Wait()
	log.Print("fim")
}
