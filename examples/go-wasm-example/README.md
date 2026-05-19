# xterm-pty Go/Wasm example (`wasm_exec.js`)

This example shows how to connect `xterm-pty` to a Go WebAssembly program running through `wasm_exec.js`.

Unlike the Emscripten integration, Go's browser runtime does not expose a PTY-compatible stdin/stdout layer, so this example uses a small JavaScript bridge:

- `goPty.write(text)` writes to the PTY slave
- `goPty.onData(callback)` forwards cooked terminal input to Go
- `goPty.onResize(callback)` forwards `SIGWINCH`
- `goPty.getSize()` returns the current terminal size

## Build

From this directory:

```sh
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" .
GOOS=js GOARCH=wasm go build -o main.wasm ./main.go
```

## Run

From the repository root:

```sh
npx http-server -p 3000
```

Then open <http://localhost:3000/examples/go-wasm-example/>.

Resize support comes from `@xterm/addon-fit` + `ResizeObserver`: each resize updates `xterm.js`, `xterm-pty` emits `SIGWINCH`, and the Go code prints the new rows/columns.
