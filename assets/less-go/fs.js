exports.statSync = function (path) {
	var res = readFile(path);
	if (typeof res !== 'string') {
		throw new Error('File not found');
	}
}

exports.readFileSync = function (path, encoding) {	
	print('read file sync');
	print(path);
	return readFile(path);
}