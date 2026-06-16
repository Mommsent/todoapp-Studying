const API_BASE_URL = 'http://localhost:5050';

/**
 * Обертка над fetch для обработки ошибок.
 * @param {string} endpoint - Путь API (например, '/tasks').
 * @param {object} options - Настройки для fetch (method, body, etc.).
 * @returns {Promise<any>} - Распарсенный JSON или текст ответа.
 */
async function request(endpoint, options = {}) {
    const url = `${API_BASE_URL}${endpoint}`;
    const config = {
        headers: {
            'Content-Type': 'application/json',
        },
        ...options,
    };

    if (config.body && typeof config.body === 'object') {
        config.body = JSON.stringify(config.body);
    }

    try {
        const response = await fetch(url, config);

        if (response.status === 204) {
            return null;
        }

        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.message || data.error || `HTTP error! status: ${response.status}`);
        }
        return data;
    } catch (error) {
        console.error(`API Error for ${endpoint}:`, error);
        throw error;
    }
}

export const api = {
    getUsers: (limit, offset) => {
        const params = new URLSearchParams();
        if (limit) params.append('limit', limit);
        if (offset) params.append('offset', offset);
        return request(`/users?${params.toString()}`);
    },
    getUser: (id) => request(`/users/${id}`),
    createUser: (userData) => request('/users', { method: 'POST', body: userData }),
    patchUser: (id, userData) => request(`/users/${id}`, { method: 'PATCH', body: userData }),
    deleteUser: (id) => request(`/users/${id}`, { method: 'DELETE' }),

    getTasks: (params = {}) => {
        const query = new URLSearchParams();
        if (params.user_id) query.append('user_id', params.user_id);
        if (params.limit) query.append('limit', params.limit);
        if (params.offset) query.append('offset', params.offset);
        return request(`/tasks?${query.toString()}`);
    },
    getTask: (id) => request(`/tasks/${id}`),
    createTask: (taskData) => request('/tasks', { method: 'POST', body: taskData }),
    patchTask: (id, taskData) => request(`/tasks/${id}`, { method: 'PATCH', body: taskData }),
    deleteTask: (id) => request(`/tasks/${id}`, { method: 'DELETE' }),

    getStatistics: (params = {}) => {
        const query = new URLSearchParams();
        if (params.user_id) query.append('user_id', params.user_id);
        if (params.from) query.append('from', params.from);
        if (params.to) query.append('to', params.to);
        return request(`/statistics?${query.toString()}`);
    },
};