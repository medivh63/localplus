use axum::routing::get;
use axum::Router;

use crate::handlers::quiz_handler::*;
use crate::AppState;

/// class 7 routes
pub(crate) fn driving_quiz_routes() -> Router<AppState> {
    Router::new().route("/class7", get(index))
}
