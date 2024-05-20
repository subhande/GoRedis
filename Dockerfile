
# Specifies a parent image
FROM golang:1.22.3-bookworm

# RUN apt install netcat-traditional
# RUN apt install curl
# RUN curl -fsSL https://packages.redis.io/gpg | gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
# RUN echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/redis.list
# RUN apt-get update
# RUN apt-get install redis



# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod download

# Builds your app with optional configuration
RUN go build -o /godocker

# Tells Docker which network port your container listens on
EXPOSE 7379

# # Specifies the executable command that runs when the container starts
CMD [ "/godocker" ]

# CMD ["go", "run", "main.go"]