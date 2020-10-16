def check_prime_number(n):
    isPrime = True if n<2 else False
    for i in range(2, n):
        if n % i == 0:
            isPrime = False
            print("Number is not prime")
            return 0
        
    print("Number is prime")

n = input("Enter number here: ")
check_prime_number(n)