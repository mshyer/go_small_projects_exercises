function hasDuplicateValue(array) {
  let steps = 0;
  let existingNumbers = [];
  for(let i = 0; i < array.length; i++) {
    steps++;
    if(existingNumbers[array[i]] === 1) {
      return true;
    } else {
      existingNumbers[array[i]] = 1;
    }
  }
  console.log(steps);
  return false;
}

hasDuplicateValue([1, 2, 3, 4, 5, 6, 7, 8])


var myArr = [
  1, 2, 3,
  [4, 5, 6],
  7,
  [8,
    [9, 10, 11,
      [12, 13, 14]
    ]
  ],
  [15, 16, 16, 18, 19, 
    [20, 21, 22,
      [23, 24, 25, 
        [26, 27, 29]
      ], 30, 31
    ], 32
  ], 33
]