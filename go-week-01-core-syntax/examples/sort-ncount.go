package coreConcepts

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func SortnCount() {
	// This function will read a file, count unique words, and sort them to find the top 3 most popular words
	scan := bufio.NewScanner(os.Stdin)
	word := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan(){
		word[scan.Text()]++
	}
	fmt.Println("unique words in file are:", len(word))

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range word{
		ss = append(ss, kv{k,v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _,s := range ss[:3]{
		fmt.Println(s.key, "appears", s.val, "times")
	}
}