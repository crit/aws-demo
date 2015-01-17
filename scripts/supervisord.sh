#!/bin/sh

verifySupervisord() {
	set +e
	if which supervisord |& grep "no supervisord";
	then
		if easy_install |& grep "command not found";
		then
			set -e
			yum install -y python-setuptools
			verifySupervisord
		else
			easy_install -U supervisor
		fi
	fi
}

main() {
	verifySupervisord
}
main