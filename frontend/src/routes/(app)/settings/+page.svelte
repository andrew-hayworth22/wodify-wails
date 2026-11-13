<script lang="ts">
    import { goto } from '$app/navigation';
    import { get } from 'svelte/store';
    import { config, type Theme } from '#lib/stores/config.ts';
    import { config as ConfigModels } from '../../../../wailsjs/go/models';
    import { SignOut } from '../../../../wailsjs/go/app/App';
    import Button from '../../../components/Button.svelte';
    import Input from '../../../components/Input.svelte';

    // Go's time.Duration is serialized as nanoseconds; the form edits seconds.
    const NS_PER_SECOND = 1_000_000_000;

    // Edit a draft copy so nothing (including the theme) applies until Save.
    // The root layout only renders pages after the config has loaded.
    const initial = get(config);
    let theme = $state<Theme>((initial.theme as Theme) || 'system');
    let baseURL = $state(initial.base_url);
    let timeoutSeconds = $state(initial.timeout / NS_PER_SECOND);

    let saving = $state(false);
    let saved = $state(false);
    let error = $state('');
    let signingOut = $state(false);

    async function save(event: SubmitEvent) {
        event.preventDefault();
        error = '';
        saved = false;
        saving = true;
        try {
            await config.save(
                new ConfigModels.Config({
                    ...get(config),
                    theme,
                    base_url: baseURL.trim(),
                    timeout: Math.round(timeoutSeconds * NS_PER_SECOND),
                }),
            );
            saved = true;
        } catch (err) {
            error = String(err);
        } finally {
            saving = false;
        }
    }

    async function signOut() {
        error = '';
        signingOut = true;
        try {
            await SignOut();
            await goto('/welcome');
        } catch (err) {
            error = String(err);
            signingOut = false;
        }
    }
</script>

<div class="max-w-md">
    <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">Settings</h1>

    <form class="mt-6 space-y-4" onsubmit={save} oninput={() => (saved = false)}>
        <div>
            <label for="theme" class="block text-sm font-medium text-gray-900 dark:text-gray-100">Theme</label>
            <select
                id="theme"
                bind:value={theme}
                disabled={saving}
                class="mt-1 block w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-gray-900
                       focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500
                       disabled:opacity-50 dark:border-gray-700 dark:bg-gray-800 dark:text-white"
            >
                <option value="system">System</option>
                <option value="light">Light</option>
                <option value="dark">Dark</option>
            </select>
        </div>

        <div>
            <label for="base-url" class="block text-sm font-medium text-gray-900 dark:text-gray-100">
                API base URL
            </label>
            <Input id="base-url" type="url" required bind:value={baseURL} disabled={saving} />
        </div>

        <div>
            <label for="timeout" class="block text-sm font-medium text-gray-900 dark:text-gray-100">
                Request timeout (seconds)
            </label>
            <Input id="timeout" type="number" min="0" step="any" required bind:value={timeoutSeconds} disabled={saving} />
            <p class="mt-1 text-xs text-gray-500 dark:text-gray-400">0 uses the client default.</p>
        </div>

        {#if error}
            <p class="text-sm text-red-600 dark:text-red-400" role="alert">{error}</p>
        {/if}

        <div class="flex items-center gap-3">
            <Button type="submit" disabled={saving}>{saving ? 'Saving…' : 'Save'}</Button>
            {#if saved}
                <span class="text-sm text-gray-600 dark:text-gray-400" role="status">Saved</span>
            {/if}
        </div>
    </form>

    <div class="mt-10 border-t border-gray-200 pt-6 dark:border-gray-800">
        <h2 class="text-lg font-medium text-gray-900 dark:text-white">Account</h2>
        <p class="mt-1 text-sm text-gray-600 dark:text-gray-400">
            Signing out removes your API key from this computer's keychain.
        </p>
        <Button variant="secondary" class="mt-4" onclick={signOut} disabled={signingOut}>
            {signingOut ? 'Signing out…' : 'Sign out'}
        </Button>
    </div>
</div>