package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

// helper will streamline error checks below
func errorOpeningFile(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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

func getInstructionFormat(content string) string {
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
		2038: "BREAK",
		1984: "STUR",
		1986: "LDUR",
		0:    "NOP",
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)
	//to keep track of the result of each loop in the scan
	//result is what gets returned
	var result = ""
	var Switcher = false
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
				Container = Holder
				opcode, _ = strconv.Atoi(Container)
				opcode = binaryToDecimal(opcode)
			}
		}
		//fmt.Println("OpCode: ", opcode)

		if opcode == 2038 {
			Answer = checkForValue(opcode, ValidInstructions)
			s := fmt.Sprintf("%s\t%s", "1 11111 10110 11110 11111 11111 100111", "BREAK")
			result += s
			result += "\n"
			Holder = ""
			jar = ""
			Switcher = true
			continue
		}

		switch Switcher {
		case false:
			//B format
			var num1 = 160
			var num2 = 191
			if opcode >= num1 && opcode <= num2 {
				Answer = "B"
				s := fmt.Sprintf("%s\t%d\t%s", jar[0:6]+" "+jar[7:32], Counter, Answer)

				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			//CB format
			if opcode >= 1440 && opcode <= 1447 || opcode >= 1448 && opcode <= 1455 {
				Answer = "CBZ"
				if opcode >= 1448 && opcode <= 1455 {
					Answer = "CBNZ"
				}
				s := fmt.Sprintf("%s\t%d\t%s", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer)

				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			//IM format
			if opcode >= 1684 && opcode <= 1687 || opcode >= 1940 && opcode <= 1943 {
				Answer = "MOVK"
				if opcode >= 1684 && opcode <= 1687 {
					Answer = "MOVZ"
				}
				register1, _ := strconv.Atoi(jar[28:32])

				register1 = binaryToDecimal(register1)

				s := fmt.Sprintf("%s\t%d\t%s\tR%d", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer, +register1)

				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			//I format response
			if opcode >= 1160 && opcode <= 1161 || opcode >= 1672 && opcode <= 1673 {
				Answer = "ADDI"
				if opcode >= 1672 && opcode <= 1673 {
					Answer = "SUBI"
				}
				register1, _ := strconv.Atoi(jar[23:27])
				register2, _ := strconv.Atoi(jar[28:32])

				register1 = binaryToDecimal(register1)
				register2 = binaryToDecimal(register2)

				s := fmt.Sprintf("%s\t%d\t%s\tR%d R%d", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer, register1, register2)

				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			//D format response
			if opcode == 1986 || opcode == 1984 {
				Answer = (checkForValue(opcode, ValidInstructions))
				register1, _ := strconv.Atoi(jar[23:27])
				register2, _ := strconv.Atoi(jar[28:32])

				register1 = binaryToDecimal(register1)
				register2 = binaryToDecimal(register2)

				s := fmt.Sprintf("%s\t%d\t%s\tR%d R%d", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer, register1, register2)

				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			//NOP response
			if opcode == 0 {
				Answer = (checkForValue(opcode, ValidInstructions))

				s := fmt.Sprintf("%s\t%d\t%s", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer)

				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			//R format response
			if opcode == 1104 || opcode == 1112 || opcode == 1360 || opcode == 1624 ||
				opcode == 1690 || opcode == 1691 || opcode == 1692 || opcode == 1872 {
				Answer = (checkForValue(opcode, ValidInstructions))
				register1, _ := strconv.Atoi(jar[23:27])
				register2, _ := strconv.Atoi(jar[12:16])
				register3, _ := strconv.Atoi(jar[28:32])

				register1 = binaryToDecimal(register1)
				register2 = binaryToDecimal(register2)
				register3 = binaryToDecimal(register3)

				//printing directly to file function takes %s "string" and &d digit value
				s := fmt.Sprintf("%s %s %s %s %s\t%d\t%s\tR%d R%d R%d", jar[0:11], jar[12:16], jar[17:22], jar[23:27], jar[28:32], Counter, Answer, register1, register2, register3)
				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			Counter += 4
		//handles break case
		case true:
			TwosCompDecimal, _ := strconv.ParseInt(string(jar[1:]), 2, 64)
			s := ""
			if string(jar[0]) == "1" {
				s = fmt.Sprintf("%s %s -%d", jar, "\t", TwosCompDecimal)
			} else {
				s = fmt.Sprintf("%s %s %d", jar, "\t", TwosCompDecimal)
			}
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}
	}
	return result
}

func writeToFile(path string, info string) {

	os.WriteFile(path, []byte(info), 0666)

}

func main() {

	//flags for input/output
	inputPath := flag.String("i", "", "The path to input file")
	outputPath := flag.String("o", "", "The path to output file")
	flag.Parse()

	if *inputPath == "" {
		fmt.Println("Please be sure to enter both an -i and -o for path to input/output file")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *outputPath == "" {
		fmt.Println("Please also enter -o for path to output file")
		flag.PrintDefaults()
		os.Exit(1)
	}

	os.Remove(*outputPath)
	formattedInstructions := getInstructionFormat(*inputPath)
	writeToFile(*outputPath, formattedInstructions)
}
