package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var idList = []int64{
		31223123,
		34324243,
		4432543545,
		656546456,
	}
	file, _ := os.Create("data.csv")
	defer file.Close()
	writer := csv.NewWriter(file)

	for _, data := range idList {
		writer.Write([]string{strconv.FormatInt(data, 10)})
	}
	writer.Flush()

	csvReader := csv.NewReader(file)

	var buf bytes.Buffer
	for _, data := range idList {
		buf.WriteString(strconv.FormatInt(data, 10))
		buf.WriteByte('\n')
	}
	bufReader := bufio.NewReader(&buf)
	csvReader2 := csv.NewReader(&buf)
	fmt.Println(csvReader, bufReader, csvReader2)
	all, err := csvReader2.ReadAll()
	if err != nil {
		return
	}
	for _, d := range all {
		fmt.Println(d)
	}

}
