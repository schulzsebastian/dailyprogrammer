def largest_digit(num):
    return int(sorted(str(num).zfill(4), reverse=True)[0])

def desc_digits(num):
    return int("".join(sorted(str(num).zfill(4), reverse=True)))

def kaprekar(num, i=0):
    while num not in [0, 6174]:
        num = int("".join(sorted(str(num).zfill(4), reverse=True))) - int("".join(sorted(str(num).zfill(4))))
        i += 1
    return i

secret_number = 0
for i in range(0, 9999):
    if kaprekar(i) > secret_number:
        secret_number = i
print(secret_number)
