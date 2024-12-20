import * as fs from 'node:fs';
const input_r = fs.readFileSync('./input.txt', 'utf-8');

const input = input_r.split('\n');

// First create the 2d array
let size = input.length;
let rows = input[0].length;
let ceres = Array(rows)
  .fill()
  .map(() => Array(size));

// Now fill it with data
let i = 0;
input.forEach((row) => {
  for (let j = 0; j < row.length; j++) {
    // Y , X
    ceres[i][j] = row[j];
  }
  i++;
});

// Now we are going to iterate over every single cell and check each direction + diagonals
let total_matches = 0;
for (let x = 0; x < size; x++) {
  for (let y = 0; y < size; y++) {
    // Now check all directions, we only increment if the check_direction is true
    total_matches += check_directions(x, y, -1, -1, 0) ? 1 : 0; // Top left
    total_matches += check_directions(x, y, 0, -1, 0) ? 1 : 0; // Top
    total_matches += check_directions(x, y, 1, -1, 0) ? 1 : 0; // Top right

    total_matches += check_directions(x, y, 1, 0, 0) ? 1 : 0; // Right
    total_matches += check_directions(x, y, -1, 0, 0) ? 1 : 0; // Left

    total_matches += check_directions(x, y, -1, 1, 0) ? 1 : 0; // Bottom Left
    total_matches += check_directions(x, y, 0, 1, 0) ? 1 : 0; // Bottom
    total_matches += check_directions(x, y, 1, 1, 0) ? 1 : 0; // Bottom right
  }
}

// Now lets check for X-MAS
let total_x_matches = 0;
for (let x = 0; x < size; x++) {
  for (let y = 0; y < size; y++) {
    total_x_matches += check_x_mas(x, y) ? 1 : 0;
  }
}

// Safely Compare Cell
function compare_cell_safe(x, y, CHAR) {
  if (x < 0 || x >= size || y < 0 || y >= size) {
    return false;
  }
  return ceres[y][x] == CHAR;
}

function check_directions(x, y, dirX, dirY, stage) {
  const xmas = 'XMAS';
  // Out of bounds check
  if (x < 0 || x >= size || y < 0 || y >= size) {
    return false;
  }

  // Is this the correct letter?
  if (ceres[y][x] == xmas[stage]) {
    // Did we just check the last value?
    if (stage >= 3) {
      return true;
    }
    // If not, check the next letter in that direction, increment X/Y by direction & increase stage
    return check_directions(x + dirX, y + dirY, dirX, dirY, ++stage);
  }

  return false;
}

function check_x_mas(x, y) {
  // We must check for diagonal MAS or SAM for the \ direction (either works) AND the / Direction

  // Lets check from the center, center must be A
  if (!compare_cell_safe(x, y, 'A')) {
    return false;
  }

  // TL:S && BR:M  OR  TL:M && BR:S      AND    BL:S && TR:M   OR   BL:M  && TR:S
  if (is_TL_MAS(x, y) && is_TR_MAS(x, y)) {
    return true;
  }
  return false;
}
function is_TL_MAS(x, y) {
  return (
    (compare_cell_safe(x - 1, y - 1, 'S') &&
      compare_cell_safe(x + 1, y + 1, 'M')) ||
    (compare_cell_safe(x - 1, y - 1, 'M') &&
      compare_cell_safe(x + 1, y + 1, 'S'))
  );
}
function is_TR_MAS(x, y) {
  return (
    (compare_cell_safe(x + 1, y - 1, 'S') &&
      compare_cell_safe(x - 1, y + 1, 'M')) ||
    (compare_cell_safe(x + 1, y - 1, 'M') &&
      compare_cell_safe(x - 1, y + 1, 'S'))
  );
}

// Print helper function
function print_puzzle() {
  for (let i = 0; i < size; i++) {
    console.log(ceres[i].join(''));
  }
}
//print_puzzle();
console.log(`Total of ${total_matches} XMAS's`);
console.log(`Total of ${total_x_matches} X-MAS's`);
