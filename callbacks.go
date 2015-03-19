/*
Callbacks exposed in Go-like style.

We need to declare this callbacks in another file because we are declaring
them with cgo export keyword to be used in the C part. At the moment we are
following exactly the callbacks functions signature defined in tgl library.

Functions signature can be checked in callbacks.c.
 */
package tgl

//#include <tgl/tgl.h>
import "C"

import (
  "fmt"
  "unsafe"
)

//export NewMsgCB
func NewMsgCB(TLS *C.struct_tgl_state, M *C.struct_tgl_message) {
  fmt.Printf("Go.NewMsgCB(): called.\n")
}

//export MarkedReadCB
func MarkedReadCB(TLS *C.struct_tgl_state, num int, list **C.struct_tgl_message) {
  fmt.Printf("Go.MarkedReadCB(): called.\n")
}

//export LogPrintfCB
func LogPrintfCB(TLS *C.struct_tgl_state) {
  fmt.Printf("Go.LogPrintfCB(): called.\n")
}

//export GetStringCB
func GetStringCB(TLS *C.struct_tgl_state, prompt *C.char, flags int) {
  fmt.Printf("Go.GetStringCB(): called.\n")
  fmt.Printf("TGL: %+v.\n", TLS)
  fmt.Printf("prompt: %+v.\n", prompt)
  fmt.Printf("flags: %+v.\n", flags)
}

//export LoggedInCB
func LoggedInCB(TLS *C.struct_tgl_state) {
  fmt.Printf("Go.LoggedInCB(): called.\n")
}

//export StartedCB
func StartedCB(TLS *C.struct_tgl_state) {
  fmt.Printf("Go.StartedCB(): called.\n")
}

//export TypeNotificationCB
func TypeNotificationCB(TLS *C.struct_tgl_state, U *C.struct_tgl_user, status C.enum_tgl_typing_status) {
  fmt.Printf("Go.TypeNotificationCB(): called.\n")
}

//export TypeInChatNotificationCB
func TypeInChatNotificationCB(TLS *C.struct_tgl_state, U *C.struct_tgl_user, C *C.struct_tgl_chat, status C.enum_tgl_typing_status) {
  fmt.Printf("Go.TypeInChatNotificationCB(): called.\n")
}

//export TypeInSecretChatNotificationCB
func TypeInSecretChatNotificationCB(TLS *C.struct_tgl_state, E *C.struct_tgl_secret_chat) {
  fmt.Printf("Go.TypeInSecretChatNotificationCB(): called.\n")
}

//export StatusNotificationCB
func StatusNotificationCB(TLS *C.struct_tgl_state, U *C.struct_tgl_user) {
  fmt.Printf("Go.StatusNotificationCB(): called.\n")
}

//export UserRegisteredCB
func UserRegisteredCB(TLS *C.struct_tgl_state, U *C.struct_tgl_user) {
  fmt.Printf("Go.UserRegisteredCB(): called.\n")
}

//export UserActivatedCB
func UserActivatedCB(TLS *C.struct_tgl_state, U *C.struct_tgl_user) {
  fmt.Printf("Go.UserActivatedCB(): called.\n")
}

//export NewAuthorizationCB
func NewAuthorizationCB(TLS *C.struct_tgl_state, device *C.char, location *C.char) {
  fmt.Printf("Go.NewAuthorizationCB(): called.\n")
}

//export ChatUpdateCB
func ChatUpdateCB(TLS *C.struct_tgl_state, C *C.struct_tgl_chat, flags uint) {
  fmt.Printf("Go.ChatUpdateCB(): called.\n")
}

//export UserUpdateCB
func UserUpdateCB(TLS *C.struct_tgl_state, C *C.struct_tgl_user, flags uint) {
  fmt.Printf("Go.UserUpdateCB(): called.\n")
}

//export SecretChatUpdateCB
func SecretChatUpdateCB(TLS *C.struct_tgl_state, C *C.struct_tgl_secret_chat, flags uint) {
  fmt.Printf("Go.SecretChatUpdateCB(): called.\n")
}

//export MsgReceiveCB
func MsgReceiveCB(TLS *C.struct_tgl_state, M *C.struct_tgl_message) {
  fmt.Printf("Go.MsgReceiveCB(): called.\n")
}

//export OurIdCB
func OurIdCB(TLS *C.struct_tgl_state, id int) {
  fmt.Printf("Go.OurIdCB(): called.\n")
}

//export NotificationCB
func NotificationCB(TLS *C.struct_tgl_state, t *C.char, message *C.char) {
  fmt.Printf("Go.NotificationCB(): called.\n")
}

//export UserStatusUpdate
func UserStatusUpdate(TLS *C.struct_tgl_state, U *C.struct_tgl_user) {
  fmt.Printf("Go.UserStatusUpdate(): called.\n")
}

//export CreatePrintNameCB
func CreatePrintNameCB(TLS *C.struct_tgl_state, id C.tgl_peer_id_t, a1 *C.char, a2 *C.char, a3 *C.char, a4 *C.char) *C.char {
  fmt.Printf("Go.CreatePrintNameCB(): called.\n")
  fmt.Printf("id: %+v.\n", id)
  fmt.Printf("a1: %+v.\n", a1)

  app := "Federator"
  aptr := C.CString(app)
  defer C.free(unsafe.Pointer(aptr))

  return aptr
}
