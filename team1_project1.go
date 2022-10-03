package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Instruction For project 2
type Instruction struct {
	typeofInstruction string
	rawInstruction    string
	lineValue         uint64
	programCnt        int
	opcode            uint64
	op                string
	rd                uint8
	rn                uint8
	rm                uint8
	im                string
	rt                uint8
}

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

		//B format
		if opcode >= 160 && opcode <= 191 {
			Answer = (checkForValue(opcode, ValidInstructions))
			// opcode 6 bits, offset(w) 26 bits
			// using 2s compliment 26 bits converted (length 8) = 0001 1010
			jar = string(jar[5] + ' ')
			//printing directly to file, function takes %s "string" and &d digit value
			//s := fmt.Sprintf("B format - %s", jar)
			s := fmt.Sprintf("B format - ", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer)

			//adding new string and linebreak
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}
		//CB format
		if opcode >= 1440 && opcode <= 1447 || opcode >= 1448 && opcode <= 1455 {
			Answer = (checkForValue(opcode, ValidInstructions))
			// opcode 8 bits, offset(w) 19 bits, conditional 5 bits
			jar = string(jar[0] + ' ')
			s := fmt.Sprintf("CB format - ", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer)

			//adding new string and linebreak
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}

		//IM format
		if opcode >= 1684 && opcode <= 1687 || opcode >= 1940 && opcode <= 1943 {
			Answer = (checkForValue(opcode, ValidInstructions))
			register1, _ := strconv.Atoi(jar[28:32])

			register1 = binaryToDecimal(register1)

			s := fmt.Sprintf("IM format - ", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer, register1)

			//adding new string and linebreak
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}
		//I format response
		if opcode >= 1160 && opcode <= 1161 || opcode >= 1672 && opcode <= 1673 {
			Answer = (checkForValue(opcode, ValidInstructions))
			register1, _ := strconv.Atoi(jar[23:27])
			register2, _ := strconv.Atoi(jar[28:32])

			register1 = binaryToDecimal(register1)
			register2 = binaryToDecimal(register2)

			s := fmt.Sprintf("I format - ", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer, register1, register2)

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

			s := fmt.Sprintf("D format - ", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32], Counter, Answer, register1, register2)

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
			s := fmt.Sprintf("%s %s %s %s %s %d %s R%d,R%d,R%d", jar[0:11], jar[12:16], jar[17:22], jar[23:27], jar[28:32], Counter, Answer, register1, register2, register3)
			//adding new string and linebreak
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}

		if opcode == 2038 {
			intVar, _ := strconv.Atoi(Tester)
			TwosCompDecimal, _ := strconv.ParseInt(string(intVar), 2, 64)
			s := fmt.Sprintf("%d", TwosCompDecimal)
			result += s
			result += "\n"
			Holder = ""
			jar = ""
		}
		Counter += 4
	}
	return result
}

// For project 2
func convertInstructionStringToStruct(file string) []*Instruction {
	//Open the file
	data, error := os.Open(file)
	//check for errors
	errorOpeningFile(error)
	//Scan each line of the file
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	var result []*Instruction
	var count = 96
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		var instruction = NewInstruction(line, count)
		result = append(result, instruction)
		count += 4
	}
	return result
}

// NewInstruction For project 2
func NewInstruction(data string, count int) *Instruction {

	value, error := strconv.ParseUint(data, 2, 32)
	errorOpeningFile(error)
	//Setting instruction
	instr := Instruction{
		rawInstruction: data,
		lineValue:      value,
		programCnt:     count,
	}
	return &instr
}

// For project 2
// Converting 11 bits to opcode(uint64)
func getOpcode(data string) uint64 {

	bits, error := strconv.Atoi(data)
	//check for errors
	errorOpeningFile(error)
	return uint64(binaryToDecimal(bits))

}

