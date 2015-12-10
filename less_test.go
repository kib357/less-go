package less

import (
	"os"
	"testing"
)

var (
	cssAssetReader = CssAssetReader{}
	cssAssetWriter = CssAssetWriter{}
	res            = []byte{}
)

type CssAssetReader struct{}

func (CssAssetReader) ReadFile(path string) ([]byte, error) {
	return []byte(".class { width: (1 + 1) }"), nil
}

type CssAssetWriter struct{}

func (CssAssetWriter) WriteFile(path string, data []byte, mode os.FileMode) error {
	res = data
	return nil
}

func TestRender(t *testing.T) {
	SetReader(cssAssetReader)
	SetWriter(cssAssetWriter)

	err := RenderFile("input", "output", map[string]interface{}{"compress": true})

	if err != nil {
		t.Error("Render error")
	}
	var expected = `.class{width:2}`
	if string(res) != expected {
		t.Error(`Render result invalid: "`, string(res), `" != "`, expected, `"`)
	}
}
