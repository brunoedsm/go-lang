package main

import (
    "fmt"
    "net"
    "sync"
)
func main() {

	var waiter sync.WaitGroup

    for i := 1; i <= 1024; i++ {

    	waiter.Add(1)
        
		go func(j int) {
        	defer waiter.Done()
            address := fmt.Sprintf("scanme.nmap.org:%d", j)
            conn, err := net.Dial("tcp", address)
            if err != nil {
                return
            }
            conn.Close()
            fmt.Printf("%d open\n", j)
        }(i)
		
    }
	waiter.Wait()
}