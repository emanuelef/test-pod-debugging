use actix_web::{get, web, App, HttpResponse, HttpServer, Responder};
use serde::Serialize;

#[derive(Debug, Serialize)]
struct GreetResponse {
    message: String,
}

#[get("/hello/{name}")]
async fn greet(name: web::Path<String>) -> impl Responder {
    let response = GreetResponse {
        message: format!("Hello, {}!", name),
    };

    HttpResponse::Ok().json(serde_json::json!(response))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(greet))
        .bind(("0.0.0.0", 8080))?
        .run()
        .await
}
