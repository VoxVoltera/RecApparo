import { writable } from 'svelte/store';

const KEY = 'recapparo-theme';

function read() {
  try { return localStorage.getItem(KEY) || 'dark'; } catch { return 'dark'; }
}

function resolve(mode) {
  if (mode === 'system' && typeof window !== 'undefined') {
    return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
  }
  return mode === 'system' ? 'dark' : mode;
}

function apply(mode) {
  if (typeof document !== 'undefined') {
    document.documentElement.setAttribute('data-theme', resolve(mode));
  }
}

function createTheme() {
  const { subscribe, set } = writable(read());

  return {
    subscribe,
    /** mode: 'system' | 'light' | 'dark' */
    setMode(mode) {
      set(mode);
      apply(mode);
      try { localStorage.setItem(KEY, mode); } catch { /* desktop webview may block */ }
    },
    /** flip between light and dark based on what's currently applied */
    toggle() {
      const applied = document.documentElement.getAttribute('data-theme');
      this.setMode(applied === 'dark' ? 'light' : 'dark');
    },
    /** call once on app start */
    init() { apply(read()); }
  };
}

export const theme = createTheme();
