 // 等待 DOM 加载完成
document.addEventListener('DOMContentLoaded', function() {
    // 初始化主题
    initializeTheme();
});

// 初始化主题设置
function initializeTheme() {
    // 获取主题切换按钮
    const themeToggle = document.getElementById('theme-toggle');
    if (!themeToggle) return;

    // 获取保存的主题或使用默认主题
    const savedTheme = localStorage.getItem('theme') || 'light';
    
    // 设置初始主题
    document.documentElement.setAttribute('data-bs-theme', savedTheme);
    updateThemeIcon(savedTheme);

    // 添加切换事件
    themeToggle.addEventListener('click', toggleTheme);
}

// 切换主题
function toggleTheme() {
    const currentTheme = document.documentElement.getAttribute('data-bs-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    
    // 更新主题
    document.documentElement.setAttribute('data-bs-theme', newTheme);
    updateThemeIcon(newTheme);
    
    // 保存设置
    localStorage.setItem('theme', newTheme);
}

// 更新主题图标
function updateThemeIcon(theme) {
    const icon = document.querySelector('#theme-toggle i');
    if (!icon) return;

    // 移除现有图标类
    icon.classList.remove('fa-sun', 'fa-moon');
    
    // 添加新图标类
    icon.classList.add(theme === 'dark' ? 'fa-sun' : 'fa-moon');
}