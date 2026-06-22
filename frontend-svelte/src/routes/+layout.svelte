<script>
  import './layout.css';
  import './app.css';
  import favicon from '$lib/assets/favicon.svg';
  import { onMount } from 'svelte';
  import { theme } from '../lib/theme.js';
  import Toasts from '../lib/Toasts.svelte';

  let { children } = $props();

  onMount(() => theme.init());
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

{@render children()}
<Toasts />
