use axum::routing::{get, post};
use axum::Router;

use crate::handlers::quiz_handler::*;
use crate::AppState;

/// class 7 routes
pub(crate) fn driving_quiz_routes() -> Router<AppState> {
    Router::new()
        .route("/", get(index))
        .route("/quiz", get(quiz))
        .route("/:quiz_id/answer", post(answer))
        .route("/reset", get(reset_quiz))
}
