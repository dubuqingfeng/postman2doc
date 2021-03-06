package postman

import (
	"fmt"
	"os"
	"github.com/fatih/color"
	"github.com/dubuqingfeng/postman2doc/utils"
	"html/template"
)

type Param struct {
	IsDirectory bool
	Template    string
	Collection  string
	Output      string
}

func Handle(param Param) {
	// Check the parameters.
	if param.Collection == "" {
		fmt.Println(color.RedString("Error: Collection File not exist."))
		os.Exit(0)
	}

	if !utils.PathOrFileExist(param.Collection, false) {
		fmt.Println(color.RedString("Error: Collection File not found."))
		os.Exit(0)
	}

	if param.Output != "" && !utils.PathOrFileExist(param.Output, param.IsDirectory) {
		fmt.Println(color.RedString("Error: Output File or Directory not found."))
		os.Exit(0)
	}

	if param.Template != "" && !utils.PathOrFileExist(param.Template, false) {
		fmt.Println(color.RedString("Error: Template not found."))
		os.Exit(0)
	}
	// parse collection and into Collection
	postmanCollection, err := getPostmanCollection(param.Collection)
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}
	//t := template.New("some template")
	t, err := template.ParseFiles(param.Template)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	f, err := os.OpenFile(param.Output, os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	err = t.Execute(f, postmanCollection)
	if err != nil {
		fmt.Println(err)
	}
}
