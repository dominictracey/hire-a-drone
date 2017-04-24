# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/dominictracey/rugby-scores

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get -v -d github.com/dominictracey/rugby-scores
RUN go install github.com/dominictracey/rugby-scores

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/rugby-scores --scheme=http --port=8080

# Document that the service listens on port 8080.
EXPOSE 8080
