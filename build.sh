#!/bin/bash
sum="sha1sum"

# VERSION=`date -u +%Y%m%d`
VERSION="v1.1.2"
LDFLAGS="-X main.VERSION=$VERSION -s -w"
GCFLAGS=""

if ! hash sha1sum 2>/dev/null; then
	if ! hash shasum 2>/dev/null; then
		echo "I can't see 'sha1sum' or 'shasum'"
		echo "Please install one of them!"
		exit
	fi
	sum="shasum"
fi

UPX=false
if hash upx 2>/dev/null; then
	UPX=true
fi





OSES=( linux darwin windows )
ARCHS=(amd64 386 )
for os in ${OSES[@]}; do
	for arch in ${ARCHS[@]}; do
		suffix=""
		if [ "$os" == "windows" ]
		then
			suffix=".exe"
		fi

		env CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o ./dist/${os}_${arch}/wasp${suffix} ./
		env CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o ./dist/${os}_${arch}/dummy-server${suffix} ./dummy-server
	
    	if $UPX; then upx -9 client_${os}_${arch}${suffix} server_${os}_${arch}${suffix};fi
		# tar -zcf ./dist/wasp-${os}-${arch}-$VERSION.tar.gz ./dist/${os}_${arch}/proxy${suffix} ./dist/${os}_${arch}/player${suffix} ./dist/${os}_${arch}/recorder${suffix}
        cd dist/${os}_${arch}/
        zip -D -q -r ../wasp-${os}-${arch}-$VERSION.zip wasp${suffix} dummy-server${suffix}
        cd ../..
    	$sum ./dist/wasp-${os}-${arch}-$VERSION.zip
	done
done
