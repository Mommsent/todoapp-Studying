import { loadComponent } from './router.js';

document.addEventListener('DOMContentLoaded', () => {
    document.querySelector('nav').addEventListener('click', (event) => {
        const button = event.target.closest('button');
        if (!button) return;

        const linkName = button.dataset.link;
        if (linkName) {
            history.pushState(null, null, `#${linkName}`);
            loadComponent(linkName);
        }
    });

    const initialHash = window.location.hash.slice(1) || 'tasks';
    loadComponent(initialHash);
    const navBtn = document.querySelector(`.nav-btn[data-link="${initialHash}"]`);
    if (navBtn) navBtn.classList.add('active');
});

window.addEventListener('popstate', () => {
    const hash = window.location.hash.slice(1) || 'tasks';
    loadComponent(hash);
});

document.querySelector('nav').addEventListener('click', function(event) {
    console.group('🎯 Анализ клика');
    console.log('1. Реальная цель клика (event.target):', event.target);
    console.log('2. Ближайшая кнопка (closest):', event.target.closest('button'));
    
    const button = event.target.closest('button');
    if (button) {
        console.log('3. data-link кнопки:', button.dataset.link);
        console.log('4. Текст кнопки:', button.textContent.trim());
    } else {
        console.log('3. Клик мимо кнопок - игнорируем');
    }
    console.groupEnd();
});