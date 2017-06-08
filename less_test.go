package less

import (
	"testing"
	"bytes"
)

func TestRender(t *testing.T) {
	w := new(bytes.Buffer)
	lessCompiler := NewLessCompiler()
	//TODO check compress:false it does not work correctly
	err := lessCompiler.Compile("assets/test/signature.less", w, map[string]interface{}{"compress": true})

	if err != nil {
		t.Error("Render error", err)
	}
	res := w.String()
	var expected = `.signer-box{border:1px solid #855b85}.signer-box .sign-actions{font-size:30px;width:100%;text-align:center;display:block;line-height:normal;vertical-align:middle;background:rgba(133,91,133,0.4);padding:15px 0}.signer-box .sign-actions>button.btn{margin:5px;white-space:pre-wrap;text-align:center;max-width:80%;display:inline-block;vertical-align:middle;white-space:nowrap;overflow:hidden;text-overflow:ellipsis}.signer-box .innerbtn-main:hover{-webkit-box-shadow:inset 0 0 0 1px #855b85;-moz-box-shadow:inset 0 0 0 1px #855b85;box-shadow:inset 0 0 0 1px #855b85}`
	if res != expected {
		t.Error(`Render result invalid: "`, res, `" != "`, expected, `"`)
	}
}
