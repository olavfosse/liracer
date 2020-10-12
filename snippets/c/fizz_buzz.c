#include <stdio.h>
#include <stdbool.h>

int main() {
	for(int i = 1; i <= 100; i++) {
		bool fizzed_or_buzzed = false;
		if(i % 3 == 0) {
			printf("Fizz");
			fizzed_or_buzzed = true;
		}
		if(i % 5 == 0) {
			printf("Buzz");
			fizzed_or_buzzed = true;
		}

		fizzed_or_buzzed ? puts("") : printf("%d\n", i);
	}
}