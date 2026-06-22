<script>
  import Icon from './Icon.svelte';
  let {
    active = 'Boards',
    boards = [
      { name: 'Engineering', dot: 'var(--ac-teal)' },
      { name: 'Marketing Q3', dot: 'var(--ac-iris)' }
    ],
    user = { initials: 'VX', name: 'Voxy', email: 'voxy@recapparo.dev' }
  } = $props();
</script>

<div class="sidebar">
  <div class="sb-item" class:active={active === 'Boards'}><Icon name="boards" size={17} />Boards</div>
  <div class="sb-item" class:active={active === 'Projects'}><Icon name="folder" size={17} />Projects</div>
  <div class="sb-item" class:active={active === 'Settings'}><Icon name="settings" size={17} />Settings</div>
  <div class="sb-section">Boards</div>
  {#each boards as b}
    <div class="sb-item">{b.name}<span class="dot" style="background:{b.dot}"></span></div>
  {/each}
  <div class="sb-user">
    <div class="av a1">{user.initials}</div>
    <div><div class="nm">{user.name}</div><div class="em">{user.email}</div></div>
  </div>
</div>
