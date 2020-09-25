package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := "AAAAUUURRRTTT333"

	var buffer bytes.Buffer
	previousValue := ""
	count := 1
	for index, value := range data {
		if len(data) == 1{
			buffer.WriteString(fmt.Sprintf("%s%d", string(value), 1))
			break
		}

		if previousValue == "" {
			previousValue = string(value)
			continue
		}

		if string(value) == previousValue {
			count++
		} else {
			buffer.WriteString(fmt.Sprintf("%s%d", previousValue, count))
			count = 1
		}

		if index == len(data)-1 {
			buffer.WriteString(fmt.Sprintf("%s%d", previousValue, count))
		}

		previousValue = string(value)
	}

	fmt.Println(buffer.String())
}
