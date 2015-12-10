package main

import (
	"testing"
)

var (
	cssAssetReader = CssAssetReader{}
	cssAssetWriter = CssAssetWriter{}
	res         = []byte{}
)

type CssAssetReader struct{}

func (CssAssetReader) ReadFile(path string) ([]byte, error) {
	return ".class { width: (1 + 1) }", nil
}

type CssAssetWriter struct{}

func (CssAssetWriter) WriteFile(path string, data []byte, mode os.FileMode) error {
	res = data
	return nil
}

func TestRender(t *testing.T) {
  less.SetReader(cssAssetReader)
	less.SetWriter(cssAssetWriter)
	
	err := less.RenderFile("input", "output")
	
	if err != nil {
		t.Error("Render error")
	}
	
	if string(res) != ".class {width: 2;}" {
		t.Error("Render result invalid")
	}
}
