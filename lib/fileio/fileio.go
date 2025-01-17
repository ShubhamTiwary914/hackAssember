package assembler

import (
	"bufio"
	"os"
)

func MakeFilePairs_sourceTemp(sourceName string, tempName string) (*bufio.Reader, *bufio.Writer) {
	source, _ := os.Open(sourceName)
	defer source.Close()
	temp, _ := os.Create(tempName)
	defer temp.Close()

	reader := bufio.NewReader(source)
	writer := bufio.NewWriter(temp)
	defer writer.Flush()

	return reader, writer
}

func RenameFile(oldName, newName string) {
	os.Rename(oldName, newName)
}

func DeleteFile(fileName string) {
	os.Remove(fileName)
}
