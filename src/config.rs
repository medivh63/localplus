use sqlx::{pool::PoolOptions, Error, SqlitePool};
use tera::Tera;

// 全局 Tera 实例
lazy_static! {
    pub static ref TEMPLATES: Tera = {
        tracing::info!("init tera ...");
        let mut tera = match Tera::new("templates/**/*.html") {
            Ok(t) => t,
            Err(e) => {
                tracing::error!("init tera error: {}", e);
                Tera::default()
            }
        };
        tera.autoescape_on(vec![".html"]);
        tera
    };
}

/// 初始化 tracing
pub fn init_tracing() {
    tracing_subscriber::fmt()
        .with_max_level(tracing::Level::TRACE)
        .with_file(true)
        .init();
}

/// 初始化数据库
pub async fn init_database() -> Result<SqlitePool, Error> {
    let database_url = dotenv::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let pool = PoolOptions::new()
        .max_connections(35)
        .connect(&database_url)
        .await
        .expect("db connect error");
    Ok(pool)
}
