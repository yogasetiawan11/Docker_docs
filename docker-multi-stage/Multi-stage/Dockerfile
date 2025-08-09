#####################
# Docker base images
#####################

# We spesify our operating system and make it Alias for sync 
FROM ubuntu AS build

# I am downloading GO web app
RUN apt-get update && apt-get install -y golang-go

# I copy this source code (Go app) in this spesific thing
COPY . .

# I am running Go build and building binary
RUN CGO_ENABLED=0 go build -o /app .

# And I execute the binary
ENTRYPOINT ["/app"]