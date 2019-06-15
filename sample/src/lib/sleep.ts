import { time } from '~/utils/utils';

/** non-blocking sleep function in milliseconds */
export async function sleep (delay: number) {
    const start = time()
    while (time() < start + delay);
}
