<script>
  import Comment from './Comment.svelte';
  let { comments = [] } = $props();
</script>

<div class="thread">
  {#each comments as c}<Comment {...c} />{/each}
</div>
