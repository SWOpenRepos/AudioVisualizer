@echo off
docker run ^
  --rm ^
  -v "%cd%/control_ert_rtw:/src" ^
  emscripten/emsdk:latest ^
  bash -c "sed -i 's/ EMSCRIPTEN_KEEPALIVE//g' control.h && emcc control.c control_data.c rt_nonfinite.c rtGetInf.c rtGetNaN.c -o control.js -s WASM=1 -s EXPORTED_FUNCTIONS=\"['_control_initialize','_control_step','getValue','setValue']\" -s EXPORTED_RUNTIME_METHODS=['cwrap','ccall'] -s ALLOW_MEMORY_GROWTH=1 -s INITIAL_MEMORY=512KB -s STANDALONE_WASM=1 --no-entry -O3"