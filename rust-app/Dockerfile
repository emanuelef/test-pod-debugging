FROM rust:latest as builder
WORKDIR /usr/src/app
COPY Cargo.toml Cargo.lock ./
RUN mkdir src && echo "fn main() {println!(\"Hello, world!\");}" > src/main.rs
RUN cargo build --release
COPY src ./src
RUN cargo build --release

FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /usr/src/app/target/release/rust-app .
EXPOSE 8080
CMD ["./rust-app"]
