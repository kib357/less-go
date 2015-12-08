package less

import (
	"encoding/json"
	"errors"
	"gopkg.in/olebedev/go-duktape.v2"
	"io/ioutil"
	"os"
)

var (
	ctx *duktape.Context
	r   Reader
	w   Writer
)

type Reader interface {
	ReadFile(string) ([]byte, error)
}

type Writer interface {
	WriteFile(string, []byte, os.FileMode) error
}

type reader struct{}

func (reader) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

type writer struct{}

func (writer) WriteFile(path string, data []byte, mode os.FileMode) error {
	return ioutil.WriteFile(path, data, mode)
}

func readFile(c *duktape.Context) int {
	var path = c.SafeToString(-1)
	if path == "" {
		return 0
	}
	bytes, err := r.ReadFile(path)
	if err != nil {
		bytes, err = Asset(path)
		if err != nil {
			return 0
		}
	}
	c.PushString(string(bytes))
	return 1
}

func readFileFromAssets(c *duktape.Context) int {
	var path = c.SafeToString(-1)
	if path == "" {
		return 0
	}
	bytes, err := Asset(path)
	if err != nil {
		return 0
	}
	c.PushString(string(bytes))
	return 1
}

func writeFile(c *duktape.Context) int {
	var data = []byte(c.SafeToString(-1))
	var path = c.SafeToString(-2)
	if path == "" {
		return 0
	}
	err := w.WriteFile(path, data, 0644)
	if err != nil {
		return 0
	}
	return 1
}

func SetReader(customReader Reader) {
	r = customReader
}

func SetWriter(customWriter Writer) {
	w = customWriter
}

func RenderFile(input, output string, mods ...map[string]interface{}) error {
	if input == "" {
		return errors.New("No input path provided")
	}
	if output == "" {
		output = input + ".css"
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
			print('Rendering less: ', data.slice(0,50) + '...');
			less.render(data, ` + string(encodedOptions) + `, function (e, output) {
				if (e == null) {
					print("Rendered");
					writeFile("` + output + `", output.css);
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
	r = reader{}
	w = writer{}
	ctx = duktape.New()
	ctx.PushGlobalGoFunction("readFile", readFile)
	ctx.PushGlobalGoFunction("readFileFromAssets", readFileFromAssets)
	ctx.PushGlobalGoFunction("writeFile", writeFile)

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

	// result := ctx.GetString(-1)
	// ctx.Pop()
	// fmt.Println("result is:", result)
}
