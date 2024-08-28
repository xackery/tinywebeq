{{ define "sidebar" }}
<aside id="sidebar">
<header>
    <a href="/">
        <img src="/assets/images/wordmark.svg" alt="Shoelace">
    </a>
    <div class="sidebar-version">2.15.0</div>
</header>
<button class="search-box" type="button" title="Press / to search" aria-label="Search" data-plugin="search">
    <sl-icon name="search" aria-hidden="true" library="default"></sl-icon>
    <span>Search</span>
</button>
<nav>
    <ul>
        <li>
        <h2>Players</h2>
        <ul>
            <li><a href="/player/view?id=1">Player</a></li>
        </ul>
        </li>
        <li>
        <h2>Zones</h2>
        <ul>
            <li><a href="/frameworks/react">Zones by Era</a></li>
            <li><a href="/frameworks/vue">Zones by Level</a></li>
        </ul>
        </li>
        <li>
            <h2>Quests</h2>
                <ul>
                    <li><a href="/quest/view?id=1001">Quest</a></li>            
                </ul>
            </li>
        <li>
        <li>
            <h2>Items</h2>
                <ul>
                    <li><a href="/item/view?id=1001">Item Search</a></li>            
                </ul>
            </li>
        <li>
        <h2>Spells</h2>
        <ul>
            <li>
            <a href="/spell/view?id=1001">Spell Search</a>
            </li>
        </ul>
        </li>
        <li>
        <h2>Bestiary</h2>
        <ul>
            <li><a href="/npc/view?id=1001">NPC Search</a></li>
            <li><a href="/npc/view?id=1001">Advanced Npc Search</a></li>
            </ul>
        </li>
        <li>
        <h2>Tradeskills</h2>
        <ul>
            <li><a href="/tutorials/integrating-with-laravel">Recipe Search</a></li>
        </ul>
        </li>
    </ul>
</nav>
</aside>
{{end}}