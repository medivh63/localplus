use axum::{
    http::header,
    response::{AppendHeaders, Html, IntoResponse},
};
use tera::Context;
use tower_cookies::Cookies;
use uuid::Uuid;

/// quiz 首页
pub async fn index(cookies: Cookies) -> impl IntoResponse {
    let practice_id = cookies.get("quiz_id").map_or_else(
        || {
            // 如果没有cookie则生成新的 quiz_id
            let id = Uuid::new_v4().to_string();
            tracing::info!("generate a new practice_id:{}", &id);
            id
        },
        |cookie| cookie.value().to_string(),
    );
    let cookie = format!(
        "quiz_id={}; Path=/class7; HttpOnly; Max-Age=10800",
        practice_id
    );
    let headers = AppendHeaders([(header::SET_COOKIE, cookie)]);
    let mut context = Context::new();
    context.insert(
        "get_practice_url",
        &format!("/class7/practice/{}", practice_id),
    );
    let html = TEMPLATES.render("class7/index.html", &context);
    match html {
        Ok(t) => (headers, Html(t)),
        Err(e) => (headers, Html(format!("错误: {}", e))),
    }
}
