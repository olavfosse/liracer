function find_deviders(number) {
	let deviders = [];
	for (let i = 0; i <= number; i++) {
		if (number % i == 0) deviders.push(i);
	}

	return deviders;
}

console.log(find_deviders(12))