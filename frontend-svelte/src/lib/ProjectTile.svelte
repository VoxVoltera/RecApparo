<script>
  import Icon from './Icon.svelte';
  let { name, icon = 'board', description = '', special = false,
        accent = 'var(--accent)', stats = [], assignees = [] } = $props();
</script>

<div class="ptile" class:special style="--ac:{accent}">
  <div class="spine"></div>
  <div class="p-icon"><Icon name={icon} size={20} /></div>
  <div>
    <h4>{name}</h4>
    {#if description}<p class="p-desc">{description}</p>{/if}
  </div>
  <div class="p-foot">
    <div class="p-stats">{#each stats as s}<span><b>{s.value}</b> {s.label}</span>{/each}</div>
    {#if assignees.length}
      <div class="av-stack">{#each assignees as a}<div class="av {a.tone}">{a.initials}</div>{/each}</div>
    {/if}
  </div>
</div>
