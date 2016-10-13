    from itertools import permutations
    def solve_mathagram(mathagram, n):
        l = list(range(1, 10)) * n
        for i in list(mathagram):
            if i.isdigit():
                l.remove(int(i))
        p = permutations(l, mathagram.count('x'))
        while True:
            g = next(p)
            x = mathagram.replace('x', '{}').format(*g).split('=')
            s1, s2 = 0, 0
            for i in x[0].split('+'):
                s1 += int(i)
            for i in x[1].split('+'):
                s2 += int(i)
            if s1 == s2:
                return "=".join(x).strip()
'''
import profile
def run():
    for i in ['xxx + xxx + xxx + 4x1 + 689 = xxx + xxx + x5x + 957',
            'xxx + xxx + xxx + 64x + 581 = xxx + xxx + xx2 + 623',
            'xxx + xxx + xxx + x81 + 759 = xxx + xxx + 8xx + 462',
            'xxx + xxx + xxx + 6x3 + 299 = xxx + xxx + x8x + 423',
            'xxx + xxx + xxx + 58x + 561 = xxx + xxx + xx7 + 993']:
        solve_mathagram(i, 3)
profile.run('run()')
'''
