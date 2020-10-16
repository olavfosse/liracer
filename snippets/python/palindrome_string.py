def check(text):
    length = len(text)-1
    if text[0] != text[length]:
        print("String is not a palindrome")
    else:
        print("String is a palindrome")

text = input("Enter your text:")
check(text)