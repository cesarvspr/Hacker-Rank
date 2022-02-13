package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

 type result struct {
	positive float64
	negative float64
	zero float64
 }
 
 func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(math.Round(num * output)) / output
}

func plusMinus(arr []int32) {
	// Write your code here
	var positives []int32
	var negatives []int32
	var zeros []int32
	//get positive values
	for i := range arr {
		
		if arr[i] > 0 {
			positives = append(positives, arr[i])
		}
		if arr[i] < 0 {
			negatives = append(negatives, arr[i])
		}
		if arr[i] == 0 {
			zeros = append(zeros, arr[i])
		}

		i++
	}
	var s result
	s.positive = toFixed(float64(len(positives))/float64(len(arr)), 6)
	s.negative = toFixed(float64(len(negatives))/float64(len(arr)), 6) 
	s.zero = toFixed(float64(len(zeros))/float64(len(arr)), 6)
	fmt.Println(s.positive)
	fmt.Println(s.negative)
	fmt.Println(s.zero)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
