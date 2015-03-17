package tgl

// #cgo CFLAGS: -I ./lib
// #cgo LDFLAGS: -L ./lib/tgl/libs -l tgl
// #include "tgl.go.h"
import "C"

import (
  "unsafe"
  "fmt"
)

type Interface struct {
  state *State
  cb *C.struct_tgl_update_callback
  net *C.struct_tgl_net_methods
  timer *C.struct_tgl_timer_methods
  ev *C.void
}

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
  cb C.struct_tgl_update_callback = {
    .new_msg = print_message_gw,
    .marked_read = mark_read_upd,
    .logprintf = logprintf,
    .get_string = do_get_string,
    .logged_in = on_login,
    .started = on_started,
    .type_notification = type_notification_upd,
    .type_in_chat_notification = type_in_chat_notification_upd,
    .type_in_secret_chat_notification = 0,
    .status_notification = 0,
    .user_registered = 0,
    .user_activated = 0,
    .new_authorization = 0,
    .user_update = user_update_gw,
    .chat_update = chat_update_gw,
    .secret_chat_update = secret_chat_update_gw,
    .msg_receive = print_message_gw,
    .our_id = our_id_gw,
    .user_status_update = user_status_upd
  }
  */
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

// Destroy destroys the Poller
func (i *Interface) Destroy() {
}
