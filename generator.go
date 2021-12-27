package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

func main() {
	var files []string
	root := "java-slack-sdk/json-logs/samples/api/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path != root {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		json, _ := ioutil.ReadFile(file)
		apiMethodName := strings.ReplaceAll(filepath.Base(file), ".json", "")
		packageName := strings.ReplaceAll(apiMethodName, ".", "_")
		name := strcase.ToCamel(packageName)
		cmd := exec.Command(
			"bash", "-c",
			fmt.Sprintf("quicktype --all-properties-optional --package %s -t %s", packageName, name))
		stdin, cmdError := cmd.StdinPipe()
		if cmdError != nil {
			log.Fatal(cmdError)
			return
		}
		stdin.Write(json)
		stdin.Close()
		stdout, _ := cmd.StdoutPipe()
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(fmt.Sprintf("Generating %s ...", apiMethodName))
			os.MkdirAll(fmt.Sprintf("webapi/%s", packageName), os.ModePerm)
			ioutil.WriteFile(
				fmt.Sprintf("webapi/%s/endpoint.go", packageName),
				[]byte(fmt.Sprintf("package %s\nconst ApiMethod string = \"%s\"",
					packageName,
					apiMethodName,
				)),
				0644)
			body, _ := ioutil.ReadAll(stdout)
			ioutil.WriteFile(
				fmt.Sprintf("webapi/%s/response.go", packageName),
				body,
				0644)
		}
	}
}
