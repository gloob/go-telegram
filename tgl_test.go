package tgl

import (
  "testing"
)

func TestTglAuth(t *testing.T) {
  state := NewState()
  defer state.Destroy()

  err := state.Dial()
  if err != nil {
    t.Errorf("VERBOSE error: %s", err)
  }
}
