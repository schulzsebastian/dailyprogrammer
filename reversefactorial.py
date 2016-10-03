for line in open('reversefactorial_input.txt', 'r').readlines():
    x = int(line.strip())
    i = x / x
    while x > 1:
        x /= i
        i += 1
        if x == 1:
            print '{} = {}!'.format(line.strip(), i - 1)
            break
    else:
        print '{} NONE'.format(line.strip())
