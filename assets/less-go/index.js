var environment = require("./environment"),
    FileManager = require("./file-manager"),
    createFromEnvironment = require("../less"),
    less = createFromEnvironment(environment, [new FileManager()]);    

less.writeError = function (ctx, options) {
    options = options || {};
    if (options.silent) { return; }
    print(less.formatError(ctx, options));
};

module.exports = less;
