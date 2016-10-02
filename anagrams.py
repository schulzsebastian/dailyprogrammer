for line in open('anagrams_input.txt', 'r').readlines():
    x, y = [[j.lower() for j in i if j.isalpha()] for i in line.strip().split('?')]
    if sorted(x) == sorted(y): print '{} is an anagram of {}'.format(line.split('?')[0].strip(), line.split('?')[1].strip())
    else: print '{} is NOT an anagram of {}'.format(line.split('?')[0].strip(), line.split('?')[1].strip())
