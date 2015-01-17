#!/bin/sh
set +e
DIR="$( cd "$( dirname "$0" )" && pwd )"

source /root/.bashrc

hostname awsdemo
service network restart

yes | cp ${DIR}/../etc/rc.d/init.d/awsdemo /etc/rc.d/init.d
chmod a+x /etc/rc.d/init.d/awsdemo
/sbin/chkconfig --level 235 awsdemo on

yes | cp ${DIR}/../etc/supervisord.conf /etc/

service awsdemo start