"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var measure_1 = require("~/lib/measure");
var sleep_1 = require("~/lib/sleep");
measure_1.measure(function () { return sleep_1.sleep(2 * 1e3); });
console.log('Execution line no-blocked');
