package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Data struct {
	Name        string
	CountryCode string
	CountryFull string
	IsMobile    bool
}

func main() {
	fileNameReadFlag := flag.String("read", ".debug/DB24/data.csv", "reading file")
	fileNameSaveFlag := flag.String("out", ".debug/res/data_isp.csv", "saved file")
	flag.Parse()

	fileNameRead := *fileNameReadFlag
	data, err := parseDataFile(fileNameRead)
	if err != nil {
		log.Println("error parse data file:", err.Error())
		os.Exit(1)
	}
	log.Println("len data =", len(data))

	fileNameSave := *fileNameSaveFlag
	errWriteDataInFile := saveDataInFile(data, fileNameSave)
	if errWriteDataInFile != nil {
		log.Println("error save data in file:", errWriteDataInFile.Error())
		os.Exit(1)
	}
	log.Println("file created", fileNameSave)
}

func parseDataFile(fileName string) (map[string]*Data, error) {
	file, errReadFile := ioutil.ReadFile(fileName)
	if errReadFile != nil {
		return nil, fmt.Errorf(errReadFile.Error())
	}

	lines := strings.Split(string(file), "\n")

	res := make(map[string]*Data, 0)
	for _, line := range lines {
		dataLine := strings.Split(line, `","`)
		if len(dataLine) != 22 {
			continue
		}

		name := strings.Trim(dataLine[10], " \"\n\t")
		_, ok := res[name]
		if ok {
			continue
		}

		d := &Data{}
		d.Name = name
		d.CountryCode = strings.Trim(dataLine[2], " \"\n\t")
		d.CountryFull = strings.Trim(dataLine[3], " \"\n\t")

		m := strings.Trim(dataLine[21], " \"\n\t")
		if strings.Contains(m, "MOB") {
			d.IsMobile = true
		}
		res[name] = d
	}

	return res, nil
}

func saveDataInFile(data map[string]*Data, fileNameSave string) error {
	df, errCreateFile := os.Create(fileNameSave)
	if errCreateFile != nil {
		return errCreateFile
	}
	defer df.Close()

	for _, d := range data {
		resData := fmt.Sprintf("%s,%s,%s,%t\n", d.Name, d.CountryCode, d.CountryFull, d.IsMobile)
		_, errWrite := df.WriteString(resData)
		if errWrite != nil {
			return errWrite
		}
	}

	return nil
}

//====================================================
func parseDataCSV(fileName string) ([]*Data, error) {
	file, errReadFile := os.Open(fileName)
	if errReadFile != nil {
		return nil, fmt.Errorf(errReadFile.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 20

	res := make([]*Data, 0, 1000000)
	for {
		r, errRead := reader.Read()
		if errRead != nil {
			if errRead.Error() == "EOF" {
				break
			}
			return nil, errRead
		}

		d := &Data{}
		d.Name = r[10]
		d.CountryCode = r[2]
		d.CountryFull = r[3]
		if strings.Contains(r[21], "MOB") {
			d.IsMobile = true
		}

		res = append(res, d)
	}
	return res, nil
}
