package main

import (
	//"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
	"io/ioutil"
)

func readFile(c *duktape.Context) int {
	var path = c.SafeToString(-1)
	if path == "" {
		return 0
	}
	bytes, err := ioutil.ReadFile(path)
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
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return 0
	}	
	return 1
}

func main() {
	ctx := duktape.New()
	ctx.PushGlobalGoFunction("readFile", readFile)
	ctx.PushGlobalGoFunction("writeFile", writeFile)

	ctx.EvalString(`
		Duktape.modSearch = function (id, require, exports, module) {
			id = id.replace(/\.js$/, "");
			var res = readFile(id + ".js");
			if (typeof res === 'string') {			
				print('Loading module from ' + id + ".js");
				return res;
			}

			var res = readFile(id + "/index.js");
			if (typeof res === 'string') {
				print('Loading module from ' + id + '/index.js');
				
				return 'module.exports = require("' + id + '/index.js")';			
			}
		    throw new Error('module not found: ' + id);
		};
	`)

	ctx.PevalString(`
		try {
			var fs = require('./assets/less-go/fs');
			var less = require('./assets/less-go');
			
			var data = fs.readFileSync('./css/app.less');		
			print('Rendering less: ', data.slice(0,50) + '...');
			less.render(data, {filename: "./css/app.less"}, function (e, output) {
				if (e == null) {
					print("Compiled");
					writeFile("./css/app.css", output.css);
				} else {		 
					print('Compile error', e.stack);
				}		 
			});										
		} catch (e) {
			print("ERROR!", e.stack);
		}
	`)
	// result := ctx.GetString(-1)
	// ctx.Pop()
	// fmt.Println("result is:", result)
}
