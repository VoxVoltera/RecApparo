<script>
  import Icon from './Icon.svelte';
  let { variant = 'primary', size = 'md', icon = null, block = false,
        type = 'button', children, ...rest } = $props();
</script>

<button {type} {...rest} class="btn btn-{variant}" class:btn-sm={size === 'sm'}
        style={block ? 'justify-content:center;width:100%' : ''}>
  {#if icon}<Icon name={icon} size={15} />{/if}{@render children?.()}
</button>
