import { readFileSync } from 'fs';

const file = readFileSync('../resources/day1.txt','utf8')

let elfCalories = <number[]>[]
let currentElf = 0
file.split(/\r?\n/).forEach(function(line){
    if(line === '') {
        elfCalories.push(currentElf)
        currentElf = 0
    } else {
        currentElf += Number.parseInt(line)
    }
  })

elfCalories = elfCalories.sort((n1,n2) => (n1 - n2) * -1);
console.log(`Solution A: ${elfCalories[0]}`)
console.log(`Solution B: ${elfCalories[0] + elfCalories[1] + elfCalories[2]}`)

/**
 * Solutions
 * 
 * a: 69883
 * b: 207576
 */