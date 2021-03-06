import * as readline from "readline";
import * as fs from "fs";

const readInterface = readline.createInterface({
  input: fs.createReadStream("01_1__input.txt"),
  output: process.stdout,
  terminal: false
});

const calculateFuelRequired = (moduleMass: number) =>
  Math.floor(moduleMass / 3) - 2;

let total = 0;

readInterface.on("line", line => {
  let fuel = calculateFuelRequired(Number(line));
  while (fuel > 0) {
    total += fuel;
    fuel = calculateFuelRequired(Number(fuel));
  }
});
readInterface.once("close", () => {
  console.log(`total fuel required: ${total}`);
});
