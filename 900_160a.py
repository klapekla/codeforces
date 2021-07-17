number_of_coins = int(input())
coins_values = list(map(int, input().split()))
coins_values.sort(reverse=True)
for i in range(number_of_coins):
  if sum(coins_values[:i+1]) > sum(coins_values[i+1:]):
    print(i+1)
    break

