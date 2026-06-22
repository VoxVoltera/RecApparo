<script>
  let { accent = 'var(--ac-teal)', title = '', crumb = '', from = '', description = '', stats = [] } = $props();
</script>

<div class="dboard" style="--ac:{accent}">
  {#if crumb || from}
    <div class="crumb">{crumb} <span style="opacity:.5">▸</span> <span class="from">from {from}</span></div>
  {/if}
  <h2>{title}</h2>
  {#if description}<p>{description}</p>{/if}
  {#if stats.length}
    <div class="d-stats">
      {#each stats as s}<span><b>{s.value}</b>{s.label}</span>{/each}
    </div>
  {/if}
</div>
