#ifndef GO_TGL_H
#define GO_TGL_H

#include <tgl/tgl.h>
#include <tgl/tgl-timers.h>
#include <tgl/tgl-net.h>

// TODO: Could we use libevent version detection similar to autotools in Go?
#include <event2/event.h>
#include <event2/bufferevent.h>

// Defining typedef of the callbacks, we can encapsulate them if necessary.
// TODO: Are we going to use them, finally?
typedef void (new_msg_cb)(struct tgl_state *TLS, struct tgl_message *M);
typedef void (marked_read_cb)(struct tgl_state *TLS, int num, struct tgl_message *list[]);
typedef void (logprintf_cb)(const char *format, ...)  __attribute__ ((format (printf, 1, 2)));
typedef void (get_string_cb)(struct tgl_state *TLS, const char *prompt, int flags, void (*callback)(struct tgl_state *TLS, char *string, void *arg), void *arg);
typedef void (logged_in_cb)(struct tgl_state *TLS);
typedef void (started_cb)(struct tgl_state *TLS);
typedef void (type_notification_cb)(struct tgl_state *TLS, struct tgl_user *U, enum tgl_typing_status status);
typedef void (type_in_chat_notification_cb)(struct tgl_state *TLS, struct tgl_user *U, struct tgl_chat *C, enum tgl_typing_status status);
typedef void (type_in_secret_chat_notification_cb)(struct tgl_state *TLS, struct tgl_secret_chat *E);
typedef void (status_notification_cb)(struct tgl_state *TLS, struct tgl_user *U);
typedef void (user_registered_cb)(struct tgl_state *TLS, struct tgl_user *U);
typedef void (user_activated_cb)(struct tgl_state *TLS, struct tgl_user *U);
typedef void (new_authorization_cb)(struct tgl_state *TLS, const char *device, const char *location);
typedef void (chat_update_cb)(struct tgl_state *TLS, struct tgl_chat *C, unsigned flags);
typedef void (user_update_cb)(struct tgl_state *TLS, struct tgl_user *C, unsigned flags);
typedef void (secret_chat_update_cb)(struct tgl_state *TLS, struct tgl_secret_chat *C, unsigned flags);
typedef void (msg_receive_cb)(struct tgl_state *TLS, struct tgl_message *M);
typedef void (our_id_cb)(struct tgl_state *TLS, int id);
typedef void (notification_cb)(struct tgl_state *TLS, char *type, char *message);
typedef void (user_status_update_cb)(struct tgl_state *TLS, struct tgl_user *U);
typedef char *(create_print_name_cb) (struct tgl_state *TLS, tgl_peer_id_t id, const char *a1, const char *a2, const char *a3, const char *a4);

#endif // GO_TGL_H
