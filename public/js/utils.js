/**
 * Показывает уведомление в шапке приложения.
 * @param {string} message - Текст уведомления.
 * @param {'success'|'error'} type - Тип уведомления.
 * @param {number} duration - Время показа в мс.
 */
export function showNotification(message, type = 'success', duration = 3000) {
    const notificationEl = document.getElementById('notification');
    if (!notificationEl) return;

    notificationEl.textContent = message;
    notificationEl.className = `notification ${type}`;
    notificationEl.classList.remove('hidden');

    setTimeout(() => {
        notificationEl.classList.add('hidden');
    }, duration);
}

/**
 * Форматирует ISO строку даты в читаемый вид.
 * @param {string} isoString - Дата в формате ISO (например, "2026-02-26T10:30:00Z").
 * @returns {string} Отформатированная дата или "Никогда".
 */
export function formatDate(isoString) {
    if (!isoString) return 'Никогда';
    try {
        const date = new Date(isoString);
        return date.toLocaleDateString('ru-RU', {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: '2-digit',
            minute: '2-digit'
        });
    } catch (e) {
        return 'Некорректная дата';
    }
}

/**
 * Ставит первую букву строки в верхний регистр.
 * @param {string} str - Входная строка.
 * @returns {string} Строка с заглавной первой буквой.
 */
export function capitalize(str) {
    if (!str) return '';
    return str.charAt(0).toUpperCase() + str.slice(1);
}