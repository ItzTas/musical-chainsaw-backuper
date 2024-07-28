#!/usr/bin/env sh

src=""
dst=""

while getopts s:d: opt; do
	case ${opt} in
	s)
		src=$OPTARG
		;;
	d)
		dst=$OPTARG
		;;
	*)
		echo "Uso: $0 -s <source_directory> -d <destination_directory>"
		;;
	esac
done
shift $((OPTIND - 1))

if [ -z "$src" ] || [ -z "$dst" ]; then
	echo "Error: origin and destination directories are required."
	echo "Uso: $0 -s <source_directory> -d <destination_directory>"
	exit 1
fi

if [ -d "musical-chainsaw-backuper" ]; then
	cd "musical-chainsaw-backuper" || printf ""
fi

go build -o out
if [ -x "./out" ]; then
	./out -src="$src" -dst="$dst"
	rm ./out
else
	echo "fail in building binary"
	exit 1
fi
