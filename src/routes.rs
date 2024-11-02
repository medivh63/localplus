use axum::Router;

use crate::AppState;

/// class 7 routes
fn driving_quiz_routes() -> Router<AppState> {
    Router::new()
        .route("/class7", get(index))
        .route("/:quiz_id", get(get_practice))
        .route("/:quiz_id/answers", post(answers))
        .route("/:quiz_id/restart", get(restart))
}
