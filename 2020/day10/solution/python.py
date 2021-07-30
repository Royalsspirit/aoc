arr = [int(line.rstrip()) for line in    
       open('input.txt', 'r').readlines()]
arr.sort()
arr.append(arr[-1]+3)

memo = {0: 1}
for r in arr:
  print("r",r)
  memo[r] = memo.get(r-3,0) + memo.get(r-2,0) + memo.get(r-1,0)
print(memo[arr[-1]])
