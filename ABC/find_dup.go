
package main

import (
"bufio"
"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func find_dup1(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func find_dup2(){
	counts := make(map[string]int)
	for _, filename := range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		if err != nil{
			fmt.Fprint(os.Stderr, "dup2: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n"){
			counts[line]++
		}
	}

	for line, n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main(){
	//find_dup1()
	find_dup2()
}