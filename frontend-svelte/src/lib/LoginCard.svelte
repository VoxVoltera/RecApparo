<script>
  import Button from './Button.svelte';
  import Icon from './Icon.svelte';
  import Segmented from './Segmented.svelte';
  import OtpInput from './OtpInput.svelte';

  let mfa = $state('totp');   // 'totp' | 'key' -> only the selected method renders
  let code = $state('');
  let remember = $state(false);
</script>

<div class="login">
  <div class="l-head">
    <span class="glyph" style="width:40px;height:40px;border-radius:12px;font-size:19px">R</span>
    <h2>Welcome back</h2>
    <p class="sub">Log in to RecApparō</p>
  </div>

  <label class="field"><span>Email</span><input class="input" placeholder="you@team.com" /></label>
  <label class="field"><span>Password</span><input class="input" type="password" placeholder="••••••••" /></label>

  <div class="field">
    <span>Two-factor</span>
    <Segmented fill bind:value={mfa}
      options={[{ value: 'totp', label: 'Authenticator' }, { value: 'key', label: 'Security key' }]} />

    {#if mfa === 'totp'}
      <div class="mfa-pane">
        <OtpInput bind:value={code} />
        <span class="mfa-hint">Enter the 6-digit code from your authenticator app.</span>
      </div>
    {:else}
      <div class="mfa-pane">
        <div class="seckey">
          <div class="sk-key"><Icon name="securitykey" size={24} /></div>
          <div class="sk-text">
            <strong>Insert your YubiKey</strong>
            <span>Then tap the gold contact when it blinks.</span>
          </div>
        </div>
        <Button variant="secondary" size="sm" icon="key" block>Use security key</Button>
      </div>
    {/if}
  </div>

  <div class="row">
    <label class="check">
      <input type="checkbox" bind:checked={remember} />
      <span class="box"><Icon name="check" size={12} sw={3} /></span>Stay signed in
    </label>
    <a href="#forgot">Forgot password?</a>
  </div>

  <Button block>Verify &amp; log in</Button>
  <p style="text-align:center;font-size:11.5px;color:var(--text-3);margin:0">
    Lost your device? <a href="#backup">Use a backup code</a>
  </p>
</div>
