# Build go-consent
FROM library/golang:alpine AS node
RUN apk update && apk add bash curl git gcc libc-dev

# Godep for vendoring
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/github.com/sqsinformatique/backend
RUN mkdir -p $APP_DIR

# Compile the binary and statically link
ADD . $APP_DIR
WORKDIR $APP_DIR
# RUN CGO_ENABLED=0 dep ensure -update
# RUN go get -u github.com/gobuffalo/packr/packr
# RUN packr
RUN CGO_ENABLED=0 ./mng.sh -b

FROM scratch

COPY --from=0 /go/src/github.com/sqsinformatique/backend/backend /usr/bin/backend

USER 1000

# Set the entrypoint
ENTRYPOINT ["backend"]
