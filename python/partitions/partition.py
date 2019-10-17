def partition(n, I=1):
    yield (n,)
    for i in range(I, n//2 + 1):
        for p in partition(n-i, i):
            yield (i,) + p

def partitions(n):
  a = partition(n)
  c = 0
  for i in a:
    c += 1
  return c

a = partitions(100)
print(a)