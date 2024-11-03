use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use sqlx::{Error, SqlitePool};

#[derive(sqlx::FromRow, Debug, Deserialize, Serialize)]
struct Quiz {
    id: u32,
    quiz_id: String,
    question_id: String,
    is_correct: i64,
    created_at: DateTime<Utc>,
}

#[derive(sqlx::FromRow, Debug, Deserialize, Serialize)]
struct Question {
    id: String,
    content: Option<String>,
    images: Option<String>,
    options: String,
}

#[derive(Debug, Deserialize, Serialize)]
struct QuestionOption {
    content: String,
    is_correct: bool,
}
