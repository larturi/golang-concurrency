package labs

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Start() {
	USE_GO_RUTINES := true
	start := time.Now()

	if USE_GO_RUTINES {
		// Ejecución concurrente
		for i := 0; i < 100; i++ {
			go getTodoFromService(i)
		}
	} else {
		// Ejecución secuencial:
		for i := 0; i < 100; i++ {
			getTodoFromService(i)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Tiempo de ejecución: %.7f segundos", elapsed.Seconds())

	var s string
	fmt.Scan(&s)
}

func getTodoFromService(id int) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(id))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("Status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
