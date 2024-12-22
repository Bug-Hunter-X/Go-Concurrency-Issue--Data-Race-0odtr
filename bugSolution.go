package main

import (

        "fmt"

        "sync"

)

var counter int
var mutex sync.Mutex

func incrementCounter() {

        for i := 0; i < 1000; i++ {

                mutex.Lock()

                counter++

                mutex.Unlock()

        }

}

func main() {

        var wg sync.WaitGroup

        for i := 0; i < 10; i++ {

                wg.Add(1)

                go func() {

                        defer wg.Done()

                        incrementCounter()

                }()

        }

        wg.Wait()

        fmt.Println("Final counter value:", counter)

} 