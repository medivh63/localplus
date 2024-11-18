// 常量定义
const ANIMATION_DURATION = 1500;
const CORRECT_ICON = '<i class="fas fa-check"></i>';
const INCORRECT_ICON = '<i class="fas fa-times"></i>';

// 下一题功能
function nextQuestion() {
    location.reload();
}

// 更新选项样式
function updateOptionStyles(option, isCorrect) {
    option.classList.add('disabled');
    option.style.pointerEvents = 'none';
    
    if (isCorrect) {
        option.classList.add('bg-success', 'text-white');
        option.querySelector('.option-icon').innerHTML = CORRECT_ICON;
    } else {
        option.classList.add('bg-danger', 'text-white');
        option.querySelector('.option-icon').innerHTML = INCORRECT_ICON;
    }
}

// 显示所有正确答案
function showCorrectAnswers() {
    document.querySelectorAll('.option[data-correct="true"]').forEach(option => {
        option.classList.add('bg-success', 'text-white');
        option.querySelector('.option-icon').innerHTML = CORRECT_ICON;
    });
}

// 禁用所有选项
function disableAllOptions() {
    document.querySelectorAll('.option').forEach(option => {
        option.classList.add('disabled');
        option.style.pointerEvents = 'none';
    });
}

// 提交答案到服务器
async function submitAnswer(quizId, questionId, isCorrect) {
    // 参数验证
    if (!quizId || !questionId) {
        throw new Error('Missing required parameters');
    }

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
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return response.json();
    } catch (error) {
        console.error('提交答案失败:', error);
        throw error;
    }
}

// 初始化选项点击事件
function initializeOptionHandlers() {
    document.querySelectorAll('.option').forEach(option => {
        option.addEventListener('click', async function() {
            // 如果已经有选项被选择，则返回
            if (document.querySelector('.option.disabled')) {
                return;
            }

            const isCorrect = this.dataset.correct === 'true';
            // 从 data 属性中获取 ID
            const questionId = document.querySelector('#quiz-container').dataset.questionId;
            const quizId = document.querySelector('#quiz-container').dataset.quizId;

            try {
                await submitAnswer(quizId, questionId, isCorrect);
                
                // 禁用所有选项
                disableAllOptions();
                
                // 显示正确答案
                showCorrectAnswers();
                
                // 更新当前选择的选项样式
                updateOptionStyles(this, isCorrect);

                // 如果答对，延迟后自动进入下一题
                if (isCorrect) {
                    setTimeout(nextQuestion, ANIMATION_DURATION);
                }

            } catch (error) {
                console.error('Error:', error);
                // 显示错误提示
                const toast = new bootstrap.Toast(document.getElementById('errorToast'));
                toast.show();
            }
        });
    });
}

// 页面加载完成后初始化
document.addEventListener('DOMContentLoaded', initializeOptionHandlers);