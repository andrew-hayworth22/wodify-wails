<script lang="ts">
    import '../style.css';
    import { onMount, tick } from 'svelte';
    import { config } from '#lib/stores/config.ts';
    import { WindowShow } from '../../wailsjs/runtime';
    import logo from '../assets/images/appicon.png';

    let { children } = $props();

    const SPLASH_MIN_MS = 3000;

    const systemDarkQuery = window.matchMedia('(prefers-color-scheme: dark)');
    let systemDark = $state(systemDarkQuery.matches);
    let configLoaded = $state(false);

    onMount(() => {
        const onChange = (e: MediaQueryListEvent) => (systemDark = e.matches);
        systemDarkQuery.addEventListener('change', onChange);

        tick().then(WindowShow);
        Promise.all([
            config.load().catch((err) => console.error('loading config:', err)),
            new Promise((resolve) => setTimeout(resolve, SPLASH_MIN_MS)),
        ]).then(() => (configLoaded = true));

        return () => systemDarkQuery.removeEventListener('change', onChange);
    })

    $effect(() => {
        const theme = $config.theme;
        const dark = theme === 'dark' || ((theme === 'system' || !theme) && systemDark);
        document.documentElement.classList.toggle('dark', dark)
    })
</script>

{#if configLoaded}
    {@render children()}
{:else}
    <div class="flex h-full flex-col items-center justify-center gap-8">
        <img src={logo} alt="Wodify" class="h-auto w-48" />
        <div
            class="size-8 animate-spin rounded-full border-4 border-gray-200 border-t-gray-800 dark:border-gray-700 dark:border-t-gray-100"
            role="status"
            aria-label="Loading"
        ></div>
    </div>
{/if}