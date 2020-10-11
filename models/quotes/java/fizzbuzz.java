public void fizzbuzz(int n) {
	for (int i = 0; i < n+1; i++) {
		String message = "";
		if (i % 3 == 0) {
			message+="Fizz";
		}
		if (i%5 == 0) {
			message+="Buzz";
		}
		if (message == "") {
			message = Integer.toString(i);
		}
		System.out.println(message);
	}
}