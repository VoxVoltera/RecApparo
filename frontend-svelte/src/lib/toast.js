import { writable } from 'svelte/store';

export const toasts = writable([]);

let seq = 0;

export function pushToast(msg, sub = '', ttl = 2600) {
  const id = ++seq;
  toasts.update((list) => [...list, { id, msg, sub }]);
  setTimeout(() => {
    toasts.update((list) => list.filter((t) => t.id !== id));
  }, ttl);
}
