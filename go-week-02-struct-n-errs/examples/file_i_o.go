package goIdioms

import (
	"fmt"
	// "io/ioutil"
	"bufio"
	"os"
	"strings"
)
func GoFomatExamples() {
	fmt.Println("Go Format Examples")
	fmt.Println("What is actually going on here")
}

func GoFileIOExamples() {
	for _, fname := range os.Args[1:]{
		var lc, wc, cc int
		file, err := os.Open(fname)
		if err != nil{
			fmt.Fprint(os.Stderr, err)
			continue
		}
		// Lets mimick the wc in linux

		scan := bufio.NewScanner(file)

		for scan.Scan(){
			// scan the text
			s := scan.Text()

			// add the word count
			wc += len(strings.Fields(s))
			// add the char count
			cc+= len(s)
			// line count
			lc++
		}

 		fmt.Printf("%7d %7d %7d %s \n ", lc, wc, cc , fname)

		// data, err := ioutil.ReadAll(file)

		// if err != nil{
		// 	fmt.Fprint(os.Stderr, err)
		// 	continue
		// }

		// fmt.Printf("The document has %d bytes of data \n", len(data))
		
		// if _, err := io.Copy(os.Stdout, file); err != nil{
		// 	fmt.Fprint(os.Stderr, err)
		// 	continue
		// }

		file.Close()
	}
	// fmt.Println("Go File I/O Examples")
}

func FileIOExamples() {
GoFomatExamples()
GoFileIOExamples()
}