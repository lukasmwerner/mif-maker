package internal

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strings"
)

type MifGenerator struct {
	converter *Uint4Converter
}

func NewMifGenerator() *MifGenerator {
	return &MifGenerator{
		converter: NewUint4Converter(),
	}
}

func (mg *MifGenerator) GenerateMif(img image.Image) (string, error) {
	bounds := img.Bounds()
	depth := bounds.Dx() * bounds.Dy()
	width := 4 * 4

	dataLines, err := mg.converter.ExtractU4Data(img)
	if err != nil {
		return "", fmt.Errorf("failed to extract U4 data: %w", err)
	}

	header := mg.formatHeader(depth, width, "HEX", "HEX")
	content := mg.formatContent(dataLines)

	var mifLines []string
	mifLines = append(mifLines, header...)
	mifLines = append(mifLines, content...)

	return strings.Join(mifLines, "\n"), nil
}

func (mg *MifGenerator) WriteMif(img image.Image, filename string) error {
	mifContent, err := mg.GenerateMif(img)
	if err != nil {
		return fmt.Errorf("failed to generate MIF content: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create MIF file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(mifContent)
	if err != nil {
		return fmt.Errorf("failed to write MIF content: %w", err)
	}
	writer.Flush()

	fmt.Printf("MIF file saved as: %s\n", filename)
	return nil
}

func (mg *MifGenerator) formatHeader(depth, width int, addrRadix, dataRadix string) []string {
	return []string{
		fmt.Sprintf("DEPTH = %d;", depth),
		fmt.Sprintf("WIDTH = %d;", width),
		"",
		fmt.Sprintf("ADDRESS_RADIX = %s;", addrRadix),
		fmt.Sprintf("DATA_RADIX = %s;", dataRadix),
		"",
	}
}

func (mg *MifGenerator) formatContent(dataLines []string) []string {
	content := []string{"CONTENT", "BEGIN"}
	content = append(content, dataLines...)
	content = append(content, "END;")
	return content
}