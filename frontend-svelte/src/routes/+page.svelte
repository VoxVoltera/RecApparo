<script>
  import { theme } from '../lib/theme.js';
  import { pushToast } from '../lib/toast.js';

  import IdChip from '../lib/IdChip.svelte';
  import Icon from '../lib/Icon.svelte';
  import ThemeToggle from '../lib/ThemeToggle.svelte';
  import Button from '../lib/Button.svelte';
  import Tag from '../lib/Tag.svelte';
  import Avatar from '../lib/Avatar.svelte';
  import Toggle from '../lib/Toggle.svelte';
  import Segmented from '../lib/Segmented.svelte';
  import TicketCard from '../lib/TicketCard.svelte';
  import Lane from '../lib/Lane.svelte';
  import EmptyState from '../lib/EmptyState.svelte';
  import DaughterBoardHeader from '../lib/DaughterBoardHeader.svelte';
  import ProjectTile from '../lib/ProjectTile.svelte';
  import Sidebar from '../lib/Sidebar.svelte';
  import LoginCard from '../lib/LoginCard.svelte';
  import SettingsPanel from '../lib/SettingsPanel.svelte';
  import Modal from '../lib/Modal.svelte';
  import CommentThread from '../lib/CommentThread.svelte';

  const themeIconOptions = [
    { value: 'system', label: 'System' },
    { value: 'light', label: 'Light', icon: 'sun' },
    { value: 'dark', label: 'Dark', icon: 'moon' }
  ];

  // C20 accent carryover
  let dboardAccent = $state('var(--ac-teal)');
  let dboardFrom = $state('RAP-131');
  const picker = [
    { ac: 'var(--ac-iris)', from: 'RAP-128' },
    { ac: 'var(--ac-teal)', from: 'RAP-131' },
    { ac: 'var(--ac-amber)', from: 'RAP-150' },
    { ac: 'var(--ac-rose)', from: 'RAP-140' },
    { ac: 'var(--ac-violet)', from: 'RAP-162' }
  ];
  function pick(p) { dboardAccent = p.ac; dboardFrom = p.from; }

  // C26 logout modal
  let modalOpen = $state(false);

  // C28 thread data
  const thread = [
    {
      initials: 'MR', tone: 'a2', name: 'Mara R.', ago: '2h ago',
      text: 'Should the special-ticket glow also show on the project tile, or just the card?',
      replies: [
        { initials: 'VX', tone: 'a1', name: 'Voxy', ago: '1h ago',
          text: 'Both — the tile gets a softer version so the folder still reads calm. See C21.' }
      ]
    }
  ];
</script>

<header class="topbar">
  <div class="wordmark"><span class="glyph">R</span> RecApparō <small>v0.1 · gallery</small></div>
  <nav class="topnav">
    <a href="#controls">Controls</a>
    <a href="#tickets">Tickets</a>
    <a href="#views">Views</a>
    <a href="#overlays">Overlays</a>
  </nav>
  <div class="spacer"></div>
  <ThemeToggle />
</header>

