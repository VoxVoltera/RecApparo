<script>
  let { open = $bindable(false), labelledby = 'modalTitle', onclose, children } = $props();
  function close() { open = false; onclose?.(); }
  function onKey(e) { if (e.key === 'Escape' && open) close(); }
</script>

<svelte:window onkeydown={onKey} />

<div class="modal-backdrop" class:open role="dialog" aria-modal="true"
     tabindex="-1" aria-labelledby={labelledby}
     onclick={(e) => { if (e.target === e.currentTarget) close(); }}>
  <div class="modal">
    {@render children?.(close)}
  </div>
</div>
