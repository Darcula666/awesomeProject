// 从控制台读取输入:
package main

import (
	"fmt"
	"net/http"
	"time"
)

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}
func main() {
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	ch := make(chan string)
	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(1000)
	if err != nil {
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}
	return employee, nil
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

func updateName(name *string) {
	*name = "David"
}
