package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func parse2Comp(s string) (int64, error) {
	sign := int64(1)
	if strings.HasPrefix(s, "1") {
		sign = -1
		b := make([]byte, len(s))
		for i := len(s) - 1; i >= 0; i-- {
			d := s[i]
			switch d {
			case '0':
				d = '1'
			case '1':
				d = '0'
			}
			b[i] = d
		}
		s = string(b)
	}
	i64, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0, err
	}
	if sign <= -1 {
		i64 = -i64 - 1
	}
	return i64, err
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
			s := fmt.Sprintf("%s\t%d\t%s", "1 11111 10110 11110 11111 11111 100111", Counter, "BREAK")
			result += s
			result += "\n"
			Holder = ""
			jar = ""
			Counter += 4
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
				register1, _ := parse2Comp(jar[7:32])
				s := fmt.Sprintf("%s \t%d\t%s\t#%d", jar[1:7]+" "+jar[7:32], Counter, Answer, register1)

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
				register1, _ := strconv.Atoi(jar[27:32])
				register1 = binaryToDecimal(register1)
				register2, _ := strconv.Atoi(jar[8:27])
				register2 = binaryToDecimal(register2)
				s := fmt.Sprintf("%s\t%d\t%s\tR%d, #%d", jar[0:8]+" "+jar[8:27]+" "+jar[27:32], Counter, Answer, register1, register2)

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
				register1, _ := strconv.Atoi(jar[27:32])
				register1 = binaryToDecimal(register1)
				register2, _ := strconv.Atoi(jar[11:27])
				register2 = binaryToDecimal(register2)
				register3, _ := strconv.Atoi(jar[9:11])
				register3 = binaryToDecimal(register3) * 16

				s := fmt.Sprintf("%s\t%d\t%s\tR%d, %d, LSL %d", jar[0:9]+" "+jar[9:11]+" "+jar[11:27]+" "+jar[27:32], Counter, Answer, register1, register2, register3)

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
				register1, _ := strconv.Atoi(jar[27:32])
				register2, _ := strconv.Atoi(jar[22:27])
				register3, _ := strconv.Atoi(jar[11:22])

				register1 = binaryToDecimal(register1)
				register2 = binaryToDecimal(register2)
				register3 = binaryToDecimal(register3)

				s := fmt.Sprintf("%s\t%d\t%s\tR%d, R%d, #%d", jar[1:11]+" "+jar[11:22]+" "+jar[22:27]+" "+jar[27:32], Counter, Answer, register1, register2, register3)

				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			//D format response
			if opcode == 1986 || opcode == 1984 {
				Answer = (checkForValue(opcode, ValidInstructions))
				register1, _ := strconv.Atoi(jar[27:32])
				register2, _ := strconv.Atoi(jar[22:27])
				register3, _ := strconv.Atoi(jar[11:20])

				register1 = binaryToDecimal(register1)
				register2 = binaryToDecimal(register2)
				register3 = binaryToDecimal(register3)

				s := fmt.Sprintf("%s\t%d\t%s\tR%d, [R%d, #%d]", jar[1:11]+" "+jar[11:20]+" "+jar[20:22]+" "+jar[22:27]+" "+jar[27:32], Counter, Answer, register1, register2, register3)

				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			//NOP response
			if opcode == 0 {
				Answer = (checkForValue(opcode, ValidInstructions))

				s := fmt.Sprintf("%s \t%d\t%s", jar[1:32], Counter, Answer)

				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			//LSL and LSR
			if opcode == 1690 || opcode == 1691 {
				Answer = (checkForValue(opcode, ValidInstructions))

				register1, _ := strconv.Atoi(jar[27:32])
				register2, _ := strconv.Atoi(jar[22:27])
				register3, _ := parse2Comp(jar[16:22])

				register1 = binaryToDecimal(register1)
				register2 = binaryToDecimal(register2)

				//printing directly to file function takes %s "string" and &d digit value
				s := fmt.Sprintf("%s %s %s %s %s \t%d\t%s\tR%d, R%d, #%d", jar[1:11], jar[11:16], jar[16:22], jar[22:27], jar[27:32], Counter, Answer, register1, register2, register3)
				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			//R format response
			if opcode == 1104 || opcode == 1112 || opcode == 1360 || opcode == 1624 || opcode == 1692 || opcode == 1872 {
				Answer = (checkForValue(opcode, ValidInstructions))
				register1, _ := strconv.Atoi(jar[27:32])
				register2, _ := strconv.Atoi(jar[22:27])
				register3, _ := strconv.Atoi(jar[11:16])

				register1 = binaryToDecimal(register1)
				register2 = binaryToDecimal(register2)
				register3 = binaryToDecimal(register3)

				//printing directly to file function takes %s "string" and &d digit value
				s := fmt.Sprintf("%s %s %s %s %s \t%d\t%s\tR%d, R%d, R%d", jar[1:11], jar[11:16], jar[16:22], jar[22:27], jar[27:32], Counter, Answer, register1, register2, register3)
				//adding new string and linebreak
				result += s
				result += "\n"
				Holder = ""
				jar = ""
			}
			Counter += 4
		//handles break case
		case true:
			twos, _ := parse2Comp(jar)
			s := ""
			s = fmt.Sprintf("%s\t%d\t%d", jar, Counter, twos)
			result += s
			result += "\n"
			Holder = ""
			jar = ""
			Counter += 4
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
