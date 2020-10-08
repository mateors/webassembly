## go web assemly
# Go in the Browser via Web Assembly

# Go + WASM is an Application, not a Library

# To compile from windows command line
> set GOOS=js
> set GOARCH=wasm
> go build -o ./assets/mat.wasm ./main.go

# To reduce the size of the wasm file we installed and used tinygo (https://tinygo.org/)
> tinygo build -o ./assets/mat.wasm -no-debug ./main.go

# Calling Go function from JavaScript
    `<script>
        const go = new Go();
        if (WebAssembly.instantiateStreaming) {

            WebAssembly.instantiateStreaming(fetch("mat.wasm"), go.importObject).then((result) => {
                go.run(result.instance);
                console.log(printMessage("hello"));//calling golang func directly from js
                console.log(add('val1','val2'));
                console.log(sum(5,4));
            });

        }
    </script>`