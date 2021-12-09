FROM golang:1.17

WORKDIR /transaction-processing-service/src/

CMD ["tail", "-f", "/dev/null"]