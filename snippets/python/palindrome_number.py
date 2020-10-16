n = input("Enter a number: ")
temp = n
reversed_number = 0

while temp>0:
    last_digit = temp % 10
    reversed_number = reversed_number * 10 + last_digit
    temp /= 10
    
if reversed_number == n:
    print(n, "is a palindrome")
else:
    print("n", "is not a palindrome")