FROM golang:1.20-alpine

# create directory folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable file with name "AIRNBN"
RUN go build -o airnbn

# run executable file
CMD ["./airnbn"]