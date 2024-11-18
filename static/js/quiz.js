function nextQuestion() {
    location.reload();
}

document.querySelectorAll('.option').forEach(option => {
    option.addEventListener('click', async function () {
        if (document.querySelector('.option.disabled')) {
            return;
        }

        const isCorrect = this.dataset.correct === 'true';
        const questionId = '{{ question.question_id }}';
        const quizId = '{{ quiz_id }}';

        try {
            const response = await fetch(`/class7/${quizId}/answer`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    question_id: questionId,
                    is_correct: isCorrect
                })
            });

            // 禁用所有选项
            document.querySelectorAll('.option').forEach(opt => {
                opt.classList.add('disabled');
                opt.style.pointerEvents = 'none';
                
                // 显示正确答案
                if (opt.dataset.correct === 'true') {
                    opt.classList.add('correct');
                    opt.querySelector('.option-icon').innerHTML = '<i class="fas fa-check"></i>';
                }
            });

            // 如果选择错误，标记当前选项
            if (!isCorrect) {
                this.classList.add('incorrect');
                this.querySelector('.option-icon').innerHTML = '<i class="fas fa-times"></i>';
            } else {
                // 只有在答对时自动进入下一题
                setTimeout(() => {
                    location.reload();
                }, 1500);
            }

        } catch (error) {
            console.error('提交答案失败:', error);
        }
    });
});