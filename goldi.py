
with open('goldi.txt') as f:
    w, t = [int(i) for i in f.readline().split()]
    print(" ".join([str(seat + 1) for seat, line in enumerate(f.readlines()) if int(line.split()[0]) >= w and int(line.split()[1]) <= t]))

print(" ".join([str(seat + 1) for seat, line in enumerate(open('goldi.txt').readlines()[1:]) if int(line.split()[0]) >= int(open('goldi.txt').readline().split()[0]) and int(line.split()[1]) <= int(open('goldi.txt').readline().split()[1])]))

(lambda f, data: f(f, data.splitlines()))((lambda this_f, data: print(" ".join([str(s + 1) for s, l in enumerate(data[1:]) if int(l.split()[0]) >= int(data[0].split()[0]) and int(l.split()[1]) <= int(data[0].split()[1])]))), open('goldi.txt').read())
