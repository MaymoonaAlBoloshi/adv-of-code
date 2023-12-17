import fs from 'fs'

const readFile = async () => {
  return await fs.promises.readFile("./d1-1.txt", "utf-8");
};

const main = async () => {
  const text = await readFile();
  
  const dataArray = text.split("\n");

  let totalNumber = 0;

  for (const line of dataArray) {
    const digits = line.match(/\d/g);

    if (digits) {
      const firstDigit = digits[0];
      const lastDigit = digits[digits.length - 1];

      const twoDigitNum = parseInt(firstDigit + lastDigit);
      totalNumber += twoDigitNum;
    }
  }

  console.log(totalNumber);
};

main();
