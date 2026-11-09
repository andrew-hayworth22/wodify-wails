<script lang="ts">
    import '../style.css';
    import { onMount } from 'svelte';
    import { config } from '#lib/stores/config.ts';

    let { children } = $props();

    onMount(() => {
        config.load();
    })

    $effect(() => {
        document.documentElement.classList.toggle('dark', $config.dark_mode)
    })
</script>

{@render children()}