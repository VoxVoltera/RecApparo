<script>
  let { length = 6, value = $bindable('') } = $props();
  let els = [];

  function sync() { value = els.map((e) => e?.value || '').join(''); }

  function onInput(i, e) {
    e.target.value = e.target.value.replace(/\D/g, '').slice(0, 1);
    if (e.target.value && i < length - 1) els[i + 1]?.focus();
    sync();
  }
  function onKey(i, e) {
    if (e.key === 'Backspace' && !els[i]?.value && i > 0) els[i - 1]?.focus();
  }
  function onPaste(e) {
    e.preventDefault();
    const d = (e.clipboardData.getData('text') || '').replace(/\D/g, '').slice(0, length).split('');
    d.forEach((c, k) => { if (els[k]) els[k].value = c; });
    sync();
    els[Math.min(d.length, length - 1)]?.focus();
  }
</script>

<div class="otp">
  {#each Array(length) as _, i}
    <input class="otp-box" inputmode="numeric" maxlength="1" placeholder=""
           bind:this={els[i]} aria-label={'Digit ' + (i + 1)}
           oninput={(e) => onInput(i, e)} onkeydown={(e) => onKey(i, e)} onpaste={onPaste} />
  {/each}
</div>
