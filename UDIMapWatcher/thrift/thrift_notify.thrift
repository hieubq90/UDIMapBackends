namespace java vng.notify.firebasecloudmessage
namespace cpp vng.up.ubus.service.notify
namespace go ub.service.notify

struct TKeyValue {
  1: string key;
  2: string value;
}

struct TNotificationPayload {
  1: string title;
  2: string body;
  3: optional string icon;
  4: optional list<TKeyValue> data;
  5: optional string click_action;
}

struct TDataPayload {
  1: list<TKeyValue> data;
}

struct TFCMMessage {
  1: TNotificationPayload notiPayload;
  2: TDataPayload dataPayload;
}

struct TResponse {
  1: i32 statusCode;
  2: string header;
  3: string body;
}

typedef string TDeviceToken
typedef list<TDeviceToken> TDeviceTokenList

service NotifyService {
  TResponse notiToDeviceToken(1: string appServerKey, 2: TFCMMessage message, 3: TDeviceToken deviceToken);

  TResponse notiToMultiDeviceToken(1: string appServerKey, 2: TFCMMessage message, 3: TDeviceTokenList tokenList);

  TResponse notiToTopic(1: string appServerKey, 2: string topic, 3: string condition, 4: TFCMMessage message);
}
