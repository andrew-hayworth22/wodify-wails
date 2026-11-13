<script lang="ts">
    import { goto } from '$app/navigation';
    import { SetAPIToken } from '../../../wailsjs/go/app/App';
    import Button from '../../components/Button.svelte';
    import Input from '../../components/Input.svelte';
    import logo from '../../assets/images/appicon.png';

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
            error = String(err);
        } finally {
            saving = false;
        }
    }
</script>

<main class="flex h-full items-center justify-center px-4">
    <div class="w-full max-w-md">
        <img src={logo} alt="" class="mx-auto h-auto w-40" />
        <h1 class="mt-8 text-center text-2xl font-semibold text-gray-900 dark:text-white">Welcome to Wodify</h1>
        <p class="mt-2 text-center text-sm text-gray-600 dark:text-gray-400">
            Enter your Wodify API key to get started. It's stored securely in your system keychain.
        </p>

        <form class="mt-6 space-y-4" onsubmit={submit}>
            <div>
                <label for="api-token" class="block text-sm font-medium text-gray-900 dark:text-gray-100">
                    API key
                </label>
                <Input
                    id="api-token"
                    type="password"
                    autocomplete="off"
                    bind:value={token}
                    disabled={saving}
                />
            </div>

            {#if error}
                <p class="text-sm text-red-600 dark:text-red-400" role="alert">{error}</p>
            {/if}

            <Button type="submit" class="w-full" disabled={saving || token.trim() === ''}>
                {saving ? 'Verifying…' : 'Continue'}
            </Button>
        </form>
    </div>
</main>