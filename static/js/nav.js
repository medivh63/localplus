document.addEventListener('DOMContentLoaded', () => {
    // 获取汉堡菜单按钮和导航链接容器
    const navToggle = document.querySelector('.nav-toggle');
    const navLinks = document.querySelector('.nav-links');
    
    // 添加点击事件监听器
    navToggle.addEventListener('click', () => {
        navLinks.classList.toggle('show');
        
        // 可选：添加动画类
        navLinks.style.transition = 'all 0.3s ease-in-out';
    });
    
    // 点击导航链接后自动关闭菜单
    document.querySelectorAll('.nav-links .nav-link').forEach(link => {
        link.addEventListener('click', () => {
            navLinks.classList.remove('show');
        });
    });
    
    // 点击页面其他地方关闭菜单
    document.addEventListener('click', (event) => {
        if (!event.target.closest('.nav-toggle') && !event.target.closest('.nav-links')) {
            navLinks.classList.remove('show');
        }
    });
});

// 添加控制台日志以帮助调试
console.log('Nav script loaded'); 