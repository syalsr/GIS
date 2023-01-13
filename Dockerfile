FROM golang:1.19 as build

WORKDIR /workspace
COPY go.* ./
RUN go mod tidy
RUN go mod download

COPY . .

RUN make all

RUN go build -o gis /workspace/cmd/GIS/main.go

FROM gcr.io/distroless/static-debian11
COPY --from=build /app/gis /
CMD ["/gis"]