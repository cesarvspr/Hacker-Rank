package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'whatFlavors' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY cost
 *  2. INTEGER money
 */

func whatFlavors(cost []int32, money int32) {
	// Write your code here
	// var lenght int = 0
	var x, y, J int = 0, 0, 0
	var result [2]int
	var ok bool
	// fmt.Println("Money is: ", money)
	mp := make(map[int]int)
	for i := 0; i < len(cost); i++ {
		x = int(cost[i])
		y = int(money) - x
		fmt.Println("")
		fmt.Println("money", money, "-", "x", x, "is: ", y)

		//looking inside de the map
		J, ok = mp[y]
		fmt.Println("mp[", y, "] is: ", mp[y])

		if ok {
			fmt.Println("J is: ", J, "and i is: ", i)
			result[0] = J + 1
			result[1] = i + 1
			break
		}
		mp[x] = i

		fmt.Println("mp: ", mp)
	}

	fmt.Println(result[0], result[1])

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
	}
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
