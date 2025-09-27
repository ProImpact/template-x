package parser_test

import (
	"encoding/json"
	"testing"

	"github.com/ProImpact/templatex/internal/parser"
)

func TestParseComment(t *testing.T) {
	parse := parser.NewParser("../../test/commet.tmpl", false)
	parse.Parse()
	data, err := json.MarshalIndent(&parse, " ", "   ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
