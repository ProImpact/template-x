package templatex

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type templateParse struct {
	buffer       *bufio.Reader
	templateName string
}

func NewTemplateParse(filePath string) (*templateParse, error) {
	file := filepath.Base(filePath)
	templateName, _ := strings.CutSuffix(file, filepath.Ext(file))
	templateName, _ = strings.CutPrefix(templateName, "./")
	templateName, _ = strings.CutPrefix(templateName, ".\\")

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	t := new(templateParse)
	t.templateName = templateName

	var buff bytes.Buffer

	_, err = buff.WriteString(fmt.Sprintf("{{ define \"%s\" }}\n", t.templateName))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(&buff, f)
	if err != nil {
		return nil, err
	}

	_, err = buff.WriteString("\n{{ end }}")
	if err != nil {
		return nil, err
	}

	t.buffer = bufio.NewReader(&buff)

	return t, nil
}

func (t templateParse) String() string {
	return fmt.Sprintf("[templateName: %s]\n", t.templateName)
}

func (t *templateParse) ReadTemplate() string {
	content, err := io.ReadAll(t.buffer)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s \n", content)
}
