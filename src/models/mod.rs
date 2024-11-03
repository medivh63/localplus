use sqlx::{Error, SqlitePool};

mod quiz;

/// 获取所有题目id
pub(crate) async fn get_all_question_ids(pool: &SqlitePool) -> Result<Vec<String>, Error> {
    let recs = sqlx::query!("SELECT question_id FROM question")
        .fetch_all(pool)
        .await?;
    Ok(recs.iter().filter_map(|r| r.question_id.clone()).collect())
}
