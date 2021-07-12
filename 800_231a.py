n = int(input())
solvable_solutions = 0
for i in range(n):
  if sum(map(int,input().split(' '))) >= 2:
    solvable_solutions += 1  
print(solvable_solutions)

