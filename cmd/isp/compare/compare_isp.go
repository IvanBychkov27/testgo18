package main

import (
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
	fileNameNew := ".debug/res/data_isp.csv"
	dataNew, err := parseDataFile_New(fileNameNew)
	if err != nil {
		log.Println("error parse data new file:", err.Error())
		os.Exit(1)
	}
	log.Println("len data new =", len(dataNew))

	fileNameOld := ".debug/real/isp_old.csv"
	dataOld, errOld := parseDataFile_Old(fileNameOld)
	if errOld != nil {
		log.Println("error parse data old file:", errOld.Error())
		os.Exit(1)
	}
	log.Println("len data old =", len(dataOld))
	log.Println("diff =", len(dataOld)-len(dataNew))

	countNew := 0
	countNewIsMobile := 0
	for name, n := range dataNew {
		o, ok := dataOld[name]
		if !ok {
			countNew++
			continue
		}
		if n.IsMobile != o.IsMobile {
			countNewIsMobile++
		}
	}
	log.Println("new data:", countNew)
	log.Println("countNewIsMobile:", countNewIsMobile)

	countOld := 0
	countOldIsMobile := 0
	for name, o := range dataOld {
		n, ok := dataNew[name]
		if !ok {
			countOld++
			continue
		}
		if n.IsMobile != o.IsMobile {
			countOldIsMobile++
		}
	}
	log.Println("old data:", countOld)
	log.Println("countOldIsMobile:", countOldIsMobile)

}

func parseDataFile_New(fileName string) (map[string]*Data, error) {
	file, errReadFile := ioutil.ReadFile(fileName)
	if errReadFile != nil {
		return nil, fmt.Errorf(errReadFile.Error())
	}

	lines := strings.Split(string(file), "\n")

	//count := 0
	res := make(map[string]*Data, 0)
	for _, line := range lines {
		dataLine := strings.Split(line, `,`)
		if len(dataLine) != 4 {
			//count++
			continue
		}

		name := strings.Trim(dataLine[0], " \"\n\t")
		_, ok := res[name]
		if ok {
			continue
		}

		d := &Data{}
		d.Name = name
		d.CountryCode = strings.Trim(dataLine[1], " \"\n\t")
		d.CountryFull = strings.Trim(dataLine[2], " \"\n\t")

		m := strings.Trim(dataLine[3], " \"\n\t")
		if strings.Contains(m, "true") {
			d.IsMobile = true
		}
		res[name] = d
	}
	//log.Println("new count:", count)
	return res, nil
}

func parseDataFile_Old(fileName string) (map[string]*Data, error) {
	file, errReadFile := ioutil.ReadFile(fileName)
	if errReadFile != nil {
		return nil, fmt.Errorf(errReadFile.Error())
	}

	//df, errCreateFile := os.Create(".debug/real/temp_isp_old.csv")
	//if errCreateFile != nil {
	//	return nil, errCreateFile
	//}
	//defer df.Close()

	lines := strings.Split(string(file), "\n")

	//count := 0
	res := make(map[string]*Data, 0)
	for i, line := range lines {
		if i == 0 {
			continue
		}

		dataLine := strings.Split(line, `,`)
		if len(dataLine) < 5 {
			//count++
			//_, errWrite := df.WriteString(line + "\n")
			//if errWrite != nil {
			//	return nil, errWrite
			//}
			continue
		}

		name := strings.Trim(dataLine[1], " \"\n\t")
		_, ok := res[name]
		if ok {
			continue
		}

		d := &Data{}
		d.Name = name
		d.CountryCode = strings.Trim(dataLine[2], " \"\n\t")
		d.CountryFull = strings.Trim(dataLine[3], " \"\n\t")

		m := strings.Trim(dataLine[4], " \"\n\t")
		if strings.Contains(m, "true") {
			d.IsMobile = true
		}
		res[name] = d
	}

	//log.Println("old count:", count)
	return res, nil
}
