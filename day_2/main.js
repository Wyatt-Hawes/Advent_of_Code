import * as fs from 'node:fs';
const input = fs.readFileSync('./input.txt', 'utf-8');

const reports = input.split('\n');
let safe_reports_1 = 0;
let safe_reports_2 = 0;

reports.forEach((report) => {
  report = report.split(' ');
  report = report.map((val) => Number(val));

  if (is_safe(report)) {
    safe_reports_1++;
    safe_reports_2++;
    return;
  }

  // Can we remove an index and make it safe?
  for (let i = 0; i < report.length; i++) {
    const tmp_report = report.slice();
    tmp_report.splice(i, 1);
    if (is_safe(tmp_report)) {
      safe_reports_2++;
      return;
    }
  }
});

console.log(`Safe Reports: ${safe_reports_1}`);
console.log(`Safe Reports Removing Index: ${safe_reports_2}`);

function is_safe(report) {
  // First check if increasing or decreasing
  let is_decreasing = false;
  if (report[1] < report[0]) {
    is_decreasing = true;
  } else if (report[1] > report[0]) {
    is_decreasing = false;
  } else {
    // First 2 values are equal so its invalid
    return false;
  }

  // Check distance between every value
  for (let i = 0; i < report.length - 1; i++) {
    // Do A-B or B-A depending on if we expect increasing or decreasing
    const diff = is_decreasing
      ? report[i] - report[i + 1]
      : report[i + 1] - report[i];

    if (diff >= 1 && diff <= 3) {
      continue;
    }
    return false;
  }

  return true;
}
