FROM ubuntu:latest

# Set working directory in the Image
WORKDIR /app

# Copy the files from the host file system to The IMAGE file system
COPY . /app

# Install all necessary package, in this case we try run app python that require python
RUN apt-get update && apt-get install python -y install python3 python-pip

# Set environment variables
ENV NAME world

# Run a Command to start the application 
CMD ["python3", "app.py"]
