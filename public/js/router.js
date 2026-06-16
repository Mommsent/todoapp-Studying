import { loadTasksPage } from '../pages/tasks.js';
import { loadUsersPage } from '../pages/users.js';
import { loadStatisticsPage } from '../pages/statistics.js';

const routes = {
    tasks: loadTasksPage,
    users: loadUsersPage,
    statistics: loadStatisticsPage,
};

/**
 * Активирует кнопку навигации.
 * @param {string} linkName - Название раздела (tasks, users, statistics).
 */
function setActiveNavButton(linkName) {
    document.querySelectorAll('.nav-btn').forEach(btn => {
        const isActive = btn.dataset.link === linkName;
        btn.classList.toggle('active', isActive);
    });
}

/**
 * Загружает компонент для указанного раздела.
 * @param {string} linkName - Название раздела.
 */
export async function loadComponent(linkName) {
    const mainContent = document.getElementById('main-content');
    if (!mainContent) {
        console.error('Element #main-content not found!');
        return;
    }

    const loadFunction = routes[linkName];
    if (!loadFunction) {
        mainContent.innerHTML = '<p class="error-message">Раздел не найден.</p>';
        return;
    }

    try {
        mainContent.innerHTML = ''; 
        await loadFunction(mainContent);
        setActiveNavButton(linkName);
    } catch (error) {
        console.error(`Error loading component ${linkName}:`, error);
        mainContent.innerHTML = `<p class="error-message">Ошибка загрузки раздела "${linkName}". ${error.message}</p>`;
    }
}