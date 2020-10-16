public static int do_search (int[] array, int target_value) {
	int min = 0;
	int max = array.length - 1;
	int guess;
	int N_of_guesses = 0;	

	while (max >= min) {
		guess = (int) ((min + max) / 2);
		N_of_guesses++;
		if (array[guess] == target_value) {
			System.out.println(N_of_guesses);
			return guess;
		} else if (array[guess] < target_value) {
			min = guess + 1;
		} else {
			max = guess - 1;
		}
	}	
	return -1;
}