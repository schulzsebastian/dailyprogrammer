import itertools
def is_consecutive(l, i=None):
    for e in l:
        if not i:
            i = e
        else:
            if i - e in [1, -1]:
                return True
            i = e
    return False
for n in open('fibonacciintegers_input.txt', 'r').readlines()[1:]:
    l, o, x = [], [], []
    a, b = 1, 1
    while a <= int(n):
        a, b = b, a + b
        if a <= int(n):
            l.append(a)
    for i in range(len(l)+1):
        for c in list(itertools.combinations(list(reversed(l)), i)):
            if sum(c) == int(n) and not is_consecutive(c):
                o.append(c)
    print n.strip() + ' = ' + ' + '.join([str(i) for i in o[0]])
