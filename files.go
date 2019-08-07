package utils

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func OpenJSON(fileName string) (ret interface{}) {
	file, err := ioutil.ReadFile("./" + fileName + ".json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &ret)
	CheckErr(err)

	return ret
}

func SaveJSON(fileName string, data interface{}, separator string) {
	f, err := os.Create(fileName + ".json")
	CheckErr(err)
	defer f.Close()

	out, err := json.MarshalIndent(data, "", separator)
	CheckErr(err)
	_, err = f.WriteString(string(out))
	CheckErr(err)
}

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
