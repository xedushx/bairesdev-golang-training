package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var s = []string{"James", "Robert", "John", "Michael", "William", "David", "Richard",
	"Joseph", "Thomas", "Charles", "Barbara", "Susan", "Jessica", "Sarah", "Karen",
	"Donna", "Michelle", "Paul"}

func main() {
	ch := make(chan string)
	var mut sync.Mutex
	counter := 0

	searchSb := "ara"

	for w := 0; w < 5; w++ {
		go func(ch chan string, mut *sync.Mutex) {
			for {
				mut.Lock()
				i := counter
				counter++
				mut.Unlock()

				if counter > len(s) {
					return
				}

				currentString := s[i]

				if strings.Contains(currentString, searchSb) {
					ch <- currentString
				} else {
					ch <- "-"
				}
			}
		}(ch, &mut)
	}

	t := time.Now()
	for i := 0; i < len(s); i++ {
		result := <-ch
		if result != "-" {
			fmt.Printf("Time: %v Line: %v Result: %v\n", t, 1, result)
		}
	}

}
