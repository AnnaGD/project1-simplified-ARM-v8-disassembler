package main

import (
	"bufio"
	"fmt"
	"os"
)

func RorDInst(r []string, d []string, str string) string {
	for _, v := range r {
		if v == str {
			return "The instruction format for this is R"
		}
	}
	for _, v := range d {
		if v == str {
			return "The instruction format for this is D"
		}
	}
	return ""
}

func contains(s []string, str string) string {
	for _, v := range s {

		if len(v) == 6 {
			if v == str {
				return "the instruction format for this is B"
			}
		}

		if len(v) == 8 {
			if v == str {
				return "the instruction format for this is CB"
			}
		}

		if len(v) == 9 {
			if v == str {
				return "the instruction format for this is IM"
			}
		}

		if len(v) == 11 {
			if v == str {
				return "the instruction format for this is R"
			}
		}

	}
	return ""
}

func errorOpeningFile(e error) {
	if e != nil {
		panic(e)
	}
}

func getInstructionFormat(content string) {
	data, error := os.Open(content)
	errorOpeningFile(error)
	var Tester = ""
	var Holder = ""

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		Tester = fileScanner.Text()
		chars := []rune(Tester)

		for i := 0; i <= 10; i++ {
			Holder += string(chars[i])
			if len(Holder) == 6 {
				BFormat := []string{"000101"}
				fmt.Println(contains(BFormat, Holder))
			}
			if len(Holder) == 8 {
				CBFormat := []string{"10110100", "10110101"}
				fmt.Println(contains(CBFormat, Holder))
			}
			if len(Holder) == 9 {
				IMFormat := []string{"110100101", "111100101"}
				fmt.Println(contains(IMFormat, Holder))
			}
			if len(Holder) == 10 {
				IFormat := []string{"1001000100", "1101000100"}
				fmt.Println(contains(IFormat, Holder))
			}
		}
		if len(Holder) == 11 {
			Rformat := []string{"10001010000", "10001011000", "10101010000", "11001011000", "11010011010", "11010011011", "11010011100", "11101010000"}
			Dformat := []string{"11111000000", "11111000010"}
			fmt.Println(RorDInst(Rformat, Dformat, Holder))
			Holder = ""
		}
	}
}

func main() {
	var Path = "src/temp/data.txt"
	File, err := os.Create("team1_out_dis.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println(err)
		File.Close()
		return
	}

	err = File.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	getInstructionFormat(Path)
}
