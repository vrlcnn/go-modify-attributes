package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

type core struct {
	file_name     string
	output        string
	transformType []string //camelcase, snakecase

	isWriteFile bool //if true write else print stdout
	reWrite     bool

	file_set *token.FileSet
}

func main() {

	if err := fucking_main(); err != nil {
		fmt.Println("There is something wrong bro! ,", err.Error())
		os.Exit(1)
	}
}

func fucking_main() error {

	b := parseCommandLine()

	ast_file, err := b.parseFile()
	if err != nil {
		return err
	}

	struct_list := b.findStructs(ast_file)

	b.findStructFieldList(struct_list)

	if err := b.astToPrint(ast_file); err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Println("Everthing translated. Go work!!")

	return nil
}

func (c *core) astToPrint(f *ast.File) error {
	var buff bytes.Buffer
	if err := format.Node(&buff, c.file_set, f); err != nil {
		return err
	}

	//if isWriteFile not set print stdout
	if c.isWriteFile == false {
		fmt.Printf("%s", buff)
		return nil
	}

	if err := ioutil.WriteFile(c.file_name, buff.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func parseCommandLine() *core {
	var (
		flagFileName = flag.String("file", "1", "hebele gubele")
		flagType     = flag.String("type", "", "Adds each tag with comma seperated. ex: json,xml, db etc.")
		flagReWrite  = flag.Bool("override", false, "Override tag")
		//flagStruct   = flag.String("struct", "", "If denifed for specific struct it should be seperated by comman ex: Game,GamePlayer. if all sturct required not define it.")
		flagIsWriteFile = flag.Bool("write-file", false, "if this writed output will be write given file else output printed to stdout.")
	)

	flag.Parse()

	b := &core{
		file_name:     *flagFileName,
		file_set:      token.NewFileSet(),
		reWrite:       *flagReWrite,
		isWriteFile:   *flagIsWriteFile,
		transformType: flagTypeSingleToSlice(*flagType),
	}

	return b
}

//split -flag tag1,tag2
func flagTypeSingleToSlice(ft string) []string {
	return strings.Split(ft, ",")
}

//parse given b.file to ast
func (c *core) parseFile() (*ast.File, error) {
	c.file_set = token.NewFileSet()

	return parser.ParseFile(c.file_set, c.file_name, nil, parser.ParseComments)
}

//find structs from ast
func (c *core) findStructs(nodeList ast.Node) *[]ast.StructType {
	structs := make([]ast.StructType, 0)

	getStructs := func(n ast.Node) bool {
		switch typ := n.(type) {
		case *ast.StructType:
			structs = append(structs, *typ)
		}

		return true
	}

	ast.Inspect(nodeList, getStructs)
	return &structs
}

//find on each field in struct and set them
func (c *core) findStructFieldList(structs *[]ast.StructType) {
	for _, strc := range *structs {
		for idx, _ := range strc.Fields.List {
			fieldList := (strc.Fields)
			x := fieldList.List[idx]

			c.setStructFieldTag(x)
		}
	}
}

//our job on struct tag field set with struct names field
func (c *core) setStructFieldTag(field *ast.Field) {
	if field.Tag == nil {
		name := ""

		//for embeded struct name will be nil
		if field.Names != nil {
			name = field.Names[0].Name

			field.Tag = &ast.BasicLit{
				Value: c.createTag(name),
				Kind:  token.STRING, ValuePos: field.Type.End() + 1,
			}
		}
	}
}

//it just return tag: value pair
func (c *core) createTag(name string) string {
	temp_tag_val := "`"
	for _, format_type := range c.transformType {
		temp_tag_val += fmt.Sprintf("%v: \"%v\" ", format_type, name)
	}
	temp_tag_val += "`"

	return temp_tag_val
}

/*
Rules:
VRLCNN => ["VRLCNN"]
FisekTavsan => ["Fisek", "Tavsan"]
MYMam => ["MY", "Mam"]
DEsignPArser => ["DEsign", "PArser"]
VERSION97 => ["VERSION", "97"]
*/
func (c *core) camelCaseToSnake(src string) {
	/*
		var u_type int
		var before_u_type int
		var runes [][]rune
	*/

}
