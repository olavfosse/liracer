def do_search(array, target_value): 
	min = 0
	max = array.length - 1
	N_of_guesses = 0

	while max >= min:
		guess = int((min + max) / 2);
		N_of_guesses+= 1
		if array[guess] == target_value:
			print(N_of_guesses);
			return guess
		elif array[guess] < target_value:
			min = guess + 1
		else:
			max = guess - 1

	return -1