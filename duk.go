package main

import (
	"fmt"
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
			var env = require("environment");
			var fs = require("fs");			
			var less = require("less");
			var fm = require('LessFileManager.js');
			var ffm = new fm();
			var compiler = less(env, [ffm]); 					 						
			print('ffm');			
			
			compiler.logger.addListener({
				debug: function(msg) {
					print(msg);
    			},	
				info: function(msg) {
					print(msg);
				},
				warn: function(msg) {
					print(msg);
				},
				error: function(msg) {
					print(msg);
				}
			});						
			
			var data = fs.readFileSync('./css/app.less');		
			print(data.slice(0,50) + '...');
			compiler.render(data, {filename: "./css/app.less"}, function (e, output) {
				print("--------------------");
				//print(output.css);
				writeFile("./css/app.css", output.css);		 
				print(e);		 
				print("--------------------");
			});
			
			print('________end__________');
									
		} catch (e) {
			print("ERROR!", e)
		}
	`)
	result := ctx.GetString(-1)
	ctx.Pop()
	fmt.Println("result is:", result)
}
