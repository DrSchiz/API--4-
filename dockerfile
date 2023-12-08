# The base go-image
FROM golang:1.14-alpine
 
# Create a directory for the app
 
# Copy all files from the current directory to the app directory
COPY . .
 
# Set working directory
WORKDIR /API
 
# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o API . 
 
# Run the server executable
CMD [ "/API/cmd" ]