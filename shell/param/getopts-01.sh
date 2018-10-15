#!/bin/bash
# 函数用法 getopts option_string varname
# 该函数不能处理长参数
#当option_string以":"开头时，getopts会区分invalid option错误和miss option argument错误。invalid option时，varname会被设成?，$OPTARG是出问题的option； miss option argument时，varname会被设成:，$OPTARG是出问题的option。
# 如果option_string不以":"开头，invalid option错误和miss option argument错误都会使varname被设成?，$OPTARG是出问题的option。

# 情况一: 不以冒号开头
# 输入 -a 但是后面没有参数的的时候，会报错误
while getopts "a:" opt; do
  case $opt in
    a)
      echo "[case1] this is -a the arg is ! $OPTARG"
      ;;
    \?)
      echo "[case1] Invalid option: -$OPTARG"
      ;;
  esac
done
# 情况二:以冒号开头
# 忽略错误
while getopts ":a:" opt; do
  case $opt in
    a)
      echo "[case2] this is -a the arg is ! $OPTARG"
      ;;
    \?)
      echo "[case2] Invalid option: -$OPTARG"
      ;;
  esac
done