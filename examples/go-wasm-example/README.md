# xterm-pty Go WebAssembly example

This example demonstrates using xterm-pty with a Go program compiled to
WebAssembly (`GOOS=js GOARCH=wasm`).

Go's `wasm_exec.js` sets up a stub `globalThis.fs` object that the runtime uses
for system-call emulation.  The `index.html` patches `fs.writeSync` and
`fs.read` to route stdout/stderr writes and stdin reads through the xterm-pty
slave, so ordinary Go code that uses `fmt.Print` and `bufio.Scanner` works
without any changes.

## Build and run

1. Copy `wasm_exec.js` from your Go installation:

   ```
   cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" .
   ```

2. Build the Go program:

   ```
   GOOS=js GOARCH=wasm go build -o main.wasm main.go
   ```

3. Serve the directory (the browser requires HTTP for WebAssembly):

   ```
   npx http-server -p 3000
   ```

4. Open <http://localhost:3000> in your browser.
