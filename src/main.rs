use sqlx::SqlitePool;

mod handlers;
mod models;
mod routes;
mod config;

#[tokio::main]
async fn main() {
    // 初始化 tracing
    config::init_tracing();
    // 初始化tera模版引擎
    let tera = config::init_tera();
    // 初始化数据库
    let pool = config::init_database().await.unwrap();
    // 初始化应用状态
    let app_state = AppState {
        pool,
        quiz_ids: vec![],
        tera,
    };
}

#[derive(Clone)]
struct AppState {
    pool: SqlitePool,
    quiz_ids: Vec<String>,
    tera: tera::Tera,
}