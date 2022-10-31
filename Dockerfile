FROM golang:latest

COPY ./ /app/src/
COPY ./conf/ /app/conf/
WORKDIR /app/
RUN cd src/ && go mod download && go build -o ../service main.go && cd ../ && rm -rf src/

EXPOSE 9001
CMD ["./service"]