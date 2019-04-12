FROM damienfontaine/lunarc:0.0.5
LABEL maintainer="Damien Fontaine <damien.fontaine@lineolia.net>"

ENV LUNARC_ENV production

ADD . /go/src/gitlab.lineolia.net/meda/voice-corpus-maker
RUN go install gitlab.lineolia.net/meda/voice-corpus-maker

WORKDIR /go/src/gitlab.lineolia.net/meda/voice-corpus-maker

ENTRYPOINT /go/bin/voice-corpus-maker -env=${LUNARC_ENV}