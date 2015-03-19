/*
Telegram Go library.
*/
package tgl

/*
#include "tgl.go.h"

// TODO: This is client logic, we need to move almost all of this into Go.
#include <fcntl.h>
#include <assert.h>

char *get_auth_key_filename(void) { return "auth"; }
char *get_state_filename(void) { return "state"; }

void empty_auth_file(struct tgl_state *TLS) {
  if (TLS->test_mode) {
    bl_do_dc_option (TLS, 1, 0, "", strlen (TG_SERVER_TEST_1), TG_SERVER_TEST_1, 443);
    bl_do_dc_option (TLS, 2, 0, "", strlen (TG_SERVER_TEST_2), TG_SERVER_TEST_2, 443);
    bl_do_dc_option (TLS, 3, 0, "", strlen (TG_SERVER_TEST_3), TG_SERVER_TEST_3, 443);
    bl_do_set_working_dc (TLS, 2);
  } else {
    bl_do_dc_option (TLS, 1, 0, "", strlen (TG_SERVER_1), TG_SERVER_1, 443);
    bl_do_dc_option (TLS, 2, 0, "", strlen (TG_SERVER_2), TG_SERVER_2, 443);
    bl_do_dc_option (TLS, 3, 0, "", strlen (TG_SERVER_3), TG_SERVER_3, 443);
    bl_do_dc_option (TLS, 4, 0, "", strlen (TG_SERVER_4), TG_SERVER_4, 443);
    bl_do_dc_option (TLS, 5, 0, "", strlen (TG_SERVER_5), TG_SERVER_5, 443);
    bl_do_set_working_dc (TLS, 4);
  }
}

void read_dc(struct tgl_state *TLS, int auth_file_fd, int id, unsigned ver) {
  int port = 0;
  assert (read (auth_file_fd, &port, 4) == 4);
  int l = 0;
  assert (read (auth_file_fd, &l, 4) == 4);
  assert (l >= 0 && l < 100);
  char ip[100];
  assert (read (auth_file_fd, ip, l) == l);
  ip[l] = 0;

  long long auth_key_id;
  static unsigned char auth_key[256];
  assert (read (auth_file_fd, &auth_key_id, 8) == 8);
  assert (read (auth_file_fd, auth_key, 256) == 256);

  bl_do_dc_option (TLS, id, 2, "DC", l, ip, port);
  bl_do_set_auth_key_id (TLS, id, auth_key);
  bl_do_dc_signed (TLS, id);
}

int need_dc_list_update;
void read_auth_file(struct tgl_state *TLS) {
  //if (binlog_enabled) { return; }
  int auth_file_fd = open(get_auth_key_filename(), O_CREAT | O_RDWR, 0600);
  if (auth_file_fd < 0) {
    empty_auth_file (TLS);
    return;
  }
  assert (auth_file_fd >= 0);
  unsigned x;
  unsigned m;
  if (read (auth_file_fd, &m, 4) < 4 || (m != DC_SERIALIZED_MAGIC)) {
    close (auth_file_fd);
    empty_auth_file (TLS);
    return;
  }
  assert (read (auth_file_fd, &x, 4) == 4);
  assert (x > 0);
  int dc_working_num;
  assert (read (auth_file_fd, &dc_working_num, 4) == 4);

  int i;
  for (i = 0; i <= (int)x; i++) {
    int y;
    assert (read (auth_file_fd, &y, 4) == 4);
    if (y) {
      read_dc (TLS, auth_file_fd, i, m);
    }
  }
  bl_do_set_working_dc (TLS, dc_working_num);
  int our_id;
  int l = read (auth_file_fd, &our_id, 4);
  if (l < 4) {
    assert (!l);
  }
  if (our_id) {
    bl_do_set_our_id (TLS, our_id);
  }
  close (auth_file_fd);
}

int unknown_user_list_pos;
int unknown_user_list[1000];

void write_state_file (struct tgl_state *TLS) {
  static int wseq;
  static int wpts;
  static int wqts;
  static int wdate;
  if (wseq >= TLS->seq && wpts >= TLS->pts && wqts >= TLS->qts && wdate >= TLS->date) { return; }
  wseq = TLS->seq; wpts = TLS->pts; wqts = TLS->qts; wdate = TLS->date;
  int state_file_fd = open (get_state_filename (), O_CREAT | O_RDWR, 0600);
  if (state_file_fd < 0) {
    logprintf ("Can not write state file '%s': %m\n", get_state_filename ());
    exit (1 ? EXIT_FAILURE : EXIT_SUCCESS);
  }
  int x[6];
  x[0] = STATE_FILE_MAGIC;
  x[1] = 0;
  x[2] = wpts;
  x[3] = wqts;
  x[4] = wseq;
  x[5] = wdate;
  assert (write (state_file_fd, x, 24) == 24);
  close (state_file_fd);
}

void net_loop(struct tgl_state *TLS) {
  int last_get_state = time (0);
  while (1) {
    event_base_loop (TLS->ev_base, EVLOOP_ONCE);

    if (time (0) - last_get_state > 3600) {
      tgl_do_lookup_state (TLS);
      last_get_state = time (0);
    }

    write_state_file (TLS);

    if (unknown_user_list_pos) {
      int i;
      for (i = 0; i < unknown_user_list_pos; i++) {
        tgl_do_get_user_info (TLS, TGL_MK_USER (unknown_user_list[i]), 0, 0, 0);
      }
      unknown_user_list_pos = 0;
    }
  }
}

#cgo CFLAGS: -I ./lib
#cgo LDFLAGS: -L ./lib/tgl/libs -l tgl
#cgo pkg-config: libevent
*/
import "C"

import (
  "unsafe"
  "fmt"
)

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

  hash := "34be6d99874fb9607fe932dbb86fe4a3"
  hptr := C.CString(hash)
  defer C.free(unsafe.Pointer(hptr))
  app := "Federator"
  aptr := C.CString(app)
  defer C.free(unsafe.Pointer(aptr))

  C.tgl_register_app_id(s.inner, 10604, hptr);
  C.tgl_set_app_version(s.inner, aptr);

  // Bug on C.tlg: If RSA file location doesn't exists it segfault at
  // tgl/mtproto-client.c:1263 (tglmp_on_start)
  C.tgl_init(s.inner)

  // Initial handshake.
  C.read_auth_file(s.inner)

  // reset_authorization
  //var P *C.struct_tgl_peer_t = C.tgl_peer_get (s.inner, TGL_MK_USER(s.inner.our_id))
  //C.set_default_username(P.user.phone);
  C.bl_do_reset_authorization(s.inner);

  C.tgl_login(s.inner)

  fmt.Printf("inner: %+v\n", s.inner)

  //C.net_loop(s.inner)

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
  var ev = C.event_base_new();
  C.tgl_set_ev_base(s.inner, unsafe.Pointer(ev));
}

