#!/bin/sh

is_dir() {
	local dir=$1

	[[ -d $dir ]]
}

set +e
supervisorctl stop awsdemo

is_dir /home/ec2-user/awsdemo && rm -rf /home/ec2-user/awsdemo
rm -f /etc/rc.d/init.d/awsdemo