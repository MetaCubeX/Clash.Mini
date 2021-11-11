#!/usr/bin/sh

if ! [ `echo $1 | egrep '^(([0-9]*f)|(n[0-9]+))$'` ]
then
  set -o errexit
  if [ -n "$1" ] ;then
    echo -e "Format error: \"$1\"\n"
  fi
  echo -e 'Example: \n    ./ctailf.sh 10f \n    -- real-timely print log with format the color, show last %n line(s)\n'
  echo -e '    ./ctailf.sh n10\n    -- print log with format the color, show last %n line(s)\n'
  echo -e '    ./ctailf.sh n10 log.txt\n    -- print log with format the color, show last %n line(s)'
  exit 1
fi

if [ ! -n "$2" ] ;then
  curDir=$(dirname $(readlink -f "$0"))
  logPath="$curDir/Clash.Mini.log"
else
  logPath="$2"
fi

tail -$1 $logPath | perl -pe 's/(^.*DEBG.*$)|(^.*INFO.*$)|(^.*WARN.*$)|(^.*EROR.*$)/\e[1;32m$1\e[0m\e[1;36m$2\e[0m\e[1;33m$3\e[0m\e[1;31m$4\e[0m/g'
#tail -$1 $logPath | perl -pe 's/(DEBG)|(INFO)|(WARN)|(EROR)/\e[1;32m$1\e[0m\e[1;36m$2\e[0m\e[1;33m$3\e[0m\e[1;31m$4\e[0m/g'
