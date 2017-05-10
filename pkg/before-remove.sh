if [ $1 -eq 0 ]; then
  systemctl disable http-logger >/dev/null || true
  systemctl stop http-logger >/dev/null || true
  if getent passwd http-logger >/dev/null ; then
    userdel http-logger
  fi

  if getent group http-logger > /dev/null ; then
    groupdel http-logger
  fi
fi
