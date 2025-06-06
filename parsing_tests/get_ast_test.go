package parsingtests_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/Borwe/fire/parsing"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/packages"
)

func TestGettingAstFromValidGoFileAndTurningBackToFile(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	assert.True(t, ok)
	dir := filepath.Dir(file) + "/"

	fileSet := token.NewFileSet()

	var filesAst []*ast.File
	var errs []packages.Error
	var err error

	fmt.Println("DIR:", dir)

	filesAst, errs, err = parsing.ParsePackage(dir, fileSet)
	assert.Empty(t, errs)
	assert.NotEmpty(t, filesAst)
	assert.Nil(t, err)

	//get ast file
	var thisFileAst *ast.File
	for _, fAst := range filesAst {
		fName := fileSet.Position(fAst.Package).Filename
		fmt.Println("fname:", fName)
		if fName == file {
			thisFileAst = fAst
		}
	}
	assert.NotNil(t, thisFileAst)

	//read file
	fileData, err := os.ReadFile(file)
	assert.Nil(t, err)
	//turn ast to file
	astFileData, err := parsing.ToBytes(thisFileAst, fileSet)
	assert.Nil(t, err)

	//get asts of file
	origFileDataAst, err := parser.ParseFile(fileSet, "test.go", fileData, parser.SpuriousErrors)
	assert.Nil(t, err)

	newDataAst, err := parser.ParseFile(fileSet, "test.go", astFileData, parser.SpuriousErrors)
	assert.Nil(t, err)

	//check that the two asts match
	assert.True(t, compareAstsIfMatch(origFileDataAst, newDataAst))
}

func compareAstsIfMatch(a *ast.File, b *ast.File) bool {
	astCmpChannel1 := make(chan *ast.Node, 1)
	astCmpChannel2 := make(chan *ast.Node, 1)
	continueChannel1 := make(chan bool, 1)
	continueChannel2 := make(chan bool, 1)
	doneParsing := make(chan bool, 1)
	parseFunc := func(astf *ast.File,
		astCmp chan<- *ast.Node, continueChan <-chan bool, done chan<- bool) {
		ast.Inspect(astf, func(node ast.Node) bool{
			fmt.Println("Send")
			astCmp<- &node
			return <-continueChan
		})
		fmt.Println("DONE")
		close(astCmp)
		done <- true
	}
	go parseFunc(a, astCmpChannel1, continueChannel1, doneParsing)
	go parseFunc(b, astCmpChannel2, continueChannel2, doneParsing)

	done := 0
	for done != 2{
		n1 := <- astCmpChannel1
		n2 := <- astCmpChannel2

		fmt.Println("GET")

		if n1 == nil || n2 == nil &&
			(reflect.TypeOf(n1) != reflect.TypeOf(n2)) {
			fmt.Println("FAILED ON NIL OR NOT SAME TYPE",n1,"\n\nAND:",n2)
			continueChannel1<- false
			continueChannel2<- false
			close(continueChannel1)
			close(continueChannel2)
			return false
		}else{
			continueChannel1<- true
			continueChannel2<- true
		}

		trydone:

		select {
		case <-doneParsing:
			fmt.Println("DONE+++")
			done+=1
			goto trydone
		case <- time.After(3* time.Millisecond):
			continue
		}
	}
	return true
}
