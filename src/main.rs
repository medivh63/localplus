use std::path::PathBuf;

use axum::{routing::get_service, Router};
use config::AppState;
use tower_cookies::CookieManagerLayer;
use tower_http::{services::ServeDir, trace::TraceLayer};
#[macro_use]
extern crate lazy_static;

mod config;
mod handlers;
mod models;
mod routes;

#[tokio::main]
async fn main() {
    // 初始化 tracing
    config::init_tracing();
    // 初始化数据库
    let pool = config::init_database().await.unwrap();
    // 查询所有题目id
    let quiz_ids = models::get_all_question_ids(&pool).await.unwrap();
    // 初始化应用状态
    let state = AppState::new(pool, quiz_ids);
    // 启动服务器
    let app = Router::new()
        .nest("/class7", routes::driving_quiz_routes())
        // 将这行放在其他路由之前，以避免路由冲突
        .nest_service(
            "/static",
            get_service(ServeDir::new(PathBuf::from("static"))),
        )
        .fallback(handlers::fallback)
        .layer(CookieManagerLayer::new())
        .layer(TraceLayer::new_for_http())
        .with_state(state);
    // run our app with hyper, listening globally on port
    let port = std::env::var("SERVER_PORT").unwrap_or_else(|_| "3000".to_string());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{}", port))
        .await
        .unwrap();
    tracing::info!("Server running on: {}", listener.local_addr().unwrap());

    axum::serve(listener, app).await.unwrap()
}
