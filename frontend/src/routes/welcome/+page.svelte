<script lang="ts">
    import { goto } from '$app/navigation';
    import { SetAPIToken } from '../../../wailsjs/go/app/App';

    let token = $state('');
    let saving = $state(false);
    let error = $state('');

    async function submit(event: SubmitEvent) {
        event.preventDefault();
        error = '';
        saving = true;
        try {
            await SetAPIToken(token);
            token = '';
            await goto('/');
        } catch (err) {
            // Wails rejects with the Go error message as a string
            error = String(err);
        } finally {
            saving = false;
        }
    }
</script>

<main class="flex h-full items-center justify-center px-4">
    <div class="w-full max-w-md">
        <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">Welcome to Wodify</h1>
        <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">
            Enter your Wodify API key to get started. It's stored securely in your system keychain.
        </p>

        <form class="mt-6 space-y-4" onsubmit={submit}>
            <div>
                <label for="api-token" class="block text-sm font-medium text-gray-900 dark:text-gray-100">
                    API key
                </label>
                <input
                    id="api-token"
                    type="password"
                    autocomplete="off"
                    bind:value={token}
                    disabled={saving}
                    class="mt-1 block w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-gray-900
                           focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500
                           disabled:opacity-50 dark:border-gray-700 dark:bg-gray-800 dark:text-white"
                />
            </div>

            {#if error}
                <p class="text-sm text-red-600 dark:text-red-400" role="alert">{error}</p>
            {/if}

            <button
                type="submit"
                disabled={saving || token.trim() === ''}
                class="w-full rounded-md bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-500
                       disabled:cursor-not-allowed disabled:opacity-50"
            >
                {saving ? 'Verifying…' : 'Continue'}
            </button>
        </form>
    </div>
</main>