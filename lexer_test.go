package templatex_test

import (
	"log"
	"testing"

	"github.com/ProImpact/templatex"
)

func TestGetTokens(t *testing.T) {
	tmpl, err := templatex.NewTemplateParse("./test/all.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	for {
		token, err := tmpl.Next()
		if err != nil {
			log.Fatal(err)
		}
		if token.NodeType == templatex.EOF {
			t.Logf("%+v \n", token)
			break
		}
		t.Logf("%+v \n", token)
	}
}
