function fizzbuzz(n) {
	for (i = 0; i < n+1; i++) {
		let message = "";
		if (i % 3 == 0) {
			message += "Fizz";
		}
		if (i % 5 == 0) {
			message += "Buzz";
		}
		if (message == "") {
			message = i;
		}

		console.log(message);
	}
}

fizzbuzz(30);