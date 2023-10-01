# Go WASM 2D Pinball
## Overview
Browser-based pinball game using 2D Canvas and WebAsssembly (WASM), compiled from Go.

Many learnings we drawn from the following projects:
- [https://github.com/go4orward/gigl](https://github.com/go4orward/gigl)
- [https://github.com/markfarnan/go-canvas](https://github.com/markfarnan/go-canvas)

## Dev Environment Setup

### Basic CLI Setup
Run the static file server:
```
cd static-server
go run main.go
```

You can now see game in the browser at [http://localhost:3000](http://localhost:3000).

Running `make` in the project root directory will recompile the Go files into to a WASM target file, and copy this into the `static-server/assets` directory. From here you just have to reload the page to see your changes.

### Further Setup with VSCode

#### Recompile Go to WASM On-Save
It's helpful to run make automatically on-save when editing Go files. This can be done in VSCode by installing the [Run On Save](https://marketplace.visualstudio.com/items?itemName=emeraldwalk.RunOnSave), with a snippet like the following in your `settings.json`:
```
    "emeraldwalk.runonsave": {
        "commands": [
            {
                "match": "\\.go$",
                "isAsync": true,
                "cmd": "make"
            }
        ]
    }
```

#### Get Intellisense to Work with `sycall/js`
The `syscall/js` package has build constraints which prevent Intellisense from working and cause VSCode to report an error with the import. One fix for this is to run VSCode from the terminal like this:
```
export GOOS=js
export GOARCH=wasm
# run the VSCode executable, which on Darwin is as follows:
/Applications/Visual\ Studio\ Code.app/Contents/Resources/app/bin/code
```