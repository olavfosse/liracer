let doSearch = function (array, targetValue) {
  let min = 0;
  let max = array.length - 1;
  let guess;
  let NOfGuesses = 0;

  while (max >= min) {
    guess = Math.floor((min + max) / 2);
    NOfGuesses++;
    if (array[guess] === targetValue) {
      println(NOfGuesses);
      return guess;
    } else if (array[guess] < targetValue) {
      min = guess + 1;
    } else {
      max = guess - 1;
    }
  }

  return -1;
};