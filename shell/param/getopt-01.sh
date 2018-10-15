#!/usr/bin/env bash
## 处理复杂的参数需要使用getopt命令
###################################
# Extract command line options &amp; values with getopt
#
set -- $(getopt -q ab:cd "$@")
#
echo
while
 [ -n "$1" ]
do
  case "$1" in
  -a)
    echo "Found the -a option"
    ;;
  -b)
    param="$2"
    echo "Found the -b option, with parameter value $param"
    shift
    ;;
  -c)
     echo "Found the -c option"
     ;;
  --)
      shift
      break
      ;;
   *)
    echo "$1 is not option"
    ;;
esac
  shift
done

#
count=1
for param in "$@"
do
  echo "Parameter #$count : $param"
  count=$[ $count + 1 ]
done