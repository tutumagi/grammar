package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/tutumagi/grammar"
)

var (
	grammarFilePath = flag.String("grammar", "./demo.bnf", "bnf grammar file in json format")
	outputFilePath  = flag.String("output", "", "output file")
)

func main() {
	flag.Parse()

	bb, err := ioutil.ReadFile(*grammarFilePath)
	if err != nil {
		panic(err)
	}
	var output io.Writer = os.Stdout
	if outputFilePath != nil && *outputFilePath != "" {
		output, err = os.OpenFile(*outputFilePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot open file:%s err:%s", *outputFilePath, err)
			os.Exit(1)
		}
	}

	g := grammar.NewGrammar(string(bb))
	firstSet, followSet, predictTable := g.MakeFirstFollowPredict()
	fmt.Fprintln(output, "FirstSet:")
	grammar.PrintFirstSet(output, firstSet)
	fmt.Fprintln(output, "FollowSet:")
	grammar.PrintFollowSet(output, followSet)
	fmt.Fprintln(output, "PredictTable:")
	grammar.PrintPredictTable(output, predictTable)
}
