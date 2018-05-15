WORDS_ALPHA='https://github.com/dwyl/english-words/raw/master/words_alpha.zip'

all: setup

.PHONY: setup
setup:
	[ ! -f words_alpha.txt ] && wget -Owords_alpha.zip $(WORDS_ALPHA) && unzip words_alpha.zip && rm words_alpha.zip

.PHONY: coverage
coverage:
	@rm coverage.out || true
	go get -u github.com/haya14busa/goverage
	'${GOPATH}/bin/goverage' -v -coverprofile coverage.out ./internal/game ./internal/wordbank
	go tool cover -html=coverage.out
