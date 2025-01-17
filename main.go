package main

import (
	FileIO "assembler/lib/fileio"
	Logs "assembler/lib/logs"
	Mappings "assembler/lib/mappings"
	Parser "assembler/lib/parser"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	instructionMap := make(map[string]string)
	Mappings.FillInstructionMap(&instructionMap)
	memoryMap := make(map[string]string)
	memMapIndex := 16
	fileName := os.Args[1]

	firstRun(fmt.Sprintf("%s.asm", fileName), fmt.Sprintf("%s.hack", fileName), &instructionMap, &memoryMap)
	//Mappings.PrintMapContents(&memoryMap)
	secondRun(fmt.Sprintf("%s.hack", fileName), fmt.Sprintf("%s.hackfinal", fileName), &memoryMap, &memMapIndex)

	FileIO.DeleteFile(fmt.Sprintf("%s.hack", fileName))
	FileIO.RenameFile(fmt.Sprintf("%s.hackfinal", fileName), fmt.Sprintf("%s.hack", fileName))

	Logs.LogExecutionTime(start, "TaskName")
	Logs.LogMemoryUsage()
}

// region firstRun
func firstRun(sourceName string, tempName string, instructionMap *map[string]string, memoryMap *map[string]string) {
	source, _ := os.Open(sourceName)
	defer source.Close()
	temp, _ := os.Create(tempName)
	defer temp.Close()

	reader := bufio.NewReader(source)
	writer := bufio.NewWriter(temp)
	defer writer.Flush()

	stepIndex := 1
	labelsCount := 0

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineString := string(line)
		//*remove whitespaces & comments
		Parser.RemoveComment(&lineString)
		if Parser.RemoveWhitespaceAndEmptyLines(&lineString) {
			if line[0] == 0x40 {
				//*A-instruction
				var aInstruction string = parseAinstruction(&lineString, instructionMap)
				writer.WriteString(aInstruction + "\n")
			} else if line[0] == 0x28 {
				////*Labels (xxx):  parse label -> 1st run
				mapLabelHeader(&lineString, memoryMap, stepIndex-labelsCount)
				labelsCount++
			} else {
				//*C-instruction
				var cInstruction string = parseCInstruction(&lineString, instructionMap)
				//fmt.Println(cInstruction)
				writer.WriteString(cInstruction + "\n")
			}
			//step indicates line -> move +1 if not label header
			stepIndex++
		}
	}
}

// region secondRun
func secondRun(sourceName string, tempName string, memoryMap *map[string]string, memMapIndex *int) {
	source, _ := os.Open(sourceName)
	defer source.Close()
	temp, _ := os.Create(tempName)
	defer temp.Close()

	reader := bufio.NewReader(source)
	writer := bufio.NewWriter(temp)
	defer writer.Flush()

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineString := string(line)
		//*remaining A-instruction (@a) (labels/variables)
		if lineString[0] == 0x40 {
			if val, exists := (*memoryMap)[lineString[1:]]; exists {
				//already marked in memory
				byte16 := Parser.To15BitBinary(val)
				writer.WriteString(fmt.Sprintf("0%s\n", byte16))
			} else {
				//variable declaration -> mark now in empty memoryIndex
				Mappings.AddToMemoryMap(memoryMap, memMapIndex, lineString[1:])
				byte16 := Parser.To15BitBinary(strconv.Itoa(*memMapIndex))
				//fmt.Println(byte16)
				writer.WriteString(fmt.Sprintf("0%s\n", byte16))
			}
		} else {
			//not a instruction (already parsed) -> copy
			writer.WriteString(fmt.Sprintf("%s\n", lineString))
		}
	}
}

// * A-instruction
func parseAinstruction(line *string, instrMap *map[string]string) string {
	*line = (*line)[1:]
	var binaryStr string

	if val, exists := (*instrMap)[*line]; exists {
		binaryStr = fmt.Sprintf("0%s", Parser.To15BitBinary(val))
	} else if (*line)[0] >= '0' && (*line)[0] <= '9' { //*if non-label A-instruction
		binaryStr = fmt.Sprintf("0%s", Parser.To15BitBinary(*line))
	} else { //*label A-instruction -> translate in 2nd run
		binaryStr = "@" + *line
	}
	return binaryStr
}

// *A Labels -1st run.
func mapLabelHeader(line *string, memMap *map[string]string, stepIndex int) {
	label := Parser.ExtractLabel(*line)
	Mappings.AddToMemoryMap(memMap, &stepIndex, label)
}

func parseCInstruction(line *string, instrMap *map[string]string) string {
	comp, dest, jump := Parser.ParseCInstruction(line)
	// fmt.Printf("comp: %s, dest: %s, jump: %s\n", comp, dest, jump)
	compBits := (*instrMap)[("c" + comp)]
	destBits := (*instrMap)[("d" + dest)]
	jumpBits := (*instrMap)[("j" + jump)]

	cInstructionBits := fmt.Sprintf("111%s%s%s", compBits, destBits, jumpBits)
	return cInstructionBits
}
