def largest_digit(num):
    return int(sorted(str(num).zfill(4), reverse=True)[0])


def desc_digits(num):
    return int("".join(sorted(str(num).zfill(4), reverse=True)))


def kaprekar(num, i=0):
    while num not in [0, 6174]:
        num = int("".join(sorted(str(num).zfill(4), reverse=True))) - int("".join(sorted(str(num).zfill(4))))
        print(num)
        i += 1
    return i

def secret_number():
    return max(set([kaprekar(i) for i in range(1000, 10000)]))
