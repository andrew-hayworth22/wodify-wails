import { redirect } from '@sveltejs/kit';
import { HasAPIToken } from '../../../wailsjs/go/app/App';
import type { PageLoad } from './$types';

// Already have a token → skip the welcome screen.
export const load: PageLoad = async () => {
    if (await HasAPIToken()) {
        redirect(307, '/');
    }
};