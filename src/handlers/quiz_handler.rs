use axum::{
    extract::State,
    http::header,
    response::{AppendHeaders, Html, IntoResponse},
};
use tera::Context;
use tower_cookies::Cookies;
use uuid::Uuid;

use crate::{config::TEMPLATES, AppState};

/// quiz 首页
pub async fn index(State(state): State<AppState>, cookies: Cookies) -> impl IntoResponse {
    let quiz_id = cookies.get("quiz_id").map_or_else(
        || {
            let id = Uuid::new_v4().to_string();
            tracing::info!("generate a new quiz_id:{}", &id);
            id
        },
        |cookie| cookie.value().to_string(),
    );
    let cookie = format!("quiz_id={}; Path=/class7; HttpOnly; Max-Age=10800", quiz_id);
    let headers = AppendHeaders([(header::SET_COOKIE, cookie)]);
    let mut context = Context::new();
    context.insert("quiz_url", &format!("/class7/quiz/{}", quiz_id));
    let html = TEMPLATES.render("class7/index.html", &context);
    match html {
        Ok(t) => (headers, Html(t)),
        Err(e) => (headers, Html(format!("错误: {}", e))),
    }
}
