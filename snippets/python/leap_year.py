def isLeapYear(year):
    if year % 4 == 0:
        print(year, "is a leap year")
    else:
        print(year, "is not a leap year.")
        
year = input("Enter a year: ")
isLeapYear(int(year))