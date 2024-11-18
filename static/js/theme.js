// 等待 DOM 加载完成
document.addEventListener('DOMContentLoaded', () => {
    // 获取主题切换按钮和图标
    const themeToggle = document.querySelector('.theme-toggle');
    const themeIcon = themeToggle.querySelector('i');
    
    // 从 localStorage 获取保存的主题，默认为 light
    const savedTheme = localStorage.getItem('theme') || 'light';
    
    // 初始化主题
    setTheme(savedTheme);
    
    // 添加点击事件监听器
    themeToggle.addEventListener('click', () => {
        // 获取当前主题
        const currentTheme = document.documentElement.getAttribute('data-theme');
        // 切换主题
        const newTheme = currentTheme === 'light' ? 'dark' : 'light';
        setTheme(newTheme);
    });
    
    // 设置主题的函数
    function setTheme(theme) {
        // 设置 HTML 根元素的 data-theme 属性
        document.documentElement.setAttribute('data-theme', theme);
        // 保存主题到 localStorage
        localStorage.setItem('theme', theme);
        // 更新图标
        themeIcon.className = theme === 'light' ? 'fas fa-sun' : 'fas fa-moon';
    }
});

// 立即执行一次主题初始化
const savedTheme = localStorage.getItem('theme') || 'light';
document.documentElement.setAttribute('data-theme', savedTheme); 