<script>
  import Comment from './Comment.svelte';
  let { initials = '', tone = 'a1', name = '', ago = '', text = '', replies = [] } = $props();
</script>

<div class="comment">
  <div class="av {tone}">{initials}</div>
  <div class="body">
    <div class="c-head"><span class="nm">{name}</span><span class="ago">{ago}</span></div>
    <div class="c-text">{text}</div>
    <div class="c-actions"><button>Reply</button><button>React</button></div>
  </div>
</div>
{#each replies as r}
  <div class="reply"><Comment {...r} /></div>
{/each}
