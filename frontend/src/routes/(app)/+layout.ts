import { redirect } from '@sveltejs/kit';
import { HasAPIToken } from '../../../wailsjs/go/app/App';
import type { LayoutLoad } from './$types';

// Guards every route in the (app) group: no API token → /welcome.
// Call invalidate('app:auth') after the token changes to re-run this check.
export const load: LayoutLoad = async ({ depends }) => {
    depends('app:auth');

    if (!(await HasAPIToken())) {
        redirect(307, '/welcome');
    }
};