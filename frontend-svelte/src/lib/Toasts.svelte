<script>
  import { fly } from 'svelte/transition';
  import { toasts } from './toast.js';
  import Icon from './Icon.svelte';
</script>

<div class="toast-wrap">
  {#each $toasts as t (t.id)}
    <div class="toast show" in:fly={{ x: 60, duration: 300 }} out:fly={{ x: 120, duration: 300 }}>
      <div class="t-ic"><Icon name="check" size={15} sw={3} /></div>
      <div class="t-msg">{t.msg}{#if t.sub}<small>{t.sub}</small>{/if}</div>
    </div>
  {/each}
</div>