// For project 2
func getTypeOfInstruction(opcode uint64, jar string) string {

	var result = ""
	switch opcode {
	case 1986, 1984:
		result = "D"
	case 1104, 1112, 1360, 1624, 1690, 1691, 1692, 1872:
		result = "R"
	case 2038:
		result = "BREAK"
	default:
		if opcode >= 160 && opcode <= 191 {
			result = "B"
		}
		if opcode >= 1440 && opcode <= 1447 || opcode >= 1448 && opcode <= 1455 {
			result = "CB"
		}
		if opcode >= 1684 && opcode <= 1687 || opcode >= 1940 && opcode <= 1943 {
			result = "IM"
		}
		if opcode >= 1160 && opcode <= 1161 || opcode >= 1672 && opcode <= 1673 {
			result = "I"
		}
	}
	return result
}

//TODO for project2
//func formattedString(element Instruction) string {
//	opcode := element.rawInstruction[0:11]
//
//	s := fmt.Sprintf(
//		"%s %d %s %d\n",
//		opcode,
//		element.programCnt,
//		element.op,
//		element.rm,
//	)
//	return s
//}

// // TODO working through writeToFile
func writeToFile(path string, info string, instructionStructs []*Instruction) {
	// TODO : for project 2
	//f, error := os.OpenFile(path,
	//	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//errorOpeningFile(error)
	//defer f.Close()
	//create a loop through each struct and call write string inside the loop
	//for _, element := range instructionStructs {
	//
	//	info := formattedString(*element)
	//	_, error := f.WriteString(info)
	//	errorOpeningFile(error)
	//}

	os.WriteFile(path, []byte(info), 0666)

}

// For project 2
func getOp(opcode uint64) string {
	ValidInstructions := map[uint64]string{
		1104: "AND",
		1112: "ADD",
		//1160: "ADDI",
		//1161: "ADDI",
		1360: "ORR",
		1624: "SUB",
		//1672: "SUBI",
		//1673: "SUBI",
		1690: "LSR",
		1691: "LSL",
		1692: "ASR",
		1872: "EOR",
	}
	return ValidInstructions[opcode]
}

//TODO for project2 to build out switch cases per instruction type
//func getFileContent(structs []*Instruction) string {
//	result := ""
//	for _, s := range structs {
//		jar := s.rawInstruction
//		switch s.typeofInstruction {
//		case "R":
//			Answer = (checkForValue(int(s.opcode, ValidInstructions))
//			register1, _ := strconv.Atoi(jar[23:27])
//			register2, _ := strconv.Atoi(jar[12:16])
//			register3, _ := strconv.Atoi(jar[28:32])
//
//			register1 = binaryToDecimal(register1)
//			register2 = binaryToDecimal(register2)
//			register3 = binaryToDecimal(register3)
//
//			//printing directly to file function takes %s "string" and &d digit value
//			s := fmt.Sprintf("%s %s %s %s %s %d %s R%d,R%d,R%d",
//			jar[0:11],
//			jar[12:16],
//			jar[17:22],
//			jar[23:27],
//			jar[28:32],
//			Counter,
//			Answer,
//			register1,
//			register2,
//			register3)

//			//demo.go version needs review
//			//s := fmt.Sprintf("R format - ", jar[0:11]+" "+jar[12:16]+" "+jar[17:22]+" "+jar[23:27]+" "+jar[28:32],
//			Counter,
//			Answer,
//			register1,
//			register2,
//			register3)
//
//			//adding new string and linebreak
//			result += s
//			result += "\n"
//		}
//	}
//	return result
//}

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

	//Loop through each line in the input file
	//and make an instruction struct
	var instructionStructSlice = convertInstructionStringToStruct(*inputPath)

	for _, element := range instructionStructSlice {
		//1st eleven bits of the structs rawInstruction
		op := element.rawInstruction[0:11]

		//Set the converted opCode string -> uint64
		element.opcode = getOpcode(op)
		element.op = getOp(element.opcode)

		element.typeofInstruction = getTypeOfInstruction(element.opcode, element.rawInstruction)
	}

	os.Remove(*outputPath)
	//TODO delete file at teh start of main
	formattedInstructions := getInstructionFormat(*inputPath)
	writeToFile(*outputPath, formattedInstructions, instructionStructSlice)
}
