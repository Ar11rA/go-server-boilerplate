FROM golang:latest 
ADD . /srv/ 
WORKDIR /app 
RUN make build 
CMD ["make", "run"]