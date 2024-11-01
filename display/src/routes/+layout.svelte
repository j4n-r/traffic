<script>
import { LayoutDashboard, Home, Settings, User } from "lucide-svelte";
let { children } = $props();
</script>

<div class="layout">
    <nav class="sidebar">
        <div class="sidebar-header">
            <h2 class="sidebar-title">Display</h2>
        </div>
        <ul class="sidebar-menu">
            <li><Home class="icon" /></li>
            <li><LayoutDashboard class="icon" /></li>
            <li><Settings class="icon" /></li>
            <li><User class="icon" /></li>
        </ul>
    </nav>
    <div class="header">
    </div>
    <div class="content">
        {@render children()}
    </div>
</div>

<style>
.layout {
    display: grid;
    grid-template-columns: 0.5fr 8fr;
    grid-template-rows: 0.5fr 8fr;
    height: 100vh;
}
.header {
    grid-column-start:2;
    background-color: var(--rp-base);
}

.sidebar {
    background-color: var(--rp-base);
    color: var(--rp-text);
    padding: 20px;
    display: flex;
    flex-direction: column;
    grid-row: span 2;
}

.sidebar-header {
    width: 100%;
    padding-bottom: 20px;
}

.sidebar-title {
    font-size: 2rem;
    color: var(--rp-iris);
}

.sidebar-menu {
    display: flex;
    flex-direction: column;
    align-items: center;
    list-style-type: none;
    margin-top: 20px;
    width: 100%;
    padding: 0;
    gap: 30px;
}

.content {
    padding: 20px;
    background-color: var(--rp-overlay);
    grid-column-start: 2;
    border-top-left-radius: 10px;
}
</style>
