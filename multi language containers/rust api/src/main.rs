use actix_web::middleware::Logger;
use actix_web::{App, HttpResponse, HttpServer, Responder, get};
use env_logger::Env;
use serde::Serialize;

#[derive(Serialize)]
struct Message {
    message: String,
}

#[get("/favicon.ico")]
async fn favicon() -> impl Responder {
    HttpResponse::NoContent().finish()
}

#[get("/")]
async fn hello() -> impl Responder {
    let msg = Message {
        message: "Welcome to the Rust actix-web API!".to_string(),
    };

    HttpResponse::Ok().json(msg)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init_from_env(Env::default().default_filter_or("info"));

    HttpServer::new(|| {
        App::new()
            .service(hello)
            .wrap(Logger::new("%a %{User-Agent}i"))
    })
    .bind(("0.0.0.0", 8080))?
    .run()
    .await
}
