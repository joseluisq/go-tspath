import { measure } from "~/lib/measure"
import { sleep } from "~/lib/sleep"

measure(() => sleep(2 * 1e3))

console.log('Execution line no-blocked')
