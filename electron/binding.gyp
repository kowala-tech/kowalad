{
    "targets": [
        {
            "target_name": "addon",
            "sources": [ "backend.cc"]
        }
    ]
}

#1 node-gyp configure to generate the appropriate project build files for the current platform
#2 node-gyp build command to generate the compiled addon.node (This will be put into the build/Release/ directory.) 
#3 Once built, the binary Addon can be used from within Node.js by pointing require() to the built addon.node module:
#
#// hello.js
#const addon = require('./build/Release/addon');
#
#console.log(addon.hello());
#// Prints: 'world'