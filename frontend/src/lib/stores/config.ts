import { writable } from "svelte/store";
import { config as ConfigModels } from "../../../wailsjs/go/models";
import { GetConfig, SetConfig, SaveConfig} from "../../../wailsjs/go/app/App";

const { subscribe, update } = writable(new ConfigModels.Config());

export const config = {
    subscribe,
    load: async () => {
        const cfg = await GetConfig();
        update(() => cfg)
    },
    setDarkMode: (darkMode: boolean) => {
        update((c) => {
            c.dark_mode = darkMode;
            return c;
        })
    }
};