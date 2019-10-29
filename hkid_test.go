package hkid

import "testing"

type correctData struct {
	input  string
	part1  string
	part2  string
	part3  string
	output string
}

type wrongFormatData struct {
	input string
	part1 string
	part2 string
	part3 string
}

func newWrongFormatData(input string, part1 string, part2 string, part3 string) wrongFormatData {
	return wrongFormatData{input: input, part1: part1, part2: part2, part3: part3}
}

func newCorrectData(input string, part1 string, part2 string, part3 string, output string) correctData {
	return correctData{input: input, part1: part1, part2: part2, part3: part3, output: output}
}

func getCorrectData() [7]correctData {
	return [7]correctData{
		newCorrectData("B111111(1)", "B", "111111", "1", "B111111(1)"),
		newCorrectData("CA182361(1)", "CA", "182361", "1", "CA182361(1)"),
		newCorrectData("ZA182361(3)", "ZA", "182361", "3", "ZA182361(3)"),
		newCorrectData("B111112(A)", "B", "111112", "A", "B111112(A)"),
		newCorrectData("B111117(0)", "B", "111117", "0", "B111117(0)"),
		newCorrectData(" B111117(0)", " B", "111117", "0", "B111117(0)"),
		newCorrectData("z109852(8)", "z", "109852", "8", "Z109852(8)"),
	}
}

func getWrongFormatData() [6]wrongFormatData {
	return [6]wrongFormatData{
		newWrongFormatData("B111111(3)", "B", "111111", "3"),
		newWrongFormatData("CAC182361(1)", "CaC", "182361", "1"),
		newWrongFormatData("111112(A)", "", "111112", "A"),
		newWrongFormatData("1B111117(0)", "1B", "111117", "0"),
		newWrongFormatData("122(0)", "1", "22", "0"),
		newWrongFormatData("111111", "1", "111", "1"),
	}
}

func TestCheckStringFalse(t *testing.T) {
	data := getWrongFormatData()
	for _, v := range data {
		t.Run(v.input, func(t *testing.T) {
			hkidResult := CheckString(v.input)
			if hkidResult.Valid == true {
				t.Errorf("%s should false", v.input)
			}
		})
	}
}

func TestCheckPartFalse(t *testing.T) {
	data := getWrongFormatData()
	for _, v := range data {
		t.Run(v.input, func(t *testing.T) {
			hkidResult := CheckPart(v.part1, v.part2, v.part3)
			if hkidResult.Valid == true {
				t.Errorf("%s should false", v.input)
			}
		})
	}
}

func TestCheckStringTrue(t *testing.T) {
	data := getCorrectData()
	for _, v := range data {
		t.Run(v.input, func(t *testing.T) {
			hkidResult := CheckString(v.input)
			if hkidResult.Valid == false {
				t.Errorf("%s should true", v.output)
			}
			if hkidResult.Format() != v.output {
				t.Errorf("%s should Equals", v.output)
			}
			if hkidResult.GetPart1() != clearString(v.part1) {
				t.Errorf("%s should Equals", v.part1)
			}
			if hkidResult.GetPart2() != clearString(v.part2) {
				t.Errorf("%s should Equals", v.part2)
			}
			if hkidResult.GetPart3() != clearString(v.part3) {
				t.Errorf("%s should Equals", v.part3)
			}
		})
	}
}

func TestCheckPartTrue(t *testing.T) {
	data := getCorrectData()
	for _, v := range data {
		t.Run(v.input, func(t *testing.T) {
			hkidResult := CheckPart(v.part1, v.part2, v.part3)
			if hkidResult.Valid == false {
				t.Errorf("%s should true", v.output)
			}
			if hkidResult.Format() != v.output {
				t.Errorf("%s should Equals", v.output)
			}
		})
	}
}
