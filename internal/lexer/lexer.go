package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Lexer struct {
	buffer       *bufio.Reader
	templateName string
	previous     *Node
	current      *Node
	LineNumber   int
	LineIndex    int
}

func (t *Lexer) Current() *Node {
	return t.current
}

func (t *Lexer) Advance() {
	token, err := t.nextToken()
	if err != nil {
		log.Fatal(err)
	}
	t.current = token
}

func NewLexer(filePath string, includeDefine bool) (*Lexer, error) {
	file := filepath.Base(filePath)
	templateName, _ := strings.CutSuffix(file, filepath.Ext(file))
	templateName, _ = strings.CutPrefix(templateName, "./")
	templateName, _ = strings.CutPrefix(templateName, ".\\")

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	t := new(Lexer)
	t.templateName = templateName

	var buff bytes.Buffer
	if includeDefine {
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
		t.LineIndex = 0
		t.LineNumber = 1
		return t, nil
	}
	_, err = io.Copy(&buff, f)
	if err != nil {
		return nil, err
	}

	t.buffer = bufio.NewReader(&buff)
	t.LineIndex = 0
	t.LineNumber = 1

	return t, nil
}

func (t Lexer) String() string {
	return fmt.Sprintf("[templateName: %s]\n", t.templateName)
}
