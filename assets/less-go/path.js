function assertPath(path) {
  if (typeof path !== 'string') {
    throw new TypeError('Path must be a string. Received ' + path);
  }
}

function isAbsolute(path) {
  assertPath(path);
  return !!path && path[0] === '/';
};

// resolves . and .. elements in a path array with directory names there
// must be no slashes or device names (c:\) in the array
// (so also no leading and trailing slashes - it does not distinguish
// relative and absolute paths)
function normalizeArray(parts, allowAboveRoot) {
  var res = [];
  for (var i = 0; i < parts.length; i++) {
    var p = parts[i];

    // ignore empty parts
    if (!p || p === '.')
      continue;

    if (p === '..') {
      if (res.length && res[res.length - 1] !== '..') {
        res.pop();
      } else if (allowAboveRoot) {
        res.push('..');
      }
    } else {
      res.push(p);
    }
  }

  return res;
}

function normalize(path) {
  assertPath(path);  

  var isAbsolutePath = isAbsolute(path),
      trailingSlash = path && path[path.length - 1] === '/';
  
  // Normalize the path
  path = normalizeArray(path.split('/'), !isAbsolutePath).join('/');

  if (!path && !isAbsolutePath) {
    path = '.';
  }
  if (path && trailingSlash) {
    path += '/';
  }

  return (isAbsolutePath ? '/' : '') + path;
};

exports.join = function () {
  var path = '';
  for (var i = 0; i < arguments.length; i++) {
    var segment = arguments[i];
    assertPath(segment);
    if (segment) {
      if (!path) {
        path += segment;
      } else {
        path += '/' + segment;
      }
    }
  }
  return normalize(path);
}