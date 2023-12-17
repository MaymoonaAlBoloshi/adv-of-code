import fs from "fs";

const readFile = async () => {
  return await fs.promises.readFile("./d1-2.txt", "utf-8");
};

/*
 * this is handy to in case of words like eightwo,
 * where if I replace 2 then we'll have eigh which is missing a letter,
 * if we leave prevouse word before and after teh replaced number
 * this will help us preserve other numbers
 * that might the same ending or bigging as the replaced number spelling
 */

const numberMappings = {
  one: "one1one",
  two: "two2two",
  three: "three3three",
  four: "four4four",
  five: "five5five",
  six: "six6six",
  seven: "seven7seven",
  eight: "eight8eight",
  nine: "nine9nine",
};

const main = async () => {
  const text = await readFile();
  let total = 0;

  let data = text.split("\n");

  for (let line of data) {
    for (let num of Object.keys(numberMappings)) {
      line = line.replaceAll(num, numberMappings[num]);
    }
    const digits = line.match(/\d/g);

    if (digits) {
      const firstDigit = digits[0];
      const lastDigit = digits[digits.length - 1];

      const twoDigitNum = parseInt(firstDigit + lastDigit);
      total += twoDigitNum;
    }
  }

  console.log(total);
};

main();
