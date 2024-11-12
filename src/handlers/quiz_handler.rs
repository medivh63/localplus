use axum::{
    extract::{Path, State},
    http::{header, StatusCode},
    response::{AppendHeaders, Html, IntoResponse},
    Json,
};
use rand::seq::SliceRandom;
use tera::Context;
use tower_cookies::{Cookie, Cookies};
use uuid::Uuid;

use crate::{
    config::{AppState, TEMPLATES},
    models::{self, QuestionOption},
};

/// quiz 首页
pub async fn index(cookies: Cookies) -> impl IntoResponse {
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
    context.insert("quiz_id", &quiz_id);
    let html = TEMPLATES.render("class7/index.html", &context);
    match html {
        Ok(t) => (headers, Html(t)),
        Err(e) => {
            tracing::error!("render index.html error: {}", e);
            (headers, Html(format!("error: {}", e)))
        }
    }
}

/// quiz 问题
pub async fn quiz(State(state): State<AppState>, cookies: Cookies) -> impl IntoResponse {
    let mut context = Context::new();
    // 如果 cookies.get("quiz_id") 为空，则返回404
    let quiz_id = match cookies.get("quiz_id") {
        None => return Html(TEMPLATES.render("404.html", &context).unwrap()),
        Some(cookie) => cookie.value().to_string(),
    };
    // 已经回答过的题目id
    let answered_questions = models::get_answered_all_question(state.pool(), &quiz_id)
        .await
        .unwrap();
    // 已回答的题目id
    let answered_question_ids = answered_questions
        .iter()
        .map(|q| q.question_id.clone())
        .collect::<Vec<_>>();
    // 统计回答错误的题目数量
    let wrong_count = answered_questions.iter().filter(|q| !q.is_correct).count();
    // 总题目数量
    context.insert("total_count", &state.quiz_ids().len());
    // 已回答的题目数量
    context.insert("answered_count", &answered_questions.len());
    // 回答错误的题目数量
    context.insert("wrong_count", &wrong_count);
    // 取差集
    let quiz_ids = state
        .quiz_ids()
        .iter()
        .filter(|id| !answered_question_ids.contains(id))
        .collect::<Vec<_>>();
    if quiz_ids.is_empty() {
        // 渲染页面
        return Html(TEMPLATES.render("class7/completed.html", &context).unwrap());
    }
    // 随机取一个题目id
    let random_quiz_id = quiz_ids.choose(&mut rand::thread_rng()).unwrap();
    // 根据题目id查询题目
    let question = models::get_question_by_id(state.pool(), random_quiz_id)
        .await
        .unwrap();
    // 如果题目没有选项，则返回错误
    if question.options.is_empty() {
        return Html(format!("error: no options for question {}", quiz_id));
    }
    context.insert("question", &question);
    context.insert("quiz_id", &quiz_id);
    let options: Vec<QuestionOption> = serde_json::from_str(&question.options).unwrap();
    context.insert("options", &options);
    let html = TEMPLATES.render("class7/quiz.html", &context);
    match html {
        Ok(t) => Html(t),
        Err(e) => {
            tracing::error!("render quiz.html error: {}", e);
            Html(format!("error: {}", e))
        }
    }
}

/// quiz 提交答案
pub async fn answer(
    State(state): State<AppState>,
    Path(quiz_id): Path<String>,
    Json(form): Json<models::Answer>,
) -> impl IntoResponse {
    // 插入quiz
    let quiz = models::Quiz {
        id: None,
        quiz_id,
        question_id: form.question_id,
        is_correct: form.is_correct,
        created_at: chrono::Local::now().naive_utc(),
    };
    models::insert_quiz(state.pool(), &quiz).await.unwrap();
    (StatusCode::OK, "success")
}

/// quiz 重置测验
pub async fn reset_quiz(cookies: Cookies) -> impl IntoResponse {
    cookies.remove(Cookie::new("quiz_id", ""));
    (StatusCode::FOUND, [("Location", "/class7")])
}
