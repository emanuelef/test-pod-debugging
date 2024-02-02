FROM rust:latest as builder
WORKDIR /usr/src/app
COPY Cargo.toml Cargo.lock ./
RUN mkdir src
COPY src ./src
RUN cargo build --release

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /usr/src/app/target/release/rust-app .
EXPOSE 8080
CMD ["./rust-app"]