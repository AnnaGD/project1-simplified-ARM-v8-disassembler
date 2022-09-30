package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func errorOpeningFile(e error) {
	if e != nil {
		panic(e)
	}
}

func binaryToDecimal(binaryNum int) int {
	var rem int
	index := 0
	DecimalNum := 0
	for binaryNum != 0 {
		rem = binaryNum % 10
		binaryNum = binaryNum / 10
		DecimalNum = DecimalNum + rem*int(math.Pow(2, float64(index)))
		index++
	}
	return DecimalNum
}

func checkForValue(val int, instructions map[int]string) string {
	for key, value := range instructions {
		if key == (val) {
			return (value)
		}
	}
	return "cannot find valid instruction"
}

func getInstructionFormat(content string) {
	data, error := os.Open(content)
	errorOpeningFile(error)
	var Tester = ""
	var Answer = ""
	var jar = ""
	var Holder = ""
	var Counter = 96
	ValidInstructions := map[int]string{
		1104: "AND",
		1112: "ADD",
		1360: "ORR",
		1624: "SUB",
		1690: "LSR",
		1691: "LSL",
		1692: "ASR",
		1872: "EOR",
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {

		Tester = fileScanner.Text()
		chars := []rune(Tester)
		Container := string(chars)
		opcode, _ := strconv.Atoi(Holder)

		for i := 0; i < 32; i++ {
			jar += string(chars[i])
		}

		for i := 0; i <= 11; i++ {
			Holder += string(chars[i])
			if len(Holder) == 11 {
				Container = string(Holder)
				opcode, _ = strconv.Atoi(Container)
				opcode = binaryToDecimal(opcode)
			}
		}

		if opcode >= 160 && opcode <= 191 {
			jar = string(jar[5] + ' ')
			fmt.Println("B format - ", jar)
			Holder = ""
			jar = ""
		}
		if opcode >= 1440 && opcode <= 1447 || opcode >= 1448 && opcode <= 1455 {
			fmt.Println("CB format - ", jar)
			Holder = ""
			jar = ""
		}
		if opcode >= 1684 && opcode <= 1687 || opcode >= 1940 && opcode <= 1943 {
			fmt.Println("IM format - ", jar)
			Holder = ""
			jar = ""
		}
		if opcode >= 1160 && opcode <= 1161 || opcode >= 1672 && opcode <= 1673 {
			fmt.Println("I format - ", jar)
			Holder = ""
			jar = ""
		}
		if opcode == 1986 || opcode == 1984 {
			fmt.Println("D format - ", jar)
			Holder = ""
			jar = ""
		}
		//Rformat response
		if opcode == 1104 || opcode == 1112 || opcode == 1360 || opcode == 1624 ||
			opcode == 1690 || opcode == 1691 || opcode == 1692 || opcode == 1872 {
			Answer = (checkForValue(opcode, ValidInstructions))
			register1, _ := strconv.Atoi(jar[23:27])
			register2, _ := strconv.Atoi(jar[12:16])
			register3, _ := strconv.Atoi(jar[28:32])

			register1 = binaryToDecimal(register1)
			register2 = binaryToDecimal(register2)
			register3 = binaryToDecimal(register3)

			fmt.Println(jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer, register1, register2, register3)
			Holder = ""
			jar = ""
		}
		Counter += 4
	}
}

func main() {
	var Path = "src/temp/data.txt"
	File, err := os.Create("team1_out_dis.txt")

	if err != nil {
		fmt.Println(err)
		File.Close()
		return
	}

	getInstructionFormat(Path)
}