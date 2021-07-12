n, k = list(map(int,input().split(' ')))
scores = list(map(int,input().split(' ')))
min_score_to_advance = scores[k-1]
print(len(list(filter(lambda score: score > 0 and score >= min_score_to_advance, scores))))
