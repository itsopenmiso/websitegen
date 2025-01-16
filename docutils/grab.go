package docutils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GrabJSONs() {}
func RemoveMarkdowns() {
	e := os.MkdirAll("rendered", 0644)
	if e != nil {
		panic(e)
	}

	_ = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Sprawdzenie, czy plik ma rozszerzenie .mdx
		if !info.IsDir() && filepath.Ext(info.Name()) == ".md" {
			fmt.Println("Removing: ", path)
			err := os.Remove(path)
			if err != nil {
				panic(err)
			}

		}

		return nil
	})

}
