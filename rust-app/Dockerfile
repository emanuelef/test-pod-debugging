FROM lukemathwalker/cargo-chef:latest-rust-1.76-alpine AS chef
WORKDIR /app

FROM chef AS planner
COPY . .
RUN cargo chef prepare --recipe-path recipe.json

FROM chef AS builder 
RUN apk update && apk add musl-dev perl make
COPY --from=planner /app/recipe.json recipe.json
RUN cargo chef cook --release --recipe-path recipe.json
COPY . .
RUN SQLX_OFFLINE=true cargo build --release

FROM alpine:latest AS runtime
RUN apk update && apk add ca-certificates && apk cache clean
COPY --from=builder /app/target/release/rust-app /usr/local/bin/rust-app
WORKDIR /usr/local/rust-app
ENV RUST_LOG="info"
EXPOSE 8080
CMD ["rust-app"]