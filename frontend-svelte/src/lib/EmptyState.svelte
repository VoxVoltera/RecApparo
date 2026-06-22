<script>
  import Icon from './Icon.svelte';
  let { icon = 'inbox', title = '', description = '', children } = $props();
</script>

<div class="empty">
  <div class="e-icon"><Icon name={icon} size={23} /></div>
  <h4>{title}</h4>
  {#if description}<p>{description}</p>{/if}
  {@render children?.()}
</div>
