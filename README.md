# pinball
# Dev Environment Setup
Run the static file server:
```
cd static-server
go run main.go
```

You can rebuild the WASM file manually by running the `make` command after editing Go files. It's also helpful to do this automatically on-save. This can be done in VSCode by installing the [Run On Save](https://marketplace.visualstudio.com/items?itemName=emeraldwalk.RunOnSave), with a snippet like the following in your `settings.json`:
```
    "emeraldwalk.runonsave": {
        "commands": [
        {
            "match": "\\.go$",
            "isAsync": true,
            "cmd": "make"
        }
    ]
```