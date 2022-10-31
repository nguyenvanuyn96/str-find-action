package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

func listFiles(include string, exclude string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		if doesFileMatch(path, include, exclude) {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}

func doesFileMatch(path string, include string, exclude string) bool {
	if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
		includeRe := regexp.MustCompile(include)
		excludeRe := regexp.MustCompile(exclude)
		return includeRe.Match([]byte(path)) && !excludeRe.Match([]byte(path))
	}
	return false
}

func findStringInFile(path string, find string) ([]string, error) {
	read, readErr := ioutil.ReadFile(path)
	check(readErr)

	re := regexp.MustCompile(find)
	results := re.FindAllString(string(read), -1)

	if len(results) > 0 {
		return results, nil
	}

	return []string{}, nil
}

func main() {
	include := os.Getenv("INPUT_INCLUDE")
	exclude := os.Getenv("INPUT_EXCLUDE")
	find := os.Getenv("INPUT_FIND")

	files, filesErr := listFiles(include, exclude)
	check(filesErr)

	fileFoundCount := 0
	resultArray := []string{}

	for _, path := range files {
		foundStrs, findStringErr := findStringInFile(path, find)
		check(findStringErr)

		if len(foundStrs) > 0 {
			fileFoundCount++
			resultArray = append(resultArray, foundStrs...)
		}
	}
1
	fmt.Println(fmt.Sprintf('fileFoundCount=%d >> $GITHUB_OUTPUT', fileFoundCount))
	fmt.Println(fmt.Sprintf('resultArray=%v >> $GITHUB_OUTPUT', resultArray))
}
