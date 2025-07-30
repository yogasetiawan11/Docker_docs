FROM ubuntu:latest

WORKDIR /app

COPY . /app

RUN apt-get update && apt-get install python -y install python3 python-pip

ENV NAME world

CMD ["python3", "app.py"]
