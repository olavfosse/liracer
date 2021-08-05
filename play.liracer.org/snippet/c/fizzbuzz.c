#include <stdio.h>

int main() {
	for(int i = 1; i <= 100; i++) {
		int fizzedOrBuzzed = 0;
		if(i % 3 == 0) {
			fizzedOrBuzzed = 1;
			printf("Fizz");
		}
		if(i % 5 == 0) {
			fizzedOrBuzzed = 1;
			printf("Buzz");
		}
		if(fizzedOrBuzzed) {
			printf("\n");
		} else {
			printf("%d\n", i);
		}
	}
}
