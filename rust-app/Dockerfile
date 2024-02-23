FROM rust:1.76.0 as builder
RUN apt-get -y update
RUN apt-get install musl-tools -y
WORKDIR /usr/src/app
COPY Cargo.toml Cargo.lock ./
RUN mkdir src
COPY src ./src
RUN cargo install --target x86_64-unknown-linux-musl --path .

FROM alpine:latest
COPY --from=builder /usr/local/cargo/bin/rust-app /usr/local/bin/rust-app
EXPOSE 8080
CMD ["./rust-app"]

