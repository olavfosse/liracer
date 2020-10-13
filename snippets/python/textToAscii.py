originalString = "Test"

def stringToAscii(string):
	NumberDec = []
	for i in string:
		NumberDec.append(ord(i))
	message = ""
	message += "Decimal numbers: "
	for i in NumberDec:
		message += str(i) + ", "
	message += "\n"

	return message;

print(stringToAscii(originalString))