const do_search = function (array, target_value) {
	let min = 0;
	let max = array.length - 1;
	let guess;
	let N_of_guesses = 0;	

	while (max >= min) {
		guess = Math.floor((min + max) / 2);
		N_of_guesses++;
		if (array[guess] === target_value) {
			println(N_of_guesses);
			return guess;
		} else if (array[guess] < target_value) {
			min = guess + 1;
		} else {
			max = guess - 1;
		}
	}	

	return -1;
};