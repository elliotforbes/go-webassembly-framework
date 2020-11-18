go 1.15 required for the example because wasm_exec.js is updated

if you see 
localhost/:1 Uncaught (in promise) LinkError: WebAssembly.instantiate(): Import #1 module="go" function="runtime.resetMemoryDataView" error: function import requires a callable
error
this is because you are using wasm_exec.js from the wrong GO version