package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	path := flag.String("path", "./data.csv", "Path of the file")
	flag.Parse()
	fileBytes, fileNPath := ReadCSV(path)
	SaveFile(fileBytes, fileNPath)
	fmt.Println(strings.Repeat("=", 10), "Done", strings.Repeat("=", 10))
}

// ReadCSV to read the content of CSV File
func ReadCSV(path *string) ([]byte, string) {
	csvFile, err := os.Open(*path)

	if err != nil {
		log.Fatal("The file is not found || wrong root")
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	content, _ := reader.ReadAll()

	if len(content) < 1 {
		log.Fatal("Something wrong, the file maybe empty or length of the lines are not the same")
	}

	headersArr := make([]string, 0)
	for _, headE := range content[0] {
		headersArr = append(headersArr, headE)
	}

	//Remove the header row
	content = content[1:]

	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, d := range content {
		buffer.WriteString("{")
		for j, y := range d {
			// \" -> '
			y = strings.Replace(y, `\"`, "'", -1)
			// """""" -> ''
			y = strings.Replace(y, `""""""`, "''", -1)
			// """" -> ''
			y = strings.Replace(y, `""""`, "''", -1)
			// """ -> "
			y = strings.Replace(y, `"""`, `"`, -1)
			// "" -> empty
			y = strings.Replace(y, `""`, "", -1)
			// '' -> ""
			y = strings.Replace(y, "''", `""`, -1)
			// trim external quotes
			y = strings.Trim(y, `"`)
			buffer.WriteString(`"` + headersArr[j] + `":`)
			if len(y) == 0 {
				buffer.WriteString(`""`)
			} else {
				_, fErr := strconv.ParseFloat(y, 32)
				_, bErr := strconv.ParseBool(y)
				if fErr == nil {
					buffer.WriteString(y)
				} else if bErr == nil {
					buffer.WriteString(strings.ToLower(y))
				} else {
					// make room for json objects and arrays
					switch string(y[0]) {
					case "{":
						fallthrough
					case "[":
						buffer.WriteString(y)
					default:
						buffer.WriteString(`"` + y + `"`)
					}
				}
			}
			//end of property
			if j < len(d)-1 {
				buffer.WriteString(",")
			}

		}
		//end of object of the array
		buffer.WriteString("}")
		if i < len(content)-1 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(`]`)
	// final correction for empty keys
	str := buffer.String()
	str = strings.Replace(str, ":}", `:""}`, -1)
	str = strings.Replace(str, `:,"`, `:"","`, -1)
	rawMessage := json.RawMessage(str)
	fmt.Println(string(rawMessage))
	x, err := json.MarshalIndent(rawMessage, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	newFileName := filepath.Base(*path)
	newFileName = newFileName[0:len(newFileName)-len(filepath.Ext(newFileName))] + ".json"
	r := filepath.Dir(*path)
	return x, filepath.Join(r, newFileName)
}

// SaveFile Will Save the file, magic right?
func SaveFile(myFile []byte, path string) {
	if err := ioutil.WriteFile(path, myFile, os.FileMode(0644)); err != nil {
		panic(err)
	}
}
