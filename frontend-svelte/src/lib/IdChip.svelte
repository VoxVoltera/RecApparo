<script>
  import { pushToast } from './toast.js';
  let { id } = $props();
  function copy() {
    if (navigator.clipboard) navigator.clipboard.writeText(id).catch(() => {});
    pushToast('Copied ' + id, 'Reference it in chat to flag a fix');
  }
</script>

<button class="id-chip" onclick={copy}>{id}</button>
