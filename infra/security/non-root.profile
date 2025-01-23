#include <tunables/global>

profile docker-infobaeapi flags=(attach_disconnected,mediate_deleted) {

  /app/** r,
  /app/bin/InfobaeAPI rx,
  /app/etc/** r,
  /app/logs/** w,

  network inet,

  capability chown,
  capability net_bind_service,

  deny /sys/**,
  deny /proc/**,
  deny /dev/**,
  deny /etc/**,
}