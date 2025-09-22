package templatex_test

import (
	"log"
	"testing"

	"github.com/ProImpact/templatex"
)

func TestGetTokens(t *testing.T) {
	tmpl, err := templatex.NewTemplateParse("./test/conditional.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	tokens, err := tmpl.Parse()
	if err != nil {
		log.Fatal(err)
	}
	for _, token := range tokens {
		t.Logf("%+v \n", token)
	}
}
