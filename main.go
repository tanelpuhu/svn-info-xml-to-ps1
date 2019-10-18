package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const svnInfoXMLToPS1Version string = "0.0.2"

var (
	flagVersion bool
	// ErrInvalidInput is returned as error if input is invalid
	ErrInvalidInput = errors.New("Invalid input?")
)

type typeXMLInfo struct {
	URL         string `xml:"entry>repository>root"`
	RelativeURL string `xml:"entry>relative-url"`
}

func getStdin() []byte {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return data
}

func parseXMLInfo(content []byte) (string, string) {
	res := typeXMLInfo{}
	err := xml.Unmarshal([]byte(content), &res)
	if err != nil {
		return "", ""
	}
	sp := strings.Split(res.URL, "/")
	res.URL = sp[len(sp)-1]
	res.RelativeURL = res.RelativeURL + "/"
	return res.URL, res.RelativeURL
}

func getPS1String(input []byte, cwd string) (string, error) {
	var (
		result   []string
		name     string
		location string
	)
	repo, relurl := parseXMLInfo(input)
	if repo != "" {
		if !strings.HasSuffix(cwd, "/") {
			cwd += "/"
		}
		if strings.HasPrefix(relurl, "^/trunk/") {
			location = "trunk"
		} else {
			re1, _ := regexp.Compile(`^\^/(branches|tags|releases)/(.*?\/)`)
			resultSlice := re1.FindStringSubmatch(relurl)
			if len(resultSlice) == 3 {
				location = resultSlice[1]
				name = strings.Trim(resultSlice[2], "/")
			}
		}
		for _, value := range []string{repo, location, name} {
			if value != "" {
				if strings.Index(cwd, "/"+value+"/") == -1 {
					result = append(result, value)
				}
			}
		}
		if len(result) == 0 {
			result = append(result, repo)
		}
		return ":" + strings.Join(result, ":"), nil
	}
	return "", ErrInvalidInput
}

func main() {
	flag.BoolVar(&flagVersion, "V", false, "Print version")
	flag.Parse()
	if flagVersion {
		fmt.Printf("svn-info-xml-to-ps1 %v\n", svnInfoXMLToPS1Version)
		os.Exit(0)
	}
	cwd, _ := os.Getwd()
	cwd, _ = filepath.Abs(cwd)
	result, err := getPS1String(getStdin(), cwd)
	if err == nil {
		fmt.Println(result)
	}
}
