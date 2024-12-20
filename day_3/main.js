import {match} from 'node:assert';
import exp from 'node:constants';
import * as fs from 'node:fs';

const input = fs.readFileSync('./input.txt', 'utf-8');

// Match the following:
// mul(
// 1+ of digits 0-9
// a literal ,
// 1+ of digits 0-9
// )
// Stores the 2 numbers in match[1] and match[2]
const regex_part_1 = /mul[(]([0-9]+),([0-9]+)[)]/g;

// Same as above but OR "do()" and "don't()" literal matches
const regex_part_2 = /mul[(]([0-9]+),([0-9]+)[)]|do[(][)]|don't[(][)]/g;

// Get all matches
const expressions_1 = [...input.matchAll(regex_part_1)];
const expressions_2 = [...input.matchAll(regex_part_2)];

// Iterate and apply all matches
let total_1 = 0;
expressions_1.forEach((match) => {
  // Stores the 2 numbers in match[1] and match[2]
  total_1 += match[1] * match[2];
});

console.log('Total: ' + total_1);
console.log('=============');

let total_2 = 0;
let enabled = true;
expressions_2.forEach((match) => {
  // match[0] contains matching string
  if (match[0] == 'do()') {
    enabled = true;
    return;
  }
  if (match[0] == "don't()") {
    enabled = false;
    return;
  }

  // Stores the 2 numbers in match[1] and match[2]
  if (enabled) {
    total_2 += match[1] * match[2];
  }
});

console.log(`Total w/ do() & don't(): ` + total_2);
console.log('=============');
