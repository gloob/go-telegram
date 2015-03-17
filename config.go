package tgl

// #cgo CFLAGS: -I ./lib
// #cgo LDFLAGS: -L ./lib/tgl/libs -l tgl
// #include "tgl.go.h"
// #include "stdlib.h"
import "C"
import "unsafe"

type Config struct {
  state *State
}

/*
// Config set
void tgl_set_binlog_mode (struct tgl_state *TLS, int mode);
- void tgl_set_binlog_path (struct tgl_state *TLS, const char *path);
- void tgl_set_auth_file_path (struct tgl_state *TLS, const char *path);
- void tgl_set_download_directory (struct tgl_state *TLS, const char *path);
- void tgl_set_rsa_key (struct tgl_state *TLS, const char *key);
- void tgl_set_app_version (struct tgl_state *TLS, const char *app_version);

// Operation modes
x void tgl_incr_verbosity (struct tgl_state *TLS);
- void tgl_set_verbosity (struct tgl_state *TLS, int val);
- void tgl_enable_pfs (struct tgl_state *TLS);
- void tgl_set_test_mode (struct tgl_state *TLS);
*/

// NewState creates a new State struct.
func NewConfig(state *State) *Config {
  c := new(Config)
  c.state = state

  return c
}

func (c *Config) setBinlogMode(mode int) {
  C.tgl_set_binlog_mode(c.state.inner, C.int(mode))
}

func (c *Config) setBinlogPath(path string) {
  pptr := C.CString(path)
  defer C.free(unsafe.Pointer(pptr))

  C.tgl_set_binlog_path(c.state.inner, pptr)
}

func (c *Config) setAuthPath(path string) {
  pptr := C.CString(path)
  defer C.free(unsafe.Pointer(pptr))

  C.tgl_set_auth_file_path(c.state.inner, pptr)
}

func (c *Config) setDownloadPath(path string) {
  pptr := C.CString(path)
  defer C.free(unsafe.Pointer(pptr))

  C.tgl_set_download_directory(c.state.inner, pptr)
}

func (c *Config) setRsaKey(key string) {
  kptr := C.CString(key)
  defer C.free(unsafe.Pointer(kptr))

  C.tgl_set_rsa_key(c.state.inner, kptr)
}

func (c *Config) setAppVersion(version string) {
  vptr := C.CString(version)
  defer C.free(unsafe.Pointer(vptr))

  C.tgl_set_app_version(c.state.inner, vptr)
}

func (c *Config) setTestMode() {
  C.tgl_set_test_mode(c.state.inner)
}

func (c *Config) setVerboseMode(val int) {
  C.tgl_set_verbosity(c.state.inner, C.int(val))
}

func (c *Config) setPfs() {
  C.tgl_enable_pfs(c.state.inner)
}

func (c *Config) Destroy() {
}
