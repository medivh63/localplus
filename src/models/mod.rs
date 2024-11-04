pub use quiz::{Answer, Question, QuestionOption, Quiz};
use sqlx::{Error, SqlitePool};

mod quiz;

/// 获取所有题目id
pub async fn get_all_question_ids(pool: &SqlitePool) -> Result<Vec<String>, Error> {
    let recs = sqlx::query!("SELECT question_id FROM question")
        .fetch_all(pool)
        .await?;
    Ok(recs.iter().map(|r| r.question_id.clone()).collect())
}

/// 查询已回答全部题目
pub async fn get_answered_all_question(
    pool: &SqlitePool,
    quiz_id: &str,
) -> Result<Vec<Quiz>, Error> {
    let recs = sqlx::query_as!(
        Quiz,
        "SELECT id, quiz_id, question_id, is_correct as \"is_correct: bool\", created_at FROM quiz WHERE quiz_id = ?",
        quiz_id
    )
    .fetch_all(pool)
    .await?;
    Ok(recs)
}

/// 根据题目id查询题目
pub async fn get_question_by_id(pool: &SqlitePool, question_id: &str) -> Result<Question, Error> {
    let rec = sqlx::query_as!(
        Question,
        "SELECT * FROM question WHERE question_id = ?",
        question_id
    )
    .fetch_one(pool)
    .await?;
    Ok(rec)
}

/// 插入quiz
pub async fn insert_quiz(pool: &SqlitePool, quiz: &Quiz) -> Result<(), Error> {
    sqlx::query!(
        "INSERT INTO quiz (quiz_id, question_id, is_correct ,created_at) VALUES (?, ?, ?, ?)",
        quiz.quiz_id,
        quiz.question_id,
        quiz.is_correct,
        quiz.created_at
    )
    .execute(pool)
    .await?;
    Ok(())
}
