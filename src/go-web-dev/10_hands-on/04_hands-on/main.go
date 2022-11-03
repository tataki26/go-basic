package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Record struct {
	Date        time.Time
	Open, Close float64
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	// parse csv
	records := parseCsv("table.csv")

	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template
	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}

}

func parseCsv(filePath string) []Record {
	// open csv file
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	// read csv file
	// rows -> [][]string
	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// make a data structure
	records := make([]Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)
		close, _ := strconv.ParseFloat(row[2], 64)

		records = append(records, Record{
			Date:  date,
			Open:  open,
			Close: close,
		})
	}

	return records

}
