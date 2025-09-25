package parser_test

import (
	"testing"

	"github.com/ProImpact/templatex/internal/parser"
)

func TestParseComment(t *testing.T) {
	parse := parser.NewParser("../../test/commet.tmpl", false)
	parse.Parse()
}
