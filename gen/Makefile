
.PHONY: adapter advertising

adapter:
	BASEDIR=.. \
	FILE_FILTER=adapter \
	xAPI_FILTER="Provisioner" \
	xMETHOD_FILTER=Attach LOG_LEVEL=trace \
		go run ./srcgen/main.go full --debug

advertising:
	BASEDIR=.. \
	FILE_FILTER=advertising-api \
	xAPI_FILTER="Provisioner" \
	xMETHOD_FILTER=Attach LOG_LEVEL=trace \
	go run ./srcgen/main.go full --debug

mesh:
	BASEDIR=.. \
	FILE_FILTER=advertising-api \
	xAPI_FILTER="Provisioner" \
	xMETHOD_FILTER=Attach LOG_LEVEL=trace \
	go run ./srcgen/main.go full --debug