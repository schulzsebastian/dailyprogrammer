def make_bricks(s, b, g):
    if any(map(lambda x, y: x <= g <= y, range(5, b * 5 + 5, 5), range(5 + s, b * 5 + 5, 5))):
        return True
    return False
