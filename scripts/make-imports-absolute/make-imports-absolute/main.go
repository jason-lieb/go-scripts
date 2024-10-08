package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide absolute import path")
		os.Exit(1)
	}

	if err := filepath.Walk("./", visit); err != nil {
		panic(err)
	}
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() || !(strings.HasSuffix(path, ".ts") || strings.HasSuffix(path, ".tsx")) {
		return nil
	}

	if err := changeRelativeImportsToAbsolute(path); err != nil {
		return err
	}

	return nil
}

func changeRelativeImportsToAbsolute(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	modifiedContent := modifyImports(filePath, string(content))
	fmt.Println("modifiedContent:\n", modifiedContent)

	// file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	// if _, err := file.WriteString(modifiedContent); err != nil {
	// 	return err
	// }

	return nil
}

func modifyImports(filePath string, content string) string {
	absoluteImportPath := os.Args[1]
	regex := regexp.MustCompile(`from '(\..*?)'`)
	matches := regex.FindAllStringSubmatch(content, -1)

	pathMap := cleanPaths(matches)

	modifiedContent := regex.ReplaceAllStringFunc(content, func(match string) string {
		relativePath := regex.FindStringSubmatch(match)[1]
		cleanedPath := pathMap[relativePath]
		return `from '` + absoluteImportPath + `/` + cleanedPath + `';`
	})

	return modifiedContent
}

func cleanPaths(matches [][]string) map[string]string {
	pathMap := make(map[string]string)

	for _, match := range matches {
		if len(match) > 1 {
			cleanedPath := path.Clean(match[1])
			for strings.HasPrefix(cleanedPath, "../") {
				cleanedPath = strings.TrimPrefix(cleanedPath, "../")
			}
			pathMap[match[1]] = cleanedPath
		}
	}
	return pathMap
}
