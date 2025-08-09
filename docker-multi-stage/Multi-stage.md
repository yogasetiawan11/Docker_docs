# What is Multi Stage in Docker 
A multi-stage Docker build is a feature that allows you to use multiple FROM instructions in a single Dockerfile. Each FROM instruction begins a new "stage" of the build. You can selectively copy artifacts like compiled code or assets—from one stage to another, discarding everything you don't need in the final image.


## Why Do We Need Them? 
The Problem with Single-Stage Builds Before multi-stage builds, a common problem was creating bloated Docker images. To build an application (like a Java, Go, or Node.js app), you needed a Docker image that contained:

- The base Operating System (e.g., Ubuntu)

- The language runtime (e.g., JDK, Go toolchain, Node.js)

- Build tools and dependencies (e.g., Maven, npm)

- Your application's source code

This resulted in:

Large Image Sizes: They have Large Images sizes Making them slow to push, pull, and deploy.

Increased Attack Surface: Every extra program or library in your image is a potential security vulnerability. or People find some kind of issues by This Images If you don't need the Java compiler (javac) in your production image, it shouldn't be there.

# What is Distroless Images 
Distroless is very minimalistic or lighweight Docker Images That contain only your application and its runtime application.

Distroless images and multi-stage builds are a perfect match for creating secure and minimal production containers. 

You can find Distroless Images in this github repository:
https://github.com/GoogleContainerTools/distroless/blob/main/README.md

Example
```bash
# ---- Stage 1: The "Builder" ----
# We use a full Go image to build our app.
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o my-app .

# ---- Stage 2: The "Final Image" with Distroless ----
# We now use a static, distroless image as our final base.
# 'gcr.io/distroless/static-debian12' is for a self-contained binary.
FROM gcr.io/distroless/static-debian12

WORKDIR /

# Copy ONLY our compiled application from the builder stage.
COPY --from=builder /app/my-app .

# Set the user to a non-root user for better security.
USER nonroot:nonroot

# Run the application.
CMD ["/my-app"]
```

By Implementing This concepts (Docker Multi Stage + Distroless Images) Your Images sizes will drastically go down, and The other Advantage your Images literally 99& not expose to vulnerability