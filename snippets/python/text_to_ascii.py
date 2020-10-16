originalString = "Test"

def string_to_ascii(string):
	number_dec = []
	for i in string:
		number_dec.append(ord(i))
	message = ""
	message += "Decimal numbers: "
	for i in number_dec:
		message += str(i) + ", "
	message += "\n"

	return message

print(string_to_ascii(originalString))