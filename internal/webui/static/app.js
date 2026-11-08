'use strict';
const $ = (selector) => document.querySelector(selector);
const state = { page: 'tasks', tasks: [], users: [], status: 'all', loaded: false, busy: false, editor: null, deleting: null, statsRequest: 0 };
const titles = { tasks: ['Задачи', 'Освободите голову. Сосредоточьтесь на важном.'], users: ['Пользователи', 'Люди, с которыми всё получается.'], statistics: ['Статистика', 'Посмотрите, как маленькие шаги складываются в результат.'] };
const escapeHTML = (value) => String(value ?? '').replace(/[&<>"']/g, (c) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' }[c]));
const initials = (name) => name.trim().split(/\s+/).slice(0, 2).map((s) => Array.from(s)[0]).join('').toUpperCase();
const dateLabel = (value) => value ? new Date(value).toLocaleDateString('ru-RU', { day: 'numeric', month: 'short' }) : '—';
const userName = (id) => state.users.find((u) => u.id === id)?.full_name || `Пользователь #${id}`;
let toastTimeout;
function toast(message) { $('#toast').textContent = message; $('#toast').hidden = false; clearTimeout(toastTimeout); toastTimeout = setTimeout(() => { $('#toast').hidden = true; }, 4000); }
function showError(selector, message) { const el = $(selector); el.textContent = message || ''; el.hidden = !message; }
async function api(path, options = {}) {
 let response;
 try { response = await fetch(`/api/v1${path}`, { ...options, headers: { 'Content-Type': 'application/json', ...options.headers }, cache: 'no-store', signal: AbortSignal.timeout(15000) }); }
 catch { throw new Error('Не удалось связаться с сервером. Проверьте подключение и повторите попытку.'); }
 if (response.status === 204) return null;
 let data;
 try { data = await response.json(); } catch { throw new Error('Сервер вернул некорректный ответ. Повторите попытку.'); }
 if (!response.ok) {
  const message = { 400: 'Проверьте заполненные поля.', 404: 'Запись уже удалена. Обновите данные.', 409: 'Запись была изменена. Обновите данные и попробуйте ещё раз.', 500: 'Ошибка сервера. Попробуйте позже.' }[response.status] || 'Не удалось выполнить запрос.';
  throw new Error(`${message}${data.message ? ` ${data.message}` : ''}`);
 }
 return data;
}
function options(selected = '', allLabel = 'Все авторы') { return `<option value="">${allLabel}</option>` + state.users.map((u) => `<option value="${u.id}" ${String(u.id) === String(selected) ? 'selected' : ''}>${escapeHTML(u.full_name)} · #${u.id}</option>`).join(''); }
function updateSelects() {
 for (const id of ['task-author', 'statistics-author']) { const el = $(`#${id}`); const value = el.value; el.innerHTML = options(value, id === 'statistics-author' ? 'Все пользователи' : 'Все авторы'); }
}
function empty(title, text, action = '') { return `<div class="empty"><span class="empty-symbol" aria-hidden="true">▤</span><h2>${title}</h2><p>${text}</p>${action}</div>`; }
function renderTasks() {
 const done = state.tasks.filter((t) => t.completed).length;
 $('#total-count').textContent = state.tasks.length; $('#active-count').textContent = state.tasks.length - done; $('#done-count').textContent = done; $('#task-count').textContent = state.tasks.length;
 const query = $('#task-search').value.toLocaleLowerCase('ru').trim(); const author = $('#task-author').value;
 const tasks = state.tasks.filter((t) => (state.status === 'all' || (state.status === 'done' ? t.completed : !t.completed)) && (!author || t.author_id === Number(author)) && `${t.title} ${t.description || ''}`.toLocaleLowerCase('ru').includes(query));
 $('#visible-count').textContent = `Показано: ${tasks.length}`; $('#task-footer-count').textContent = `${tasks.length} / ${state.tasks.length}`;
 $('#task-list').innerHTML = tasks.length ? tasks.map((t) => `<article class="task-row ${t.completed ? 'completed' : ''}"><input class="task-check" type="checkbox" data-toggle="${t.id}" ${t.completed ? 'checked' : ''} ${state.busy ? 'disabled' : ''} aria-label="${t.completed ? 'Вернуть в работу:' : 'Выполнить:'} ${escapeHTML(t.title)}"><div class="task-body"><button type="button" class="task-title" data-edit-task="${t.id}">${escapeHTML(t.title)}</button>${t.description ? `<p class="task-description">${escapeHTML(t.description)}</p>` : ''}<div class="task-meta"><span class="person"><span class="avatar" aria-hidden="true">${escapeHTML(initials(userName(t.author_id)))}</span>${escapeHTML(userName(t.author_id))}</span><span>Создана ${dateLabel(t.created_at)}</span>${t.completed_at ? `<span>Выполнена ${dateLabel(t.completed_at)}</span>` : ''}<span>#${t.id}</span></div></div><div class="row-actions"><span class="badge">${t.completed ? 'Выполнено' : 'В работе'}</span><button class="icon-button" type="button" data-edit-task="${t.id}" aria-label="Редактировать задачу ${escapeHTML(t.title)}">✎</button><button class="icon-button" type="button" data-delete-task="${t.id}" aria-label="Удалить задачу ${escapeHTML(t.title)}">×</button></div></article>`).join('') : empty(state.tasks.length ? 'Ничего не нашлось' : 'Место для ваших планов', state.tasks.length ? 'Попробуйте другой запрос или измените фильтры.' : 'Создайте первую задачу — и сделайте первый шаг к результату.', !state.tasks.length ? '<button class="primary" type="button" data-new-task>＋ Создать задачу</button>' : '');
}
function renderUsers() {
 const query = $('#user-search').value.toLocaleLowerCase('ru').trim();
 const users = state.users.filter((u) => `${u.full_name} ${u.phone_number || ''} ${u.id}`.toLocaleLowerCase('ru').includes(query));
 $('#user-list').innerHTML = users.length ? users.map((u) => `<article class="user-card"><span class="avatar" aria-hidden="true">${escapeHTML(initials(u.full_name))}</span><div class="user-info"><h2>${escapeHTML(u.full_name)}</h2><p>${escapeHTML(u.phone_number || 'Телефон не указан')}</p><p>ID ${u.id} · Задач: ${state.tasks.filter((t) => t.author_id === u.id).length}</p></div><div class="row-actions"><button class="icon-button" type="button" data-edit-user="${u.id}" aria-label="Редактировать ${escapeHTML(u.full_name)}">✎</button><button class="icon-button" type="button" data-delete-user="${u.id}" aria-label="Удалить ${escapeHTML(u.full_name)}">×</button></div></article>`).join('') : empty(query ? 'Никого не нашли' : 'Начните с команды', query ? 'Попробуйте другое имя, телефон или ID.' : 'Добавьте пользователя, чтобы назначать ему задачи.', !query ? '<button class="primary" type="button" data-new-user>＋ Добавить пользователя</button>' : '');
}
async function loadData() {
 if (state.busy) return;
 state.busy = true;
 document.querySelectorAll('.refresh-button').forEach((b) => { b.disabled = true; });
 if (!state.loaded) { $('#task-list').innerHTML = empty('Загружаем задачи…', 'Готовим ваше рабочее пространство.'); $('#user-list').innerHTML = empty('Загружаем пользователей…', ''); }
 showError('#load-error', '');
 try {
  const [tasks, users] = await Promise.all([api('/tasks'), api('/users')]);
  state.tasks = tasks || []; state.users = users || []; state.loaded = true; updateSelects(); renderTasks(); renderUsers();
 } catch (error) {
  showError('#load-error', error.message);
  if (!state.loaded) { $('#task-list').innerHTML = empty('Данные недоступны', 'Нажмите «Обновить», чтобы повторить загрузку.'); $('#user-list').innerHTML = empty('Данные недоступны', 'Нажмите «Обновить», чтобы повторить загрузку.'); }
 } finally { state.busy = false; if (state.loaded) renderTasks(); document.querySelectorAll('.refresh-button').forEach((b) => { b.disabled = false; }); }
 if (state.page === 'statistics') await loadStatistics();
}
async function loadStatistics() {
 const request = ++state.statsRequest;
 const form = $('#statistics-filter'); const values = new FormData(form);
 if (values.get('from') && values.get('to') && values.get('from') >= values.get('to')) { showError('#statistics-error', 'Конец периода должен быть позже начала. Дата «до» не входит в период.'); clearStatistics(); return; }
 const params = new URLSearchParams(); for (const [key, value] of values) if (value) params.set(key, value);
 showError('#statistics-error', ''); clearStatistics();
 try {
  const data = await api(`/statistics?${params}`); if (request !== state.statsRequest) return;
  $('#stat-created').textContent = data.created_task; $('#stat-completed').textContent = data.completed_task;
  $('#stat-rate').textContent = data.completed_task_rate == null ? '—' : `${Number(data.completed_task_rate).toLocaleString('ru-RU', { maximumFractionDigits: 1 })}%`;
  $('#stat-duration').textContent = formatDuration(data.arg_time);
  const rate = Math.min(100, Math.max(0, data.completed_task_rate || 0)); $('#progress-bar').style.width = `${rate}%`; $('#progress').setAttribute('aria-valuenow', rate);
  $('#progress-caption').textContent = data.created_task ? `Выполнено ${data.completed_task} из ${data.created_task} задач за выбранный период.` : 'За выбранный период задач пока нет.';
 } catch (error) { if (request === state.statsRequest) showError('#statistics-error', error.message); }
}
function clearStatistics() { for (const id of ['created', 'completed', 'rate', 'duration']) $(`#stat-${id}`).textContent = '—'; $('#progress-bar').style.width = '0%'; $('#progress').removeAttribute('aria-valuenow'); $('#progress-caption').textContent = 'Нет данных для отображения.'; }
function formatDuration(value) {
 if (!value) return '—';
 const parts = [...value.matchAll(/([\d.]+)(h|ms|µs|ns|m|s)/g)];
 const seconds = parts.reduce((sum, [, n, unit]) => sum + Number(n) * ({ h: 3600, m: 60, s: 1, ms: .001, 'µs': .000001, ns: .000000001 }[unit]), 0);
 if (seconds < 60) return '< 1 мин';
 const minutes = Math.floor(seconds / 60); const hours = Math.floor(minutes / 60); const days = Math.floor(hours / 24);
 return days ? `${days} д ${hours % 24} ч` : hours ? `${hours} ч ${minutes % 60} мин` : `${minutes} мин`;
}
function navigate() {
 const requested = location.hash.slice(1); state.page = Object.hasOwn(titles, requested) ? requested : 'tasks';
 const [title, description] = titles[state.page]; $('#page-title').textContent = title; $('#breadcrumb').textContent = title; $('#page-description').textContent = description; document.title = `${title} — Порядок`;
 for (const page of Object.keys(titles)) { $(`#${page}-page`).hidden = state.page !== page; const nav = $(`[data-page="${page}"]`); nav.classList.toggle('active', state.page === page); if (state.page === page) nav.setAttribute('aria-current', 'page'); else nav.removeAttribute('aria-current'); }
 $('#create-button').hidden = state.page === 'statistics'; $('#create-button').textContent = state.page === 'users' ? '＋ Добавить пользователя' : '＋ Новая задача';
 if (state.page === 'statistics') loadStatistics();
}
async function openEditor(kind, id = null) {
 if (state.busy) return;
 showError('#editor-error', ''); state.editor = { kind, id }; let record = {};
 if (kind === 'tasks' && !id && !state.users.length) { toast('Сначала добавьте пользователя — автора задачи.'); location.hash = 'users'; return; }
 if (id) {
  state.busy = true;
  try { record = await api(`/${kind}/${id}`); } catch (error) { showError('#load-error', error.message); return; } finally { state.busy = false; }
 }
 $('#editor-title').textContent = kind === 'tasks' ? (id ? 'Редактировать задачу' : 'Новая задача') : (id ? 'Редактировать пользователя' : 'Новый пользователь');
 $('#editor-fields').innerHTML = kind === 'users' ? `<label>Имя и фамилия<input name="full_name" required minlength="3" maxlength="100" autocomplete="name" placeholder="Например, Анна Иванова" value="${escapeHTML(record.full_name)}"></label><label>Телефон <span class="muted">· необязательно</span><input name="phone_number" type="tel" autocomplete="tel" pattern="\\+[0-9]{9,14}" maxlength="15" placeholder="+79991234567" value="${escapeHTML(record.phone_number)}"><small>От 10 до 15 символов, начиная с +</small></label>` : `<label>Название<input name="title" required maxlength="100" placeholder="Что нужно сделать?" value="${escapeHTML(record.title)}"></label><label>Описание <span class="muted">· необязательно</span><textarea name="description" maxlength="1000" placeholder="Детали, которые помогут не забыть главное">${escapeHTML(record.description)}</textarea></label>${id ? `<p class="muted">Автор: ${escapeHTML(userName(record.author_id))} · Создана ${dateLabel(record.created_at)}${record.completed_at ? ` · Выполнена ${dateLabel(record.completed_at)}` : ''}</p><label>Статус<select name="completed"><option value="false" ${!record.completed ? 'selected' : ''}>В работе</option><option value="true" ${record.completed ? 'selected' : ''}>Выполнено</option></select></label>` : `<label>Автор<select name="author_id" required>${options('', 'Выберите пользователя')}</select></label>`}`;
 $('#editor').showModal(); $('#editor-fields input').focus();
}
async function saveEditor(event) {
 event.preventDefault(); if (state.busy) return;
 const { kind, id } = state.editor; const values = new FormData(event.target); let body;
 if (kind === 'users') body = { full_name: values.get('full_name').trim(), phone_number: values.get('phone_number').trim() || null };
 else { body = { title: values.get('title').trim(), description: values.get('description').trim() || null }; if (id) body.completed = values.get('completed') === 'true'; else body.author_id = Number(values.get('author_id')); }
 if (kind === 'users' ? Array.from(body.full_name).length < 3 : !body.title) { showError('#editor-error', kind === 'users' ? 'Имя должно содержать не менее 3 символов.' : 'Введите название задачи.'); return; }
 state.busy = true; $('#save-button').disabled = true; showError('#editor-error', '');
 try { await api(`/${kind}${id ? `/${id}` : ''}`, { method: id ? 'PATCH' : 'POST', body: JSON.stringify(body) }); $('#editor').close(); toast(id ? 'Изменения сохранены' : kind === 'tasks' ? 'Задача создана' : 'Пользователь добавлен'); }
 catch (error) { showError('#editor-error', error.message); }
 finally { state.busy = false; $('#save-button').disabled = false; }
 if (!$('#editor').open) await loadData();
}
function confirmDelete(kind, id) {
 if (state.busy) return;
 const record = state[kind].find((item) => item.id === id); if (!record) return;
 state.deleting = { kind, id }; $('#delete-title').textContent = kind === 'users' ? 'Удалить пользователя?' : 'Удалить задачу?';
 $('#delete-description').textContent = `«${kind === 'users' ? record.full_name : record.title}» будет удалён${kind === 'users' ? ' вместе со всеми его задачами' : 'а'}. Это действие нельзя отменить.`;
 showError('#delete-error', ''); $('#delete-dialog').showModal();
}
async function deleteRecord(event) {
 event.preventDefault(); if (state.busy) return;
 const { kind, id } = state.deleting; const button = event.target.querySelector('[type=submit]'); state.busy = true; button.disabled = true;
 try { await api(`/${kind}/${id}`, { method: 'DELETE' }); $('#delete-dialog').close(); toast('Запись удалена'); } catch (error) { showError('#delete-error', error.message); } finally { state.busy = false; button.disabled = false; }
 if (!$('#delete-dialog').open) await loadData();
}
async function toggleTask(input) {
 if (state.busy) { input.checked = !input.checked; return; }
 state.busy = true; input.disabled = true;
 try { const task = await api(`/tasks/${input.dataset.toggle}`, { method: 'PATCH', body: JSON.stringify({ completed: input.checked }) }); state.tasks = state.tasks.map((t) => t.id === task.id ? task : t); toast(task.completed ? 'Ещё одна задача выполнена!' : 'Задача возвращена в работу'); }
 catch (error) { input.checked = !input.checked; showError('#load-error', error.message); }
 finally { state.busy = false; renderTasks(); renderUsers(); }
}
document.addEventListener('click', (event) => {
 const target = event.target.closest('button'); if (!target) return;
 if (target.dataset.close) { if (!state.busy) $(`#${target.dataset.close}`).close(); return; }
 if (target.matches('[data-status]')) { state.status = target.dataset.status; document.querySelectorAll('[data-status]').forEach((b) => { b.classList.toggle('selected', b === target); b.setAttribute('aria-pressed', b === target); }); renderTasks(); }
 if (target.id === 'create-button') openEditor(state.page);
 if (target.hasAttribute('data-new-task')) openEditor('tasks');
 if (target.hasAttribute('data-new-user')) openEditor('users');
 if (target.classList.contains('refresh-button')) loadData();
 for (const [attr, kind] of [['editTask', 'tasks'], ['editUser', 'users']]) if (target.dataset[attr]) openEditor(kind, Number(target.dataset[attr]));
 for (const [attr, kind] of [['deleteTask', 'tasks'], ['deleteUser', 'users']]) if (target.dataset[attr]) confirmDelete(kind, Number(target.dataset[attr]));
});
document.addEventListener('change', (event) => { if (event.target.matches('[data-toggle]')) toggleTask(event.target); });
$('#task-search').addEventListener('input', renderTasks); $('#task-author').addEventListener('change', renderTasks); $('#user-search').addEventListener('input', renderUsers);
$('#editor-form').addEventListener('submit', saveEditor); $('#delete-form').addEventListener('submit', deleteRecord);
$('#statistics-filter').addEventListener('submit', (event) => { event.preventDefault(); loadStatistics(); }); $('#statistics-filter').addEventListener('reset', () => { setTimeout(loadStatistics, 0); });
for (const id of ['editor', 'delete-dialog']) $(`#${id}`).addEventListener('cancel', (event) => { if (state.busy) event.preventDefault(); });
$('#today').textContent = new Date().toLocaleDateString('ru-RU', { day: 'numeric', month: 'long', weekday: 'long' });
window.addEventListener('hashchange', navigate); navigate(); loadData();
