FROM golang:1.22-alpine3.18

LABEL Maintainer="Muhammet Hadi KAMAT"

# set working directory
WORKDIR /app

# Copy the source code
COPY . . 

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o api .

#EXPOSE the port
EXPOSE 3000

# Run the executable
CMD ["./api"]