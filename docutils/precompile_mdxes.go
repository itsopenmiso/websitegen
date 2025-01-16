package docutils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrecompileMarkdowns() {
	// Check if there is a website folder in current catalog

	// Create websitegen

	// Create websitegen/stage1

	precompileFolder("commands")
	precompileFolder("docs")
	precompileFolder("plugins")

}

func precompileFolder(dir string) {
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Sprawdzenie, czy plik ma rozszerzenie .mdx
		if !info.IsDir() && filepath.Ext(info.Name()) == ".mdx" {
			precompileFile(path)
		}
		return nil
	})
}

func precompileFile(filename string) {
	fmt.Println("Precompiling", filename)
	newFilename := strings.Replace("../../output/"+filename, ".mdx", ".md", -1)
	os.MkdirAll(filepath.Dir(newFilename), 0755)
	old, _ := os.Open(filename)
	defer old.Close()
	newFile, u := os.OpenFile(newFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if u != nil {
		panic(u)
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(old)
	writer := bufio.NewWriter(newFile)
	for scanner.Scan() {
		line := scanner.Text()
		// If line begins with @include - open the partials file and input its contents
		if strings.HasPrefix(line, "@include") {
			includePath := strings.Replace(line, "@include", "", -1)
			purePath := "partials/" + strings.Replace(strings.TrimSpace(includePath), "\"", "", -1)
			includefile, rde := os.ReadFile(purePath)
			if rde != nil {
				panic(rde)
			}
			writer.WriteString(purifyLine(string(includefile)))

		} else {
			_, e := writer.WriteString(purifyLine(line) + "\n")
			if e != nil {
				panic(e)
			}
		}
	}
	err := writer.Flush()
	if err != nil {
		panic(err)
	}

}

func purifyLine(line string) string {
	return strings.Replace(line, "(/waypoint/", "(../", -1)
}
