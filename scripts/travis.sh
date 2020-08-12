#! /bin/bash

set -e

git config --global url.https://github.com/.insteadOf git@github.com:
git config --global url.https://rhysd@bitbucket.org/.insteadOf git@bitbucket.org:
echo -e "Host github.com\n\tVerifyHostKeyDNS no\n" >> ~/.ssh/config
echo -e "Host github.com\n\tStrictHostKeyChecking no\n" >> ~/.ssh/config
echo -e "Host bitbucket.org\n\tVerifyHostKeyDNS no\n" >> ~/.ssh/config
echo -e "Host bitbucket.org\n\tStrictHostKeyChecking no\n" >> ~/.ssh/config

if [[ "$TRAVIS_OS_NAME" == "osx" ]]; then
    brew update
    set +e
    brew upgrade go
    set -e
    go get -t -d -v ./...
    go test -v ./src/
else
    go get github.com/axw/gocov/gocov
    go get github.com/mattn/goveralls
    go get golang.org/x/tools/cmd/cover
    go get -t -d -v ./...
    go vet
    cd src/ && go vet && cd -
    go test -v -coverprofile=coverage.out ./src/
    if [[ "$COVERALLS_TOKEN" != "" ]]; then
        $HOME/gopath/bin/goveralls -coverprofile coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN
    fi
fi

