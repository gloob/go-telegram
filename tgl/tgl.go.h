#ifndef GO_TGL_H
#define GO_TGL_H

#include <tgl/tgl.h>
#include <tgl/tgl-timers.h>
#include <tgl/tgl-net.h>
#include <tgl/tgl-binlog.h>
#include <tgl/config.h>

#ifdef EVENT_V2
#include <event2/event.h>
#else
#include <event.h>
#include "event-old.h"
#endif

#include <fcntl.h>
#include <stdio.h>

#define DC_SERIALIZED_MAGIC 0x868aa81d
#define STATE_FILE_MAGIC 0x28949a93
#define SECRET_CHAT_FILE_MAGIC 0x37a1988a

extern struct tgl_update_callback upd_cb;

// Wrapping some useful C macros.
tgl_peer_id_t macro_TGL_MK_USER(int id);

#endif // GO_TGL_H
