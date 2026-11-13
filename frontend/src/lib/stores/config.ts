import { writable } from "svelte/store";
import { config as ConfigModels } from "../../../wailsjs/go/models";
import { GetConfig, SetConfig, SaveConfig} from "../../../wailsjs/go/app/App";

export type Theme = 'system' | 'light' | 'dark';

const { subscribe, set } = writable(new ConfigModels.Config());

export const config = {
    subscribe,
    load: async () => {
        set(await GetConfig());
    },
    // Applies cfg in the backend, writes it to disk, then publishes it to subscribers.
    save: async (cfg: ConfigModels.Config) => {
        await SetConfig(cfg);
        await SaveConfig();
        set(cfg);
    },
};