FROM rust:1.75.0 as builder
WORKDIR /usr/src/app
COPY Cargo.toml Cargo.lock ./
RUN mkdir src
COPY src ./src
RUN cargo build --release

FROM rust:1.75.0
WORKDIR /app
COPY --from=builder /usr/src/app/target/release/rust-app .
EXPOSE 8080
CMD ["./rust-app"]