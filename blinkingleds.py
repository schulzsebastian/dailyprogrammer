import re

'''
for n in re.findall(r'ld a,(\d+)', open('leds.txt', 'r').read()):
    leds = bin(int(n)).split('b')[1].replace('1', '*').replace('0', '.')
    while len(leds) < 8:
        leds = '.' + leds
    print leds
'''
output = []
for line in open('leds.txt').readlines():
    if line.startswith('ld'):
        output.append(bin(
            int(line.strip().split(',')[1])
            ).split('b')[1].replace('1', '*').replace('0', '.').rjust(8, '.'))
    elif not line.strip():
        output.append(output[len(output) - 1])
print "\n".join(output)