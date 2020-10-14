def doSearch(array, targetValue): 
	min = 0;
	max = array.length - 1;
	NOfGuesses = 0;

	while max >= min:
		guess = int((min + max) / 2);
		NOfGuesses+= 1
		if array[guess] == targetValue:
			print(NOfGuesses);
			return guess;
		elif array[guess] < targetValue:
			min = guess + 1;
		else:
			max = guess - 1;

	return -1;