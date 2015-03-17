package tgl

/*
#include "tgl.go.h"

void new_msg_cb_cgo(struct tgl_state *TLS, struct tgl_message *M) { NewMsgCB(TLS, M); }
void marked_read_cb_cgo(struct tgl_state *TLS, int num, struct tgl_message *list[]) { MarkedReadCB(TLS, num, list); }

struct tgl_update_callback upd_cb = {
  .new_msg = new_msg_cb_cgo,
  .marked_read = marked_read_cb_cgo,
  .logprintf = 0,
  .get_string = 0,
  .logged_in = 0,
  .started = 0,
  .type_notification = 0,
  .type_in_chat_notification = 0,
  .type_in_secret_chat_notification = 0,
  .status_notification = 0,
  .user_registered = 0,
  .user_activated = 0,
  .new_authorization = 0,
  .user_update = 0,
  .chat_update = 0,
  .secret_chat_update = 0,
  .msg_receive = 0,
  .our_id = 0,
  .user_status_update = 0
};

#cgo CFLAGS: -I ./lib
#cgo LDFLAGS: -L ./lib/tgl/libs -l tgl
#cgo pkg-config: libevent
*/
import "C"
import "unsafe"

// State wraps a tgl_state struct.
type State struct {
  inner *C.struct_tgl_state
}

// NewState creates a new State struct.
func NewState() *State {
  return &State{C.tgl_state_alloc()}
}

func (s *State) Dial () *error{
  config := NewConfig(s)
  defer config.Destroy()

  config.setRsaKey("./rsa.pub")
  s.EnableCallbacks()

  hash := "36722c72256a24c1225de00eb6a1ca74"
  hptr := C.CString(hash)
  defer C.free(unsafe.Pointer(hptr))
  app := "gotgl 0.1"
  aptr := C.CString(app)
  defer C.free(unsafe.Pointer(aptr))

  C.tgl_register_app_id(s.inner, 2899, hptr);
  C.tgl_set_app_version(s.inner, aptr);

  //C.tgl_init(s.inner)

  return nil
}

// Destroy destroys the state struct.
func (s *State)  Destroy() {
  C.tgl_free_all(s.inner)
}

func (s *State) EnableCallbacks() {
  C.tgl_set_callback(s.inner, &C.upd_cb)
  C.tgl_set_timer_methods(s.inner, &C.tgl_libevent_timers)
  C.tgl_set_net_methods(s.inner, &C.tgl_conn_methods)
  //ev *C.struct_event_base = C.event_base_new();
  //C.tgl_set_ev_base(s.inner, ev);
}

