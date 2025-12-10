/**
 * @author Natnael Debesay
 * @version 1.0.0
 * @date 2025-12-09
 * @fileoverview This program prints the ASCII value between two prompted numbers between 32 and 126.
 */

// variables
let startingCharacterASCIINumber: number = 0;
let endingCharacterASCIINumber: number = 0;

// input the numbers
startingCharacterASCIINumber = parseInt(prompt("Please enter a starting number larger than 32, and less than 126: ") || "0");
endingCharacterASCIINumber = parseInt(prompt("Please enter an ending number larger than 32, and less than 126: ") || "0")
// print the numbers
for (let counter: number = startingCharacterASCIINumber; counter <= endingCharacterASCIINumber; counter++) {
    console.log(String.fromCharCode(counter));
}

console.log("\nDone.");
