import random

alphabet_char = [chr(i) for i in range(ord('a'), ord('z')+1)] + [chr(i) for i in range(ord('A'), ord('Z')+1)]
weird_char = [
    '±', 'µ', 'Ω', 'ß', 'ç', 'ñ', '∆', '•', '♪', '★', '✦', '☯', '☃', '☠', '🙂',
    '©', '®', '€', '£', '¥', '§', '¶', '¿', '¡', '°', '·', '—', '…', '‽'
]
numbers = [int(i) for i in range(1, 110)]

stringfy_alphabet_char = str(random.choice(alphabet_char)) 
stringfy_weird_char = str(random.choice(weird_char))

password_length = 20
password = ""

for char in range(password_length):
    password += str(random.choice(alphabet_char)) + str(random.choice(weird_char))
    
    
print(password)