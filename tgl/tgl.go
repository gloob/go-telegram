/*
Telegram Go library.
*/
package tgl

/*
#include "tgl.go.h"

#cgo CFLAGS: -I ./lib
#cgo LDFLAGS: -L ./lib/tgl/libs -l tgl
#cgo pkg-config: libevent
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// State wraps a tgl_state struct.
type State struct {
	// We need to have in mind that we cannot modify the internal structure
	// because in the library C-side the struct is packed.
	// This is golang-cgo issue #7560 https://github.com/golang/go/issues/7560
	// That's the reason that some fields are not accessible from the Go side.
	inner *C.struct_tgl_state
}

// NewState creates a new State struct.
func NewState() *State {
	return &State{C.tgl_state_alloc()}
}

func (s *State) Dial() *error {
	config := NewConfig(s)
	defer config.Destroy()

	config.setBinlog("binlog", 10)

	config.setRsaKey("./rsa.pub")

	hash := "34be6d99874fb9607fe932dbb86fe4a3"
	hptr := C.CString(hash)
	defer C.free(unsafe.Pointer(hptr))
	app := "Federator"
	aptr := C.CString(app)
	defer C.free(unsafe.Pointer(aptr))

	C.tgl_register_app_id(s.inner, C.int(10604), hptr)
	C.tgl_set_app_version(s.inner, aptr)

	s.EnableCallbacks()

	fmt.Printf("inner: %+v\n", s.inner)

	// Bug on C.tlg: If RSA file location doesn't exists it segfault at
	// tgl/mtproto-client.c:1263 (tglmp_on_start)
	C.tgl_init(s.inner)

	// Initial handshake.
	C.read_auth_file(s.inner)

	// reset_authorization
	P := C.tgl_peer_get(s.inner, C.macro_TGL_MK_USER(s.inner.our_id))
	fmt.Printf("P: %+v\n", P)
	//C.set_default_username(P.user.phone);
	C.bl_do_reset_authorization(s.inner)

	C.tgl_login(s.inner)

	fmt.Printf("inner: %+v\n", s.inner)

	//C.net_loop(s.inner)

	return nil
}

// Destroy destroys the state struct.
func (s *State) Destroy() {
	C.tgl_free_all(s.inner)
}

func (s *State) EnableCallbacks() {
	C.tgl_set_callback(s.inner, &C.upd_cb)
	C.tgl_set_timer_methods(s.inner, &C.tgl_libevent_timers)
	C.tgl_set_net_methods(s.inner, &C.tgl_conn_methods)
	var ev = C.event_base_new()
	C.tgl_set_ev_base(s.inner, unsafe.Pointer(ev))
}
