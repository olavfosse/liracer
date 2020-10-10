def fibonacci_sequence(n)
	return [0, 1].take(n) if n <= 2
	start = fibonacci_sequence(n-1)
	start + [start[-2..-1].sum]
end

puts fibonacci_sequence(gets.to_i)