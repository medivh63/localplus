use crate::config::TEMPLATES;
use axum::response::Html;
use tera::Context;
pub(crate) mod quiz_handler;

/// 404页面
pub(crate) async fn fallback() -> Html<String> {
    let html = TEMPLATES.render("404.html", &Context::new());
    match html {
        Ok(t) => Html(t),
        Err(e) => Html(format!("Error: {}", e)),
    }
}
