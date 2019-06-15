import { time } from '~/utils/utils';

/** measure execution of one function in milliseconds */
export async function measure (fn: Function) {
  const start = time()

  await fn()

  const end = time() - start

  console.log('Execution time: %dms', end)
}
