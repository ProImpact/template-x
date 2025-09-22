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
	buffer       *bufio.ReadWriter
	templateName string
}

func NewTemplateParse(filePath string) (*templateParse, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	t := new(templateParse)
	templateName := fmt.Sprintf("{{ define \"%s\" }}\n", t.templateName)
	buff := bytes.NewBuffer([]byte(templateName))
	_, err = io.Copy(buff, f)
	if err != nil {
		return nil, err
	}
	_, err = buff.Write([]byte("\n{{ end }}"))
	if err != nil {
		return nil, err
	}
	io.ReadAll(buff)
	t.buffer = bufio.NewReadWriter(bufio.NewReader(buff), bufio.NewWriter(buff))
	file := filepath.Base(filePath)
	file, _ = strings.CutSuffix(filePath, filepath.Ext(file))
	file, _ = strings.CutPrefix(file, "./")
	file, _ = strings.CutPrefix(file, ".\\")
	t.templateName = file
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