<main class="wrap">
  <div class="lede">
    <h1>Component <em>gallery</em></h1>
    <p>Every building block for RecApparō on one page. Each tile carries a mono ID — click it to copy. Tell me which ID is off and I'll fix just that piece.</p>
  </div>
  <div class="legend">
    <span>🔖 <strong>Click any <kbd>ID</kbd> chip to copy it</strong></span>
    <span>🌗 Theme toggle, top-right — both modes are live</span>
    <span>✨ Special-ticket accent blooms into the daughter board in <kbd>C20</kbd></span>
  </div>

  <!-- FOUNDATIONS -->
  <section class="section" id="foundations">
    <div class="eyebrow">Foundations</div>
    <div class="grid">
      <div class="tile span2">
        <div class="tile-head"><div class="meta"><h3>Surface &amp; accent tokens</h3><p>Theme-driven surfaces, plus the accent family special tickets draw from.</p></div><IdChip id="C01" /></div>
        <div class="demo">
          <span class="label">SURFACES</span>
          <div class="swatch-grid">
            <div class="swatch"><div class="chip" style="background:var(--bg)"></div><span class="nm">bg</span></div>
            <div class="swatch"><div class="chip" style="background:var(--surface)"></div><span class="nm">surface</span></div>
            <div class="swatch"><div class="chip" style="background:var(--surface-2)"></div><span class="nm">surface-2</span></div>
            <div class="swatch"><div class="chip" style="background:var(--surface-3)"></div><span class="nm">surface-3</span></div>
            <div class="swatch"><div class="chip" style="background:var(--text-2)"></div><span class="nm">text-2</span></div>
          </div>
          <span class="label">ACCENT FAMILY</span>
          <div class="swatch-grid">
            <div class="swatch"><div class="chip" style="background:var(--ac-iris)"></div><span class="nm">iris</span><code>brand</code></div>
            <div class="swatch"><div class="chip" style="background:var(--ac-teal)"></div><span class="nm">teal</span></div>
            <div class="swatch"><div class="chip" style="background:var(--ac-amber)"></div><span class="nm">amber</span></div>
            <div class="swatch"><div class="chip" style="background:var(--ac-rose)"></div><span class="nm">rose</span></div>
            <div class="swatch"><div class="chip" style="background:var(--ac-lime)"></div><span class="nm">lime</span></div>
            <div class="swatch"><div class="chip" style="background:var(--ac-violet)"></div><span class="nm">violet</span></div>
          </div>
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Type scale</h3><p>Space Grotesk display · Plus Jakarta body · Space Mono data.</p></div><IdChip id="C02" /></div>
        <div class="demo">
          <div class="type-row"><span class="tag-sz">34 / display</span><span style="font-family:var(--f-display);font-size:26px;font-weight:700;letter-spacing:-.02em">Sprint board</span></div>
          <div class="type-row"><span class="tag-sz">18 / heading</span><span style="font-family:var(--f-display);font-size:18px;font-weight:600">Backlog grooming</span></div>
          <div class="type-row"><span class="tag-sz">14 / body</span><span style="font-size:14px">Move a ticket to start work.</span></div>
          <div class="type-row"><span class="tag-sz">11 / mono</span><span style="font-family:var(--f-mono);font-size:12px;color:var(--text-3)">RAP-1042</span></div>
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Radii &amp; elevation</h3><p>Generous corners; soft, low shadows.</p></div><IdChip id="C03" /></div>
        <div class="demo">
          <div class="demo-row">
            <div style="width:54px;height:54px;background:var(--surface-2);border:1px solid var(--border);border-radius:var(--r-sm)"></div>
            <div style="width:54px;height:54px;background:var(--surface-2);border:1px solid var(--border);border-radius:var(--r-md)"></div>
            <div style="width:54px;height:54px;background:var(--surface-2);border:1px solid var(--border);border-radius:var(--r-lg)"></div>
            <div style="width:54px;height:54px;background:var(--surface-2);border:1px solid var(--border);border-radius:var(--r-xl)"></div>
          </div>
          <div class="demo-row">
            <div style="width:90px;height:48px;background:var(--surface);border-radius:var(--r-md);box-shadow:var(--shadow-sm)"></div>
            <div style="width:90px;height:48px;background:var(--surface);border-radius:var(--r-md);box-shadow:var(--shadow)"></div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- CONTROLS -->
  <section class="section" id="controls">
    <div class="eyebrow">Controls</div>
    <div class="grid">
      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Buttons</h3><p>Primary, secondary, ghost, danger.</p></div><IdChip id="C04" /></div>
        <div class="demo">
          <div class="demo-row"><Button>New ticket</Button><Button variant="secondary">Cancel</Button></div>
          <div class="demo-row"><Button variant="ghost">Skip</Button><Button variant="danger">Delete board</Button><Button size="sm" icon="plus">Add</Button></div>
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Icon buttons</h3><p>For toolbars and card affordances.</p></div><IdChip id="C05" /></div>
        <div class="demo"><div class="demo-row">
          <button class="icon-btn" aria-label="Search"><Icon name="search" size={17} /></button>
          <button class="icon-btn" aria-label="Add"><Icon name="plus" size={17} /></button>
          <button class="icon-btn" aria-label="Notifications"><Icon name="bell" size={17} /></button>
          <button class="icon-btn" aria-label="Settings"><Icon name="settings" size={17} /></button>
        </div></div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Text inputs</h3><p>Default, focused, and error states.</p></div><IdChip id="C06" /></div>
        <div class="demo">
          <label class="field"><span>Board name</span><input class="input" placeholder="e.g. Q3 Marketing" /></label>
          <label class="field"><span>Email</span><input class="input error" value="not-an-email" /><span class="hint"><Icon name="alert" size={13} />Enter a valid email address.</span></label>
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Search field</h3><p>Filter tickets and boards.</p></div><IdChip id="C07" /></div>
        <div class="demo"><div class="search"><Icon name="search" size={16} /><input class="input" placeholder="Search tickets…" /></div></div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Toggle switch</h3><p>On/off settings.</p></div><IdChip id="C08" /></div>
        <div class="demo">
          <Toggle checked label="Email notifications" />
          <Toggle label="Compact card density" />
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Checkbox &amp; radio</h3><p>Multi- and single-select.</p></div><IdChip id="C09" /></div>
        <div class="demo">
          <label class="check"><input type="checkbox" checked /><span class="box"><Icon name="check" size={12} sw={3} /></span>Notify assignees</label>
          <label class="check"><input type="checkbox" /><span class="box"><Icon name="check" size={12} sw={3} /></span>Archive when done</label>
          <label class="check"><input type="radio" name="r" checked /><span class="box round"></span>All activity</label>
          <label class="check"><input type="radio" name="r" /><span class="box round"></span>Mentions only</label>
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Select</h3><p>Choose a column or assignee.</p></div><IdChip id="C10" /></div>
        <div class="demo"><div class="select"><select><option>In progress</option><option>To do</option><option>In review</option><option>Done</option></select><Icon name="chevron" size={15} /></div></div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Segmented control</h3><p>Drives the appearance setting.</p></div><IdChip id="C11" /></div>
        <div class="demo"><Segmented options={themeIconOptions} value={$theme} onchange={(v) => theme.setMode(v)} /></div>
      </div>
    </div>
  </section>

  <!-- DATA & PEOPLE -->
  <section class="section" id="data">
    <div class="eyebrow">Data &amp; people</div>
    <div class="grid">
      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Tags &amp; labels</h3><p>Board-scoped tags carry accent; global tags stay neutral.</p></div><IdChip id="C12" /></div>
        <div class="demo">
          <span class="label">BOARD-SCOPED</span>
          <div class="demo-row"><Tag variant="scoped">design</Tag><Tag variant="t-teal">backend</Tag><Tag variant="t-amber">blocked</Tag><Tag variant="t-rose">urgent</Tag></div>
          <span class="label">GLOBAL</span>
          <div class="demo-row"><Tag variant="global">bug</Tag><Tag variant="global">chore</Tag><Tag variant="global">docs</Tag></div>
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Avatars &amp; assignees</h3><p>Single avatar and overflow stack.</p></div><IdChip id="C13" /></div>
        <div class="demo">
          <div class="demo-row"><Avatar initials="VX" tone="a1" /><Avatar initials="MR" tone="a2" /><Avatar initials="JD" tone="a3" /></div>
          <div class="demo-row"><div class="av-stack"><div class="av a1">VX</div><div class="av a2">MR</div><div class="av a3">JD</div><div class="av a4">KP</div><div class="av-more">+3</div></div></div>
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Ticket key &amp; status</h3><p>Mono key plus column status pill.</p></div><IdChip id="C14" /></div>
        <div class="demo">
          <div class="demo-row"><span class="tkey">RAP-128</span><span class="tkey">RAP-1042</span></div>
          <div class="demo-row"><span class="pill todo">To do</span><span class="pill doing">In progress</span><span class="pill done">Done</span></div>
        </div>
      </div>
    </div>
  </section>

  <!-- TICKETS -->
  <section class="section" id="tickets">
    <div class="eyebrow">Tickets</div>
    <div class="grid">
      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Ticket card — standard</h3><p>Title, tags, assignees, comment count.</p></div><IdChip id="C16" /></div>
        <div class="demo">
          <TicketCard tkey="RAP-204" title="Wire column drag-and-drop to the ordering endpoint"
            status={{ label: 'In progress', variant: 'doing' }}
            tags={[{ label: 'frontend', variant: 'scoped' }, { label: 'chore', variant: 'global' }]}
            assignees={[{ initials: 'MR', tone: 'a2' }, { initials: 'VX', tone: 'a1' }]}
            comments={3} />
        </div>
      </div>

      <div class="tile span2">
        <div class="tile-head"><div class="meta"><h3>Ticket card — special</h3><p>Accent border + background bloom. Used for tickets that own a sub-board. Shown across the accent family.</p></div><IdChip id="C17" /></div>
        <div class="demo">
          <div class="grid" style="grid-template-columns:repeat(auto-fill,minmax(300px,1fr));gap:16px">
            <TicketCard special subboard accent="var(--ac-iris)" tkey="RAP-128" title="Marketing campaign — Q3"
              description="Plan, brief and schedule the third-quarter push across paid and owned channels."
              tags={[{ label: 'campaign', variant: 'scoped' }, { label: 'planning', variant: 'global' }]}
              assignees={[{ initials: 'JD', tone: 'a3' }, { initials: 'VX', tone: 'a1' }]} comments={12} time="2d" />
            <TicketCard special subboard accent="var(--ac-teal)" tkey="RAP-131" title="Backend migration"
              description="Move the legacy column store onto the new schema without dropping ticket history."
              tags={[{ label: 'migration', variant: 't-teal' }, { label: 'backend', variant: 'global' }]}
              assignees={[{ initials: 'MR', tone: 'a2' }]} comments={5} time="6h" />
            <TicketCard special subboard accent="var(--ac-rose)" tkey="RAP-140" title="Launch checklist"
              description="Final sign-off items every release needs before it ships to production."
              tags={[{ label: 'launch', variant: 't-rose' }, { label: 'urgent', variant: 'global' }]}
              assignees={[{ initials: 'VX', tone: 'a1' }, { initials: 'KP', tone: 'a4' }]} comments={9} time="1d" />
          </div>
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Column / lane</h3><p>First-class column with name, count, add row.</p></div><IdChip id="C18" /></div>
        <div class="demo">
          <Lane name="In progress" dot="var(--ac-amber)" count={2}>
            <TicketCard tkey="RAP-204" title="Drag-and-drop ordering" />
            <TicketCard tkey="RAP-209" title="Theme toggle persistence" />
          </Lane>
        </div>
      </div>
    </div>
  </section>

  <!-- SURFACES & VIEWS -->
  <section class="section" id="views">
    <div class="eyebrow">Surfaces &amp; views</div>
    <div class="grid">
      <div class="tile span2">
        <div class="tile-head"><div class="meta"><h3>Empty state</h3><p>What a column, board, or search shows before there's anything in it — an invitation to act, not a dead end.</p></div><IdChip id="C19" /></div>
        <div class="demo">
          <div class="grid" style="grid-template-columns:repeat(auto-fill,minmax(244px,1fr));gap:16px">
            <EmptyState icon="inbox" title="No tickets yet" description="Add the first ticket to this column, or drag one over from another lane.">
              <Button size="sm" icon="plus">New ticket</Button>
            </EmptyState>
            <EmptyState icon="search" title="No matches" description="Nothing matches “schema migration”. Try a different term or clear the filter.">
              <Button variant="secondary" size="sm">Clear search</Button>
            </EmptyState>
          </div>
        </div>
      </div>

      <div class="tile span2">
        <div class="tile-head"><div class="meta"><h3>Daughter board header — inherited glow</h3><p>Open a special ticket's sub-board and its accent carries onto the page. Pick an accent to watch it propagate.</p></div><IdChip id="C20" /></div>
        <div class="demo">
          <div class="accent-picker" role="group" aria-label="Accent">
            {#each picker as p}
              <button data-ac={p.ac} style="background:{p.ac}" aria-pressed={dboardAccent === p.ac} onclick={() => pick(p)} aria-label="Accent"></button>
            {/each}
          </div>
          <DaughterBoardHeader accent={dboardAccent} from={dboardFrom} crumb="Projects / Marketing" title="Backend migration"
            description="This board lives inside its mother ticket. Its lanes, tags and members are scoped here — and it wears the ticket's colour so you always know where you are."
            stats={[{ value: 3, label: 'columns' }, { value: 8, label: 'tickets' }, { value: 4, label: 'members' }]} />
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Project tile</h3><p>One board in the projects folder. Special ones glow.</p></div><IdChip id="C21" /></div>
        <div class="demo">
          <ProjectTile special accent="var(--ac-rose)" name="Product launch" icon="folder" description="Cross-team launch tracking"
            stats={[{ value: 24, label: 'tickets' }, { value: 4, label: 'lanes' }]}
            assignees={[{ initials: 'VX', tone: 'a1' }, { initials: 'JD', tone: 'a3' }]} />
        </div>
      </div>

      <div class="tile span2">
        <div class="tile-head"><div class="meta"><h3>Projects folder</h3><p>The grid users land on. Mix of standard and accented boards.</p></div><IdChip id="C22" /></div>
        <div class="demo">
          <div class="grid" style="grid-template-columns:repeat(auto-fill,minmax(210px,1fr));gap:14px">
            <ProjectTile name="Engineering" icon="board" stats={[{ value: 61, label: 'tickets' }]} />
            <ProjectTile special accent="var(--ac-iris)" name="Marketing Q3" icon="board" stats={[{ value: 12, label: 'tickets' }]} />
            <ProjectTile name="Design system" icon="board" stats={[{ value: 18, label: 'tickets' }]} />
          </div>
        </div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Sidebar / nav</h3><p>Boards, projects, settings, and the user chip with logout.</p></div><IdChip id="C23" /></div>
        <div class="demo"><Sidebar active="Boards" /></div>
      </div>

      <div class="tile span2">
        <div class="tile-head"><div class="meta"><h3>Login</h3><p>Password plus a two-factor step — authenticator code or security key.</p></div><IdChip id="C24" /></div>
        <div class="demo" style="align-items:center"><LoginCard /></div>
      </div>

      <div class="tile span2">
        <div class="tile-head"><div class="meta"><h3>Settings panel</h3><p>Appearance, notifications, and the logout / danger zone.</p></div><IdChip id="C25" /></div>
        <div class="demo"><SettingsPanel onlogout={() => (modalOpen = true)} /></div>
      </div>
    </div>
  </section>

  <!-- OVERLAYS & FEEDBACK -->
  <section class="section" id="overlays">
    <div class="eyebrow">Overlays &amp; feedback</div>
    <div class="grid">
      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Modal — log out confirm</h3><p>Click to open the real dialog.</p></div><IdChip id="C26" /></div>
        <div class="demo"><Button variant="secondary" onclick={() => (modalOpen = true)}>Preview log-out dialog</Button></div>
      </div>

      <div class="tile">
        <div class="tile-head"><div class="meta"><h3>Toast</h3><p>Transient confirmation, bottom-right.</p></div><IdChip id="C27" /></div>
        <div class="demo"><Button variant="secondary" onclick={() => pushToast('Ticket moved', 'RAP-204 → In progress')}>Trigger a toast</Button></div>
      </div>

      <div class="tile span2">
        <div class="tile-head"><div class="meta"><h3>Comment thread</h3><p>Threaded via parent reference — replies nest under their parent.</p></div><IdChip id="C28" /></div>
        <div class="demo"><CommentThread comments={thread} /></div>
      </div>
    </div>
  </section>

  <footer>
    <span class="f-mono">RecApparō · component gallery · 28 components</span>
    <span>Click an <strong>ID</strong> to copy · then tell me what to change</span>
  </footer>
</main>

<Modal bind:open={modalOpen}>
  {#snippet children(close)}
    <div class="m-icon"><Icon name="logout" size={23} /></div>
    <h3 id="modalTitle">Log out of RecApparō?</h3>
    <p>You'll be signed out on this device. Any unsaved ticket drafts stay in your browser.</p>
    <div class="m-actions">
      <Button variant="ghost" onclick={close}>Stay signed in</Button>
      <Button variant="danger" onclick={close}>Log out</Button>
    </div>
  {/snippet}
</Modal>
