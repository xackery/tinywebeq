{{ define "header" }}
<button id="menu-toggle" type="button" aria-label="Menu">
    <svg width="148" height="148" viewBox="0 0 148 148" xmlns="http://www.w3.org/2000/svg">
    <g stroke="currentColor" stroke-width="18" fill="none" fill-rule="evenodd" stroke-linecap="round">
        <path d="M9.5 125.5h129M9.5 74.5h129M9.5 23.5h129"></path>
    </g>
    </svg>
</button>
<div id="icon-toolbar">
    <a href="https://github.com/shoelace-style/shoelace" title="View Shoelace on GitHub" class="external-link" rel="noopener noreferrer" target="_blank">
    <sl-icon name="github" aria-hidden="true" library="default"></sl-icon>
    </a>

    <a href="https://twitter.com/shoelace_style" title="Follow Shoelace on Twitter" class="external-link" rel="noopener noreferrer" target="_blank">
    <sl-icon name="twitter" aria-hidden="true" library="default"></sl-icon>
    </a>

    <sl-dropdown id="theme-selector" placement="bottom-end" distance="3">
    <sl-button slot="trigger" size="small" variant="text" caret="" title="Press \ to toggle" data-optional="" data-valid="" data-user-valid="">
        <sl-icon class="only-light" name="sun-fill" aria-hidden="true" library="default"></sl-icon>
        <sl-icon class="only-dark" name="moon-fill" aria-hidden="true" library="default"></sl-icon>
    </sl-button>
    <sl-menu role="menu">
        <sl-menu-item type="checkbox" value="light" tabindex="-1" role="menuitemcheckbox" aria-checked="false" aria-disabled="false">Light</sl-menu-item>
        <sl-menu-item type="checkbox" value="dark" tabindex="-1" role="menuitemcheckbox" aria-checked="false" aria-disabled="false">Dark</sl-menu-item>
        <sl-divider role="separator" aria-orientation="horizontal"></sl-divider>
        <sl-menu-item type="checkbox" value="auto" tabindex="0" role="menuitemcheckbox" aria-checked="true" aria-disabled="false" checked="">System</sl-menu-item>
    </sl-menu>
    </sl-dropdown>
</div>
{{ end }}