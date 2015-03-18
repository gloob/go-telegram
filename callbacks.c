#include <tgl/tgl.h>
#include "callbacks.h"
#include "_cgo_export.h"

void new_msg_cb_cgo(struct tgl_state *TLS, struct tgl_message *M) { NewMsgCB(TLS, M); }
void marked_read_cb_cgo(struct tgl_state *TLS, int num, struct tgl_message *list[]) { MarkedReadCB(TLS, num, list); }
//void logprintf_cb_cgo(const char *format, ...)  __attribute__ ((format (printf, 1, 2))) { LogPrintfCB((char *) format); }
void get_string_cb_cgo(struct tgl_state *TLS, const char *prompt, int flags, void (*callback)(struct tgl_state *TLS, char *string, void *arg), void *arg) { GetStringCB(TLS, (char *) prompt, flags); } // TODO: Add callback to the call.
void logged_in_cb_cgo(struct tgl_state *TLS) { LoggedInCB(TLS); }
void started_cb_cgo(struct tgl_state *TLS) { StartedCB(TLS); }
void type_notification_cb_cgo(struct tgl_state *TLS, struct tgl_user *U, enum tgl_typing_status status) { TypeNotificationCB(TLS, U, status); }
void type_in_chat_notification_cb_cgo(struct tgl_state *TLS, struct tgl_user *U, struct tgl_chat *C, enum tgl_typing_status status) { TypeInChatNotificationCB(TLS, U, C, status); }
void type_in_secret_chat_notification_cb_cgo(struct tgl_state *TLS, struct tgl_secret_chat *E) { TypeInSecretChatNotificationCB(TLS, E); }
void status_notification_cb_cgo(struct tgl_state *TLS, struct tgl_user *U) { StatusNotificationCB(TLS, U); }
void user_registered_cb_cgo(struct tgl_state *TLS, struct tgl_user *U) { UserRegisteredCB(TLS, U); }
void user_activated_cb_cgo(struct tgl_state *TLS, struct tgl_user *U) { UserActivatedCB(TLS, U); }
void new_authorization_cb_cgo(struct tgl_state *TLS, const char *device, const char *location) { NewAuthorizationCB(TLS, (char *) device, (char *) location); }
void chat_update_cb_cgo(struct tgl_state *TLS, struct tgl_chat *C, unsigned flags) { ChatUpdateCB(TLS, C, flags); }
void user_update_cb_cgo(struct tgl_state *TLS, struct tgl_user *C, unsigned flags) { UserUpdateCB(TLS, C, flags); }
void secret_chat_update_cb_cgo(struct tgl_state *TLS, struct tgl_secret_chat *C, unsigned flags) { SecretChatUpdateCB(TLS, C, flags); }
void msg_receive_cb_cgo(struct tgl_state *TLS, struct tgl_message *M) { MsgReceiveCB(TLS, M); }
void our_id_cb_cgo(struct tgl_state *TLS, int id) { OurIdCB(TLS, id); }
void notification_cb_cgo(struct tgl_state *TLS, char *type, char *message) { NotificationCB(TLS, type, message); }
void user_status_update_cb_cgo(struct tgl_state *TLS, struct tgl_user *U) { UserStatusUpdate(TLS, U); }
char *create_print_name_cb_cgo(struct tgl_state *TLS, tgl_peer_id_t id, const char *a1, const char *a2, const char *a3, const char *a4) { CreatePrintNameCB(TLS, id, (char *) a1, (char *) a2, (char *) a3, (char *) a4); }

struct tgl_update_callback upd_cb = {
  .new_msg = new_msg_cb_cgo,
  .marked_read = marked_read_cb_cgo,
  .logprintf = 0,
  .get_string = get_string_cb_cgo,
  .logged_in = logged_in_cb_cgo,
  .started = started_cb_cgo,
  .type_notification = type_notification_cb_cgo,
  .type_in_chat_notification = type_in_chat_notification_cb_cgo,
  .type_in_secret_chat_notification = type_in_secret_chat_notification_cb_cgo,
  .status_notification = status_notification_cb_cgo,
  .user_registered = user_registered_cb_cgo,
  .user_activated = user_activated_cb_cgo,
  .new_authorization = new_authorization_cb_cgo,
  .user_update = user_update_cb_cgo,
  .chat_update = chat_update_cb_cgo,
  .secret_chat_update = secret_chat_update_cb_cgo,
  .msg_receive = msg_receive_cb_cgo,
  .our_id = our_id_cb_cgo,
  .user_status_update = user_status_update_cb_cgo
};
