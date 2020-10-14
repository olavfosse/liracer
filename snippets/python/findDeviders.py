def find_deviders (number):
	deviders = [];
	for i in range(1, number+1):
		if number % i == 0:
			deviders.append(i);

	return deviders;

print(find_deviders(12))