package tgl

// #cgo CFLAGS: -I ./tgl/lib
// #cgo LDFLAGS: -L ./tgl/lib/tgl/libs -l tgl
// #include "tgl.go.h"
import "C"

func login(state C.struct_tgl_state) C.struct_tgl_state { return state }
