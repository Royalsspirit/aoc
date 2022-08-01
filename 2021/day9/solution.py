from collections import deque
from functools import reduce
import heapq

def mark_basin(start, heightmap):
    dirs = [[0, 1], [0, -1], [1, 0], [-1, 0]]
    queue = deque()
    queue.append(start)
    size = 0
    while queue:
        row, col = queue.popleft()
        if heightmap[row][col] != -1:
            size += 1
            heightmap[row][col] = -1
            for rc, cc in dirs:
                new_row = row + rc
                new_col = col + cc
                if 0 <= new_row < len(heightmap) and 0 <= new_col < len(heightmap[0]) and heightmap[new_row][new_col] != -1 and heightmap[new_row][new_col] != 9:
                    queue.append((new_row, new_col))
    # print("size",size)
    return size

def smoke_basin(heightmap):
    dirs = [[0, 1], [0, -1], [1, 0], [-1, 0]]
    rows = len(heightmap)
    cols = len(heightmap[0])
    risk_sum = 0
    sizes = [float('-inf'), float('-inf'), float('-inf')]
    heapq.heapify(sizes)
    for row in range(rows):
        for col in range(cols):
            
            # Part 1
            all_less = True
            for rc, cc in dirs:
                new_row = row + rc
                new_col = col + cc
                if 0 <= new_row < rows and 0 <= new_col < cols and heightmap[row][col] >= heightmap[new_row][new_col]:
                    all_less = False
            if all_less:
                print("row",row,"col",col)
                risk_sum += heightmap[row][col] + 1

            # Part 2
            if heightmap[row][col] != -1 and heightmap[row][col] != 9:
                heapq.heappushpop(sizes, mark_basin((row, col), heightmap))
 
    return reduce((lambda x, y: x * y), sizes)

if __name__ == "__main__":
    heightmap = []
    with open('input-ex.txt', 'r') as input:
        for line in input:
            heightmap.append([int(char) for char in line.rstrip()])
    
    print(smoke_basin(heightmap))
