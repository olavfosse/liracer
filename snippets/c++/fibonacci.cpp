#include <iostream>

int fibonacci(int);

int main() {
	int n;
	std::cin >> n;
	std::cout << fibonacci(n) << "\n";
}

int fibonacci(int n) {
	switch(n) {
		case 0:
		case 1:
			return n;
		default:
			return fibonacci(n - 1) + fibonacci(n - 2);
	}
}