FROM golang:1.16-buster AS build

WORKDIR /

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

COPY ./api ./api

RUN go build -o /gserver

RUN ls -ltra /

EXPOSE 3000

ENTRYPOINT ["/gserver"]

# ## Deploy
# FROM alpine:latest

# # RUN useradd -ms /bin/bash nonroot 

# WORKDIR /

# COPY --from=build /gserver /gserver

# EXPOSE 3000

# # USER nonroot:nonroot

# RUN ls -ltra /

# RUN ["/gserver"]
