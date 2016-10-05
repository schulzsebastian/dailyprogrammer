import itertools
for n in open('fibonacciintegers_input.txt', 'r').readlines()[1:]:
    l, o = [], []
    a, b = 1, 1
    while a < int(n):
        a, b = b, a + b
        if a < int(n):
            l.append(a)
    for i in range(len(l)+1):
        for c in list(itertools.combinations(list(reversed(l)), i)):
            if sum(c) == int(n):
                o.append(n.strip() + ' = ' + ' + '.join([str(i) for i in c]))
    print o[0]
