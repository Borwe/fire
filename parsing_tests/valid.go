package parsingtests

import (
	"fmt"
	"log"
	"os"

	"go/ast"
	"go/parser"
	"go/token"
)

type Wembe struct {
	l1 string
	l2 int64
}

func trollll(){}

func yoyo(x Wembe){}

func yoyo4(x Wembe, y float64){}
func yoyo5(x *Wembe, y float64){}

func main(){
	dir, err := os.Getwd()
	if err!=nil {
		log.Fatalln(err)
	}
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir,nil, parser.SpuriousErrors)
	if err != nil{
		log.Fatalf("Failed to parse with error: %s\n",err)
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Files{
			ast.Print(fset,file.Scope)
		}
	} 
	fmt.Println("\n\n🔥FIRE, yet another hot reloader🔥")
}
