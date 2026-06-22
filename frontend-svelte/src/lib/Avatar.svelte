<script>
  let { initials = '', tone = 'a1', children } = $props();
</script>

<span class="av {tone}">{#if children}{@render children()}{:else}{initials}{/if}</span>
