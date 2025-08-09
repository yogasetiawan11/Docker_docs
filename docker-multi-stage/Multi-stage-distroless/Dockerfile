FROM ubuntu AS build

RUN apt-get update && apt-get install -y golang-go

COPY . .

RUN CGO_ENABLED=0 go build -o /app .

ENTRYPOINT ["/app"]


#########################
#
#########################

# Scratch only works for GO Language
# Because GO doesn't need GO Runtime If you want to use other Language you can
# choose Distroless Images In link GitHub that I've Provided
FROM scratch

# Copy the compiled binary from the build stage above
COPY --from=build /app /app

# Set entrypoint for the container to run the binary
ENTRYPOINT ["/app"]
