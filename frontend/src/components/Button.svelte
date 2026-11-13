<script lang="ts">
    import type { HTMLButtonAttributes } from 'svelte/elements';

    type Props = HTMLButtonAttributes & {
        variant?: 'primary' | 'secondary';
    };

    let { variant = 'primary', type = 'button', class: className = '', children, ...rest }: Props = $props();

    const variants = {
        primary: 'bg-blue-600 text-white hover:bg-blue-500',
        secondary:
            'border border-gray-300 bg-white text-gray-900 hover:bg-gray-50 ' +
            'dark:border-gray-700 dark:bg-gray-800 dark:text-white dark:hover:bg-gray-700',
    };
</script>

<button
    {type}
    class="rounded-md px-4 py-2 font-medium disabled:cursor-not-allowed disabled:opacity-50 {variants[variant]} {className}"
    {...rest}
>
    {@render children?.()}
</button>