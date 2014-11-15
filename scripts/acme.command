#!/usr/local/plan9/bin/rc

TERM=dumb
PAGER=nobs
SHELL=rc

killall plumber
plumber

exec acme -a &

