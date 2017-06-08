package less

import (
	"encoding/json"
	"errors"
	"gopkg.in/olebedev/go-duktape.v2"
	"io/ioutil"
	"bytes"
)

var (
	ctx *duktape.Context
)

type LessCompiler struct{
	writer *bytes.Buffer
}

func (l *LessCompiler) readFile(c *duktape.Context) int {
	var path = c.SafeToString(-1)
	if path == "" {
		return 0
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		b, err = Asset(path)
		if err != nil {
			return 0
		}
	}
	c.PushString(string(b))
	return 1
}

func (l *LessCompiler) readFileFromAssets(c *duktape.Context) int {
	var path = c.SafeToString(-1)
	if path == "" {
		return 0
	}
	b, err := Asset(path)
	if err != nil {
		return 0
	}
	c.PushString(string(b))
	return 1
}

func (l *LessCompiler) writeFile(c *duktape.Context) int {
	var data = []byte(c.SafeToString(-1))
	var path = c.SafeToString(-2)
	if path == "" {
		return 0
	}
	l.writer.Write(data)
	return 1
}

func NewLessCompiler() *LessCompiler{
	l := &LessCompiler{}
	ctx = duktape.New()
	ctx.PushGlobalGoFunction("readFile", l.readFile)
	ctx.PushGlobalGoFunction("readFileFromAssets", l.readFileFromAssets)
	ctx.PushGlobalGoFunction("writeFile", l.writeFile)
	return l
}

func (l *LessCompiler) Compile(input string, wb *bytes.Buffer, mods ...map[string]interface{}) error {
	l.writer = wb

	ctx.EvalString(`
		Duktape.modSearch = function (id, require, exports, module) {
			id = id.replace(/\.js$/, "");
			var res = readFileFromAssets(id + ".js");
			if (typeof res === 'string') {
				return res;
			}

			var res = readFileFromAssets(id + "/index.js");
			if (typeof res === 'string') {
				return 'module.exports = require("' + id + '/index.js")';
			}
		    throw new Error('module not found: ' + id);
		};
	`)

	if input == "" {
		return errors.New("No input path provided")
	}
	var options = map[string]interface{}{}
	if len(mods) > 0 {
		options = mods[0]
	}
	options["filename"] = input
	encodedOptions, err := json.Marshal(options)
	if err != nil {
		return err
	}
	ctx.EvalString(`
		try {
			var fs = require('./assets/less-go/fs');
			var less = require('./assets/less-go');

			var data = fs.readFileSync("` + input + `");
			less.render(data, ` + string(encodedOptions) + `, function (e, output) {
				if (e == null) {
					print("Rendered");
					writeFile(null, output.css);
				} else {
					print('Render error', e.stack);
					e.stack
				}
			});
		} catch (e) {
			print("ERROR!", e.stack);
			e.stack
		}
	`)
	result := ctx.GetString(-1)
	ctx.Pop()
	if result != "" {
		return errors.New(result)
	}
	return nil
}

func init() {


	// result := ctx.GetString(-1)
	// ctx.Pop()
	// fmt.Println("result is:", result)
}
