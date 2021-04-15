for(i = 0; i <= 100; i++) {
	let s = ""
	if(i % 3 === 0) {
		s += "Fizz"
	}
	if(i % 5 === 0) {
		s += "Buzz"
	}
	console.log(s !== "" ? s : i)
}
