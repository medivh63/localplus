document.addEventListener('DOMContentLoaded', function() {
    const form = document.querySelector('form');
    const submitBtn = document.querySelector('.submit-btn');

    form.addEventListener('submit', function(e) {
        const selectedOption = document.querySelector('input[name="answer"]:checked');
        if (!selectedOption) {
            e.preventDefault();
            alert('请选择一个答案后再提交。');
        } else {
            submitBtn.disabled = true;
            submitBtn.textContent = '提交中...';
        }
    });
});
