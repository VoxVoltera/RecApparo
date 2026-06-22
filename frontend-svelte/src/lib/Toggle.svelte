<script>
  let { checked = $bindable(false), label = '' } = $props();
</script>

<label class="switch">
  <input type="checkbox" bind:checked />
  <span class="track"></span>
  {#if label}<span>{label}</span>{/if}
</label>
