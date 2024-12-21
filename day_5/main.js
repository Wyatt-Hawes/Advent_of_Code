import * as fs from 'node:fs';
import {exit} from 'node:process';

const fname = './input.txt';
// const fname = './test.txt';

const input_r = fs.readFileSync(fname, 'utf-8');
const input = input_r.split('\r\n');
const rules = {};
const updates = [];

// FIRST CREATE INPUT, make dictionary of rules (reversed)
let i = 0;
while (input[i] != '') {
  let vals = input[i].split('|');
  vals = vals.map((val) => Number(val));
  if (!rules[vals[1]]) {
    rules[vals[1]] = [];
  }
  rules[vals[1]].push(vals[0]);
  i++;
}

i++;
while (input[i]) {
  let vals = input[i].split(',');
  updates.push(vals.map((val) => Number(val)));
  i++;
}

// Now check the orderings
const valid_orderings = [];
let orderings_to_fix = {};
updates.forEach((update) => {
  i = 0;
  let seen = {};
  let valid_ordering = true;

  seen = calculate_seen(update);

  for (let j = 0; j < update.length; j++) {
    let num = update[j];
    const num_rules = rules[num];
    num_rules?.forEach((rule) => {
      // Make sure index of rule is before our number, if not, bad
      if (seen[rule] > seen[num]) {
        valid_ordering = false;
        // Lets keep track of orderings to fix for part 2
        orderings_to_fix[update] = true;
      }
    });
  }

  if (valid_ordering) {
    console.log('===VALID===');
    valid_orderings.push(update);
    return;
  }
  console.log('===INVALID===');
});

// Now lets turn the orderings we need to fix into a [][] of numbers
orderings_to_fix = Object.keys(orderings_to_fix);
orderings_to_fix = orderings_to_fix.map((v) =>
  v.split(',').map((v2) => Number(v2))
);

// Loop over all orders we need to fix
for (let x = 0; x < orderings_to_fix.length; x++) {
  let seen = calculate_seen(orderings_to_fix[x]);

  // Look at every number
  for (let j = 0; j < orderings_to_fix[x].length; j++) {
    let num = orderings_to_fix[x][j];
    const num_rules = rules[num];
    let swapped = false;
    // Look at all rules for this number, are there any invalid orderings?
    num_rules?.forEach((rule) => {
      // If an invalid ordering, swap them
      if (!swapped && seen[rule] > seen[num]) {
        orderings_to_fix[x][seen[rule]] = num;
        orderings_to_fix[x][seen[num]] = rule;
        seen = calculate_seen(orderings_to_fix[x]);
        j = -1; // -1 since j gets incremented to 0 on loop end
        swapped = true;
        // Lets keep track of orderings to fix for part 2
      }
    });
  }
}

let puzzle_input = 0;
valid_orderings.forEach((update) => {
  const index = Math.floor(update.length / 2);
  puzzle_input += update[index];
});

let fixed_input = 0;
orderings_to_fix.forEach((update) => {
  const index = Math.floor(update.length / 2);
  fixed_input += update[index];
});

console.log('Solution: ', puzzle_input);
console.log('Fixed Solution: ', fixed_input);

function calculate_seen(order) {
  const seen = {};
  let i = 0;
  order.forEach((num) => {
    seen[num] = i;
    i++;
  });
  return seen;
}
