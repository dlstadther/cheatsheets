package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flagAdjustment := flag.String("type", "", "adjustment type (pre/post)")
	flag.Parse()

	switch *flagAdjustment {
	case "pre":
		preAdjustment()
	case "post":
		postAdjustment()
	default:
		break
	}
}

func preAdjustment() {
}

func postAdjustment() {
	basePath, _ := os.Getwd()
	bookPath := filepath.Join(basePath, "_book")

	err := filepath.Walk(bookPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) != ".html" {
			return nil
		}

		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		htmlString := string(buf)

		// ==== remove invalid lang tag for EPUB validation
		htmlString = strings.Replace(htmlString, ` lang="" xml:lang=""`, "", -1)

		// ==== inject github stars button
		buttonToFind := `</body>`
		buttonReplacement := `<div style="position: fixed; top: 10px; right: 30px; padding: 10px; background-color: rgba(255, 255, 255, 0.7);"><a class="github-button" href="https://github.com/dlstadther/cheatsheets" data-icon="octicon-star" data-size="large" data-show-count="true" aria-label="Star dlstadther/cheatsheets on GitHub">Star</a>&nbsp;<a class="github-button" href="https://github.com/novalagung" data-size="large" aria-label="Follow @dlstadther on GitHub">Follow @dlstadther</a><script async defer src="https://buttons.github.io/buttons.js"></script></div></body>`
		htmlString = strings.Replace(htmlString, buttonToFind, buttonReplacement, -1)

		// ==== inject github stars js script
		buttonScriptToFind := `</head>`
		buttonScriptReplacement := `<script async defer src="https://buttons.github.io/buttons.js"></script></head>`
		htmlString = strings.Replace(htmlString, buttonScriptToFind, buttonScriptReplacement, -1)

		// ==== update file
		err = ioutil.WriteFile(path, []byte(htmlString), info.Mode())
		if err != nil {
			return err
		}

		fmt.Println("  ==>", path)

		return nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	// ===== sitemap adjustment
	siteMapPath := filepath.Join(basePath, "_book", "sitemap.xml")
	buf, err := ioutil.ReadFile(siteMapPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	sitemapContent := strings.Replace(string(buf), `<changefreq>weekly</changefreq>`, `<changefreq>daily</changefreq>`, -1)
	err = ioutil.WriteFile(siteMapPath, []byte(sitemapContent), 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("  ==>", siteMapPath)
}
