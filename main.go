package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	jobs := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for domain := range jobs {
				_, err := net.LookupHost(domain)
				if err != nil {
					continue
				}

				fmt.Println(domain)
			}
		}()
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := sc.Text()
		jobs <- domain
	}
	close(jobs)

	wg.Wait()
}
