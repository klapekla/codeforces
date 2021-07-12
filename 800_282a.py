n = int(input())
x = 0
for i in range(n):
  statement = input()
  x = x+1 if '++' in statement else x-1
print(x)
