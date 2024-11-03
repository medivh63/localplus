use std::sync::Arc;

use axum::Router;
use sqlx::SqlitePool;
use tower_cookies::CookieManagerLayer;
use tower_http::trace::TraceLayer;
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
    // 初始化应用状态
    let state = AppState::new(pool, Vec::new());
    // 启动服务器
    let app = Router::new()
        .nest("/class7", routes::driving_quiz_routes())
        .fallback(handlers::fallback)
        .layer(CookieManagerLayer::new())
        .layer(TraceLayer::new_for_http())
        .with_state(state);
    // run our app with hyper, listening globally on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    tracing::info!("Server running on: {}", listener.local_addr().unwrap());
    axum::serve(listener, app).await.unwrap()
}

#[derive(Clone)]
struct AppState {
    pool: Arc<SqlitePool>,
    quiz_ids: Arc<Vec<String>>,
}

impl AppState {
    pub fn new(pool: SqlitePool, quiz_ids: Vec<String>) -> Self {
        Self {
            pool: Arc::new(pool),
            quiz_ids: Arc::new(quiz_ids),
        }
    }

    // 只提供getter方法
    pub fn pool(&self) -> &SqlitePool {
        &self.pool
    }

    pub fn quiz_ids(&self) -> &Vec<String> {
        &self.quiz_ids
    }
}
