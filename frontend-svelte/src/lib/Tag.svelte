<script>
  let { variant = 'scoped', children } = $props();
</script>

<span class="tag {variant}">{@render children?.()}</span>
