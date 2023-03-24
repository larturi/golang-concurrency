package labs

import (
	"fmt"
	"os"
	"strings"
)

func StartHexa() {

	f, err := os.Create("mi_archivo.txt")
	if err != nil {
		panic(err)
	}

	outCh := make(chan string)
	doneWrite := make(chan struct{})

	go func() {
		for s := range outCh {
			_, err := f.WriteString(s)
			if err != nil {
				panic(err)
			}
		}
		doneWrite <- struct{}{}
	}()

	numGoRoutines := 10
	doneCh := make(chan struct{})

	final := 16777215

	for i := 0; i < final; i = i + (final / numGoRoutines) + 1 {
		paso := i + (final / numGoRoutines)

		if paso > final {
			paso = final
		}

		fmt.Printf("Ejecutando %d %d\n", i, paso)
		go calcNums(i, paso, outCh, doneCh)
	}

	doneNum := 0

	for doneNum < numGoRoutines {
		<-doneCh
		fmt.Println("Terminó uno")
		doneNum++
	}

	close(outCh)
	<-doneWrite
	fmt.Println("Listo!")

}

func calcNums(start, end int, resultCh chan string, doneCh chan struct{}) {

	var sBuilder strings.Builder

	for i := start; i <= end; i++ {
		fmt.Fprint(&sBuilder, fmt.Sprintf("%06x\n", i))
	}

	resultCh <- sBuilder.String()
	doneCh <- struct{}{}
}
