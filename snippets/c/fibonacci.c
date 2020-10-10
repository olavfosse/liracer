#include <stdio.h>

int fibonacci(int);

int main() {
	int n;
	scanf("%i", &n);
	printf("%i\n", fibonacci(n));
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