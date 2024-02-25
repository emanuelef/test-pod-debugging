use actix_web::{get, middleware::Logger, web, App, HttpResponse, HttpServer, Responder};
use serde::Serialize;

#[derive(Debug, Serialize)]
struct GreetResponse {
    message: String,
}

#[get("/health")]
async fn health() -> impl Responder {
    HttpResponse::Ok().body("OK")
}

#[get("/hello/{name}")]
async fn greet(name: web::Path<String>) -> impl Responder {
    log::warn!("<---- /hello, name: {}", name);
    let response = GreetResponse {
        message: format!("Hello, {}!", name),
    };

    HttpResponse::Ok().json(serde_json::json!(response))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    std::env::set_var("RUST_LOG", "actix_web=info");
    env_logger::init();
    log::warn!("Starting web server on 0.0.0.0:8080");

    // Start the Actix web server
    HttpServer::new(|| App::new().wrap(Logger::default()).service(greet))
        .bind(("0.0.0.0", 8080))?
        .run()
        .await
}
