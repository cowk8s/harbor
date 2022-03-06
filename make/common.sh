#!/bin/bash

set +e
set -o noglob

#
# Set Colors
#
bold=$(tput bold)
underline=$(tput sgr 0 1)
reset=$(tput sgr0)

red=$(tput setaf 1)
green=$(tput setaf 76)
white=$(tput setaf 7)
tan=$(tput setaf 202)
blue=$(tput setaf 25)

warn() { printf "${tan}➜ %s${reset}\n" "$@"
}
note() { printf "\n${underline}${bold}${blue}Note:${reset} ${blue}%s${reset}\n" "$@"
}

set -e

function check_golang {
	if ! go version &> /dev/null
	then
		warn "No golang package in your enviroment. You should use golang docker image build binary."
		return
	fi

	# docker has been installed and check its version
	if [[ $(go version) =~ (([0-9]+)\.([0-9]+)([\.0-9]*)) ]]
	then
		golang_version=${BASH_REMATCH[1]}
		golang_version_part1=${BASH_REMATCH[2]}
		golang_version_part2=${BASH_REMATCH[3]}

		# the version of golang does not meet the requirement
		if [ "$golang_version_part1" -lt 1 ] || ([ "$golang_version_part1" -eq 1 ] && [ "$golang_version_part2" -lt 12 ])
		then
			warn "Better to upgrade golang package to 1.12.0+ or use golang docker image build binary."
			return
		else
			note "golang version: $golang_version"
		fi
	else
		warn "Failed to parse golang version."
		return
	fi
}