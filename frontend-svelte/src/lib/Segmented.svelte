<script>
  import Icon from './Icon.svelte';
  let { options = [], value = $bindable(), fill = false, onchange } = $props();
  function select(v) { value = v; onchange?.(v); }
</script>

<div class="seg" role="tablist" style={fill ? 'width:100%' : ''}>
  {#each options as o}
    <button role="tab" aria-selected={value === o.value}
            style={fill ? 'flex:1;justify-content:center' : ''}
            onclick={() => select(o.value)}>
      {#if o.icon}<Icon name={o.icon} size={14} />{/if}{o.label}
    </button>
  {/each}
</div>
