public static int doSearch (int[] array, int targetValue) {
	int min = 0;
	int max = array.length - 1;
	int guess;
	int NOfGuesses = 0;	

	while (max >= min) {
		guess = (int) ((min + max) / 2);
		NOfGuesses++;
		if (array[guess] == targetValue) {
			System.out.println(NOfGuesses);
			return guess;
		} else if (array[guess] < targetValue) {
			min = guess + 1;
		} else {
			max = guess - 1;
		}
	}	
	return -1;
};