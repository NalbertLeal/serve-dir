package internal

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	CurrentDir string

	htmlReplacer = strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		// "&#34;" is shorter than "&quot;".
		`"`, "&#34;",
		// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
		"'", "&#39;",
	)
)

func GetCurrentDirectory() string {
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	return currDir
}

func directoryContain(dirPath string, fileName string) bool {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return false
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), fileName) {
			return true
		}
	}
	return false
}

func CreateDirHtml(dirPath string) string {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return ""
	}

	filesHTMLSlice := []string{}
	for _, file := range files {
		// name := htmlReplacer.Replace(file.Name())
		htmlElem := "<a href=" + file.Name() + ">" + file.Name() + "</a>"
		filesHTMLSlice = append(filesHTMLSlice, htmlElem)
	}
	sort.Strings(filesHTMLSlice)

	filesHTML := ""
	for _, file := range filesHTMLSlice {
		filesHTML += file + "<br/>"
	}

	return filesHTML
}
