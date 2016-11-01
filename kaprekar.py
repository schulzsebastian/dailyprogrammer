class Kaprekar:
    def __init__(self, start, end):
        self.start = start
        self.end = end
        print(self.find_kaprekars())

    def is_kaprekar(self, n):
        sqrt = n ** 2
        str_sqrt = str(sqrt)
        if sqrt == n:
            return True
        for i in range(1, len(str_sqrt)):
            f, s = int(str_sqrt[:i]), int(str_sqrt[i:])
            if f == 0 or s == 0:
                continue
            if f + s == n:
                return True

    def find_kaprekars(self):
        kaprekars = []
        for i in range(self.start, self.end):
            if self.is_kaprekar(i):
                kaprekars.append(str(i))
        return " ".join(kaprekars)

Kaprekar(1, 50)
Kaprekar(2, 100)
Kaprekar(101, 9000)
