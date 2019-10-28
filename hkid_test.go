package hkid

import "testing"

type correctData struct {
	input  string
	part1  string
	part2  string
	part3  string
	output string
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
