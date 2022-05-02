package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

type PkgRawData struct {
	Name       string
	ImportPath string
	Imports    []string
}

func main() {
	//filePath := flag.String("filepath", "", "set path to file")
	flag.Parse()
	fmt.Println(flag.Arg(0))
	if flag.Arg(0) == "" {
		log.Fatalln("please set /path/to/your/file.txt")
	}

	pkgRawData := parseFile(flag.Arg(0))
	pkgClearData := clearFile(pkgRawData)
	generatePlantUML(pkgClearData)
}

func clearFile(rawData []*PkgRawData) []*PkgRawData {
	var mainPkgPath string
	for _, d := range rawData {
		if d.Name == "main" {
			mainPkgPath = d.ImportPath
		}
	}

	var clearPkgRawData []*PkgRawData
	for _, d := range rawData {
		if strings.Contains(d.ImportPath, mainPkgPath) {
			clearPkgRawData = append(clearPkgRawData, d)
		}
	}

	return clearPkgRawData
}

func generatePlantUML(pkgRawData []*PkgRawData) {
	f, err := os.OpenFile("test2.puml", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fileContent := "@startuml\n\n"
	for _, pkgRawDatum := range pkgRawData {
		fileContent += "namespace " + pkgRawDatum.Name + "{}\n"
	}
	fileContent += "\n"
	//"logger" <|-- "main"
	for _, pkgRawDatum1 := range pkgRawData {
		for _, pkgRawDatum2 := range pkgRawData {
			if pkgRawDatum1.ImportPath == pkgRawDatum2.ImportPath {
				continue
			}
			for _, i := range pkgRawDatum1.Imports {
				if i == pkgRawDatum2.ImportPath {
					fileContent += `"` + pkgRawDatum1.Name + `" --> "` + pkgRawDatum2.Name + `"` + "\n"
				}
			}

		}

	}

	fileContent += "\n@enduml"

	_, err = f.Write([]byte(fileContent))
	if err != nil {
		log.Fatal(err)
	}

}

func parseFile(fileName string) []*PkgRawData {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var pkgRawData []*PkgRawData

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sLine := strings.Split(line, "__")
		pkgName := sLine[0]
		pkgImportPath := sLine[1]
		pkgImportsRaw1 := strings.ReplaceAll(sLine[2], "[", "")
		pkgImportsRaw2 := strings.ReplaceAll(pkgImportsRaw1, "]", "")
		pkgImports := strings.Split(pkgImportsRaw2, " ")

		pkgRawData = append(pkgRawData, &PkgRawData{
			Name:       pkgName,
			ImportPath: pkgImportPath,
			Imports:    pkgImports,
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return pkgRawData
}

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !strings.Contains(s, "vendor") {
		if d.IsDir() {
			println(d.Name() + "___" + s)
		}
	}
	return nil
}
