<script>
  let { name, dot = 'var(--text-3)', count = null, addLabel = '+ Add ticket', showAdd = true, children } = $props();
</script>

<div class="lane">
  <div class="lane-head">
    <span class="l-name"><span class="dot" style="background:{dot}"></span>{name}</span>
    {#if count != null}<span class="count">{count}</span>{/if}
  </div>
  {@render children?.()}
  {#if showAdd}<div class="lane-add">{addLabel}</div>{/if}
</div>
