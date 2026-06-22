<script>
  import Icon from './Icon.svelte';
  import { theme } from './theme.js';
</script>

<button class="icon-btn" aria-label="Toggle theme" title="Toggle theme" onclick={() => theme.toggle()}>
  <Icon name="moon" size={17} />
</button>
