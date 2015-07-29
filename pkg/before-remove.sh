if [ $1 -eq 0 ]; then
  /sbin/stop http-logger >/dev/null 2>&1 || true
  if getent passwd http-logger >/dev/null ; then
    userdel http-logger
  fi

  if getent group http-logger > /dev/null ; then
    groupdel http-logger
  fi
fi
