FROM golang:1.11

RUN mkdir -p /go/src/github.com/beoneill/gokittest

ADD . /go/src/github.com/beoneill/gokittest

#RUN echo "[url \"git@bitbucket.org:beoneill\"]\n\tinsteadOf = https://beoneill@bitbucket.org/" >> /root/.gitconfig
#
#ARG SSH_PRIVATE_KEY
#RUN mkdir /root/.ssh/
#RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa

RUN go get  -t -v ./...

RUN ls -al

RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/beoneill/gokittest/flights/cmd  -watch github.com/beoneill/gokittest/flights
