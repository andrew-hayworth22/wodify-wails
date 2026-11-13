import { redirect } from '@sveltejs/kit';
import { HasAPIToken } from '../../../wailsjs/go/app/App';
import type { LayoutLoad } from './$types';
import { AUTH_DEP } from '#lib/auth.ts';

// Guards every route in the (app) group: no API token → /welcome.
export const load: LayoutLoad = async ({ depends }) => {
    depends(AUTH_DEP);

    if (!(await HasAPIToken())) {
        redirect(307, '/welcome');
    }
};