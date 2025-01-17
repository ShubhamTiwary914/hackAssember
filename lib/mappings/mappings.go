package assembler

import (
	"fmt"
	"strconv"
)

func FillInstructionMap(instrMap *map[string]string) {
	//*A-instruction
	//R0 to R15
	for i := 0; i <= 15; i++ {
		(*instrMap)[fmt.Sprintf("R%d", i)] = strconv.Itoa(i)
	}
	(*instrMap)["SCREEN"] = "16384"
	(*instrMap)["KBD"] = "24576"
	(*instrMap)["SP"] = "0"
	(*instrMap)["LCL"] = "1"
	(*instrMap)["ARG"] = "2"
	(*instrMap)["THIS"] = "3"
	(*instrMap)["THAT"] = "4"

	//*for C-instruction
	//>NOTE:
	//>For comp:  format: acccccc  (prefix (a) = 0/1 added here)
	//>For all(comp/dest/jump) -> key has prefix (c/d/j),  ex:  null -> "cnull", to handle key collisions.
	//COMP
	(*instrMap)["cnull"] = "0101010" // null
	(*instrMap)["c0"] = "0110000"    // 0
	(*instrMap)["c1"] = "1111111"    // 1
	(*instrMap)["c-1"] = "1110100"   // -1
	(*instrMap)["cD"] = "0001100"    // D
	(*instrMap)["cA"] = "0110000"    // A
	(*instrMap)["cM"] = "1110000"    // M
	(*instrMap)["c!D"] = "0001101"   // !D
	(*instrMap)["c!A"] = "0110001"   // !A
	(*instrMap)["c!M"] = "1110001"   // !M
	(*instrMap)["c-D"] = "0001111"   // -D
	(*instrMap)["c-A"] = "0110011"   // -A
	(*instrMap)["c-M"] = "1110011"   // -M
	(*instrMap)["cD+1"] = "0011111"  // D+1
	(*instrMap)["cA+1"] = "0110111"  // A+1
	(*instrMap)["cM+1"] = "1110111"  // M+1
	(*instrMap)["cD-1"] = "0001110"  // D-1
	(*instrMap)["cA-1"] = "0110010"  // A-1
	(*instrMap)["cM-1"] = "1110010"  // M-1
	(*instrMap)["cD+A"] = "0000010"  // D+A
	(*instrMap)["cD+M"] = "1000010"  // D+M
	(*instrMap)["cD-A"] = "0010011"  // D-A
	(*instrMap)["cD-M"] = "1010011"  // D-M
	(*instrMap)["cA-D"] = "0000111"  // A-D
	(*instrMap)["cM-D"] = "1000111"  // M-D
	(*instrMap)["cD&A"] = "0000000"  // D&A
	(*instrMap)["cD&M"] = "1000000"  // D&M
	(*instrMap)["cD|A"] = "0010101"  // D|A
	(*instrMap)["cD|M"] = "1010101"  // D|M

	//DEST
	(*instrMap)["dnull"] = "000" // null
	(*instrMap)["dM"] = "001"    // M
	(*instrMap)["dD"] = "010"    // D
	(*instrMap)["dA"] = "100"    // A
	(*instrMap)["dDM"] = "011"   // M+D
	(*instrMap)["dAM"] = "101"   // A+M
	(*instrMap)["dAD"] = "110"   // A+D
	(*instrMap)["dADM"] = "111"  // A+D+M

	//JMP
	(*instrMap)["jnull"] = "000" // null
	(*instrMap)["jJGT"] = "001"  // JGT
	(*instrMap)["jJEQ"] = "010"  // JEQ
	(*instrMap)["jJGE"] = "011"  // JGE
	(*instrMap)["jJLT"] = "100"  // JLT
	(*instrMap)["jJNE"] = "101"  // JNE
	(*instrMap)["jJLE"] = "110"  // JLE
	(*instrMap)["jJMP"] = "111"  // JMP
}

func AddToMemoryMap(memMap *map[string]string, index *int, varname string) {
	(*memMap)[varname] = strconv.Itoa((*index))
	(*index)++
}

func PrintMapContents(mapObj *map[string]string) {
	for key, val := range *mapObj {
		fmt.Println(fmt.Sprintf("%s: %s", key, val))
	}
}
