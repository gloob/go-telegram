/*
 * Callback exposed in Go-like style. We need to declare this callbacks
 * in another file because we are declaring them with export.
 *
 */
package tgl

//#include <tgl/tgl.h>
import "C"

import (
  "fmt"
)

//export NewMsgCB
func NewMsgCB(TLS *C.struct_tgl_state, M *C.struct_tgl_message) {
  fmt.Printf("Go.NewMsgCB(): called.\n")
}

//export MarkedReadCB
func MarkedReadCB(TLS *C.struct_tgl_state, num int, list *C.struct_tgl_message) {
  fmt.Printf("Go.MarkedReadCB(): called.\n")
}
