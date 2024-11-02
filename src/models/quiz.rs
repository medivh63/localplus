struct Quiz {
    id: i32,
    title: String,
    description: String,
}

struct Question {
    id: i32,
    quiz_id: i32,
    content: String,
}

struct QuestionOption {
    id: i32,
    question_id: i32,
    content: String,
    is_correct: bool,
}
