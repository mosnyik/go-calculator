package coreConcepts

import (
	"fmt"
	"strings"	
	"os"
	"bufio"
)

func SearchReplace() {

	if len(os.Args) < 3{
		fmt.Fprintln(os.Stderr, "Usage: search-replace <search> <replace>")
		os.Exit(-1)
	}
	
	search,replace := os.Args[1], os.Args[2]

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan(){
		s := strings.Split(scan.Text(), search)
		t := strings.Join(s, replace)

		fmt.Println(t)

		// lets say I want to collect data from the user
		
		
	}
}