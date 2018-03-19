package postman

import (
	"fmt"
	"os"
	"github.com/fatih/color"
	"github.com/dubuqingfeng/postman2doc/utils"
)

type Param struct {
	IsDirectory bool
	Template    string
	Collection  string
	Output      string
}

func init() {

}

func Handle(param Param) {
	// Check the parameters.
	CheckParam(param)
	fmt.Println(param.Collection)
	fmt.Println(param.Output)
}
func CheckParam(param Param) {
	if param.Collection == "" {
		fmt.Println(color.RedString("Error: Collection File not exist."))
		os.Exit(0)
	}

	if !utils.PathOrFileExist(param.Collection, false) {
		fmt.Println(color.RedString("Error: Collection File not found."))
		os.Exit(0)
	}

	if param.Output != "" && !utils.PathOrFileExist(param.Output, true) {
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

	//
	fmt.Println(postmanCollection)
}
