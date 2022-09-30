package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	//"log"//possible need
)

// helper will streamline error checks below
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

func getInstructionFormat(content string) string {
	data, error := os.Open(content)
	errorOpeningFile(error)
	var Tester = ""
	var Answer = ""
	var jar = ""
	var Holder = ""
	var Counter = 96 //memory location starting point
	//should increment by 4 bytes for each instruction
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
	//to keep track of the result of each loop in the scan
	//result is what gets returned
	var result = ""
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
			//printing directly to file, function takes %s "string" and &d digit value
			s := fmt.Sprintf("B format - %s", jar)

			//adding new string and linebreak
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}
		if opcode >= 1440 && opcode <= 1447 || opcode >= 1448 && opcode <= 1455 {
			//TODO: set jar
			//printing directly to file function takes %s "string" and &d digit value
			s := fmt.Sprintf("CB format - %s", jar)

			//adding new string and linebreak
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}
		if opcode >= 1684 && opcode <= 1687 || opcode >= 1940 && opcode <= 1943 {
			//TODO: set jar
			//printing directly to file function takes %s "string" and &d digit value
			s := fmt.Sprintf("IM format - %s", jar)

			//adding new string and linebreak
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}
		if opcode >= 1160 && opcode <= 1161 || opcode >= 1672 && opcode <= 1673 {
			//TODO: set jar
			//printing directly to file function takes %s "string" and &d digit value
			s := fmt.Sprintf("I format - %s", jar)

			//adding new string and linebreak
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}
		if opcode == 1986 || opcode == 1984 {
			//TODO: set jar
			//printing directly to file function takes %s "string" and &d digit value
			s := fmt.Sprintf("D format - %s", jar)

			//adding new string and linebreak
			result += s
			result += "\n"
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

			//printing directly to file function takes %s "string" and &d digit value
			s := fmt.Sprintf("%s %s %s %s %s %d %s %d %d %d", jar[0:11], jar[12:16], jar[17:22], jar[23:27], jar[28:32], Counter, Answer, register1, register2, register3)

			//adding new string and linebreak
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}
		Counter += 4
	}
	return result
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

	//TODO: Use os.Create to read/write rather than saving everything in memory and dumping it in os.WriteFile
	//create a file and check for errors
	//File, err := os.Create(*outputPath)
	//defer File.Close()

	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	instructionFormat := getInstructionFormat(*inputPath)
	fmt.Println("InputPath: ", *inputPath)
	os.WriteFile(*outputPath, []byte(instructionFormat), 0666)
	os.Exit(0)

}