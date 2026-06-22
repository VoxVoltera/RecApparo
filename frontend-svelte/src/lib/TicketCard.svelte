<script>
  import Icon from './Icon.svelte';
  let {
    tkey, title, description = '', tags = [], assignees = [],
    comments = null, time = '', status = null,
    special = false, subboard = false, accent = 'var(--accent)'
  } = $props();
</script>

<div class="ticket" class:special style={special ? `--ac:${accent}` : ''}>
  <div class="t-top">
    <span class="tkey">{tkey}</span>
    {#if status}<span class="pill {status.variant}">{status.label}</span>{/if}
    {#if subboard}<span class="subboard"><Icon name="board" size={12} />Board</span>{/if}
  </div>

  <div class="t-title">{title}</div>
  {#if description}<div class="t-desc">{description}</div>{/if}

  {#if tags.length}
    <div class="demo-row" style="gap:6px">
      {#each tags as tg}<span class="tag {tg.variant}">{tg.label}</span>{/each}
    </div>
  {/if}

  {#if assignees.length || comments != null || time}
    <div class="t-foot">
      <div class="av-stack">
        {#each assignees as a}<div class="av {a.tone}">{a.initials}</div>{/each}
      </div>
      <div class="t-meta">
        {#if comments != null}<span><Icon name="comment" size={13} />{comments}</span>{/if}
        {#if time}<span><Icon name="clock" size={13} />{time}</span>{/if}
      </div>
    </div>
  {/if}
</div>
