def test(content):
    f = 0
    r = set([0])
    while True:
        for i in content:
            f += int(i)
            if f in r:
                return f
            else:
                r.add(f)


with open('input') as f:
    content = f.readlines()
print(test(content))
