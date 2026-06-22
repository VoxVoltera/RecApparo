<script>
  import Toggle from './Toggle.svelte';
  import Segmented from './Segmented.svelte';
  import Button from './Button.svelte';
  import { theme } from './theme.js';

  let { onlogout } = $props();

  const themeOptions = [
    { value: 'system', label: 'System' },
    { value: 'light', label: 'Light' },
    { value: 'dark', label: 'Dark' }
  ];
</script>

<div class="settings">
  <div class="set-group">
    <h4>Appearance</h4>
    <p class="g-sub">How RecApparō looks on this device.</p>
    <div class="set-row">
      <div class="r-text">Theme<small>System follows your OS setting</small></div>
      <Segmented options={themeOptions} value={$theme} onchange={(v) => theme.setMode(v)} />
    </div>
  </div>

  <div class="set-group">
    <h4>Notifications</h4>
    <p class="g-sub">Choose what reaches your inbox.</p>
    <div class="set-row"><div class="r-text">Assigned to me</div><Toggle checked /></div>
    <div class="set-row">
      <div class="r-text">Comment replies<small>When someone replies in a thread you're in</small></div>
      <Toggle checked />
    </div>
    <div class="set-row"><div class="r-text">Weekly digest</div><Toggle /></div>
  </div>

  <div class="set-group danger-zone">
    <h4>Account</h4>
    <p class="g-sub">Sign out of this session.</p>
    <div class="set-row">
      <div class="r-text">Log out<small>You'll need to sign in again</small></div>
      <Button variant="danger" icon="logout" onclick={() => onlogout?.()}>Log out</Button>
    </div>
  </div>
</div>
