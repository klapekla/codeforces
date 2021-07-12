import re

players = input()

if re.search('0{7}|1{7}', players):
    print('YES')
else:
    print('NO')