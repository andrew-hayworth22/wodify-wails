<script lang="ts">
    import { onMount } from 'svelte';
    import { invalidate } from '$app/navigation';
    import { EventsOn } from '../../../wailsjs/runtime';
    import { AUTH_DEP, AUTH_EXPIRED_EVENT } from '#lib/auth.ts';

    let { children } = $props();

    onMount(() => EventsOn(AUTH_EXPIRED_EVENT, () => invalidate(AUTH_DEP)));
</script>

<!-- Placeholder shell for authenticated routes -->
<div class="flex h-full flex-col">
    <header class="flex items-center gap-6 border-b border-gray-200 px-4 py-3 dark:border-gray-800">
        <span class="font-semibold text-gray-900 dark:text-white">Wodify</span>
        <nav class="flex gap-4 text-sm text-gray-600 dark:text-gray-400">
            <a href="/" class="hover:text-gray-900 dark:hover:text-white">Home</a>
            <a href="/settings" class="hover:text-gray-900 dark:hover:text-white">Settings</a>
        </nav>
    </header>

    <main class="flex-1 overflow-auto p-4">
        {@render children()}
    </main>
</div>