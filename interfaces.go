package tgl

// #include <tgl/tgl.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type Interface struct {
	state *State
	cb    *C.struct_tgl_update_callback
	net   *C.struct_tgl_net_methods
	timer *C.struct_tgl_timer_methods
	ev    *C.void
}

// TODO: Slowly moving the interfaces logic here.

/*
// Callbacks
void tgl_set_callback (struct tgl_state *TLS, struct tgl_update_callback *cb);
void tgl_set_net_methods (struct tgl_state *TLS, struct tgl_net_methods *methods);
void tgl_set_timer_methods (struct tgl_state *TLS, struct tgl_timer_methods *methods);
void tgl_set_ev_base (struct tgl_state *TLS, void *ev_base);
*/

func NewInterface(state *State) (*Interface, error) {
	fmt.Printf("Calling NewInterface, let's create some methods.\n")

	/*
	   i := &Interface{
	     state: state,
	     cb: cb,
	     net: net,
	     timer: timer,
	     ev: ev,
	   }

	   return i, nil
	*/

	return nil, nil
}

func (i *Interface) Callback() {
	C.tgl_set_callback(i.state.inner, i.cb)
}

func (i *Interface) Net() {
	C.tgl_set_net_methods(i.state.inner, i.net)
}

func (i *Interface) Timer() {
	C.tgl_set_timer_methods(i.state.inner, i.timer)
}

func (i *Interface) Ev() {
	C.tgl_set_ev_base(i.state.inner, unsafe.Pointer(i.ev))
}

func (i *Interface) Destroy() {
}
