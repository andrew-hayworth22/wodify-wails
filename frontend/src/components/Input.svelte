<script lang="ts">
    import type { HTMLInputAttributes } from 'svelte/elements';

    let { value = $bindable(), class: className = '', ...rest }: HTMLInputAttributes = $props();
</script>

<input
    bind:value
    class="mt-1 block w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-gray-900
           focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500
           disabled:opacity-50 dark:border-gray-700 dark:bg-gray-800 dark:text-white {className}"
    {...rest}
/>