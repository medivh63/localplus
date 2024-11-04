use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};

#[derive(sqlx::FromRow, Debug, Deserialize, Serialize)]
pub struct Quiz {
    pub id: Option<i64>,
    pub quiz_id: String,
    pub question_id: String,
    pub is_correct: bool,
    pub created_at: NaiveDateTime,
}

#[derive(sqlx::FromRow, Debug, Deserialize, Serialize)]
pub struct Question {
    pub id: Option<i64>,
    pub question_id: String,
    pub content: Option<String>,
    pub question_type: i64,
    pub images: Option<String>,
    pub options: String,
    pub created_at: NaiveDateTime,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct QuestionOption {
    pub content: String,
    pub is_correct: bool,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct Answer {
    pub question_id: String,
    pub is_correct: bool,
}
