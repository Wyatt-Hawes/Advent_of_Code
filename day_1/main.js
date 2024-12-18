import * as fs from 'node:fs';

const input = fs.readFileSync('./input.txt', 'utf-8');
const r1 = [];
const r2 = [];

// Make 2 arrays from the columns
const rows = input.split('\n');
rows.forEach((r) => {
  const row = r.split('   ');
  r1.push(row[0]);
  r2.push(row[1]);
});

// Sort them
r1.sort((a, b) => a - b);
r2.sort((a, b) => a - b);

// Calculate distance between both values
let total_distance = 0;
let i = 0;
r1.forEach((val) => {
  total_distance += Math.abs(val - r2[i]);
  i++;
});

console.log(`Total Distance: ${total_distance}`);

// Part 2
// Create count of occurrences (Yes we could combine it above but its less confusing this way)
const r2_duplicate_count = {};
r2.forEach((val) => {
  val = Number(val);

  // If value doesnt exist, initialize to 0
  if (!r2_duplicate_count[val]) {
    r2_duplicate_count[val] = 0;
  }
  r2_duplicate_count[val]++;
});

let similarity_score = 0;
r1.forEach((val) => {
  val = Number(val);
  // If value exists, use it, else 0
  const occurr = r2_duplicate_count[val] ? r2_duplicate_count[val] : 0;
  similarity_score += val * occurr;
});

console.log(`Similarity Score: ${similarity_score}`);
