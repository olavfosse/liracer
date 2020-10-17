"""
prompt
======

An input replacement.
"""

def prompt(text, _type):
 	"""
 	:param text: the text that appear in console.
 	:param _type: the type.
 	"""

 	accepted = False
 	while not accepted:
		try:
			result = _type(input(text))
			accepted = True
		except:
			pass
			
	return result
