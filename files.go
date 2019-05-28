package utils

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func OpenCSV(fileName string) (ret [][]float64) {
	csvfile, err := os.Open(fileName + ".csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	for {
		line, err := r.Read()
		var lineData []float64
		if err == io.EOF {
			break
		}
		CheckErr(err)
		for j := 0; j < len(line)-1; j++ {
			val, err := strconv.ParseFloat(line[j], 64)
			CheckErr(err)
			lineData = append(lineData, val)
		}
		ret = append(ret, lineData)
	}
	return ret
}

func LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count + 1, nil

		case err != nil:
			return count + 1, err
		}
	}
}
