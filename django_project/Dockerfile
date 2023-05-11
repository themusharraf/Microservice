FROM python:alpine

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV MYPATH=/home/app/code
RUN apk add --update py-pip

WORKDIR $MYPATH
COPY requirements.txt $MYPATH
RUN pip install -r requirements.txt
COPY . $MYPATH