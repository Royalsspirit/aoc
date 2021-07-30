def forest_walk_C(forest_rows, horizontal_increment, vertical_increment):

    tree_counter = 0
    forest_width = len(forest_rows[0].strip())
    horizontal_position = 0
    row_number = 0

    for row in forest_rows:
        if vertical_increment != 1 and row_number % vertical_increment != 0:
            row_number += 1
            continue
        row_stripped = row.strip()
        if row_stripped[horizontal_position] == '#':
            tree_counter += 1
        horizontal_position += horizontal_increment
        if horizontal_position > (forest_width - 1):
            horizontal_position -= forest_width
        row_number +=1

    return tree_counter

def day_03_C():
    file = open('input.txt', 'r')

    forest_rows = []

    for line in file:
        forest_rows.append(line)
    
    print(forest_walk_C(forest_rows, 1, 1))
    print(forest_walk_C(forest_rows, 3, 1))
    print(forest_walk_C(forest_rows, 5, 1))
    print(forest_walk_C(forest_rows, 7, 1))
    print(forest_walk_C(forest_rows, 1, 2))

    print(forest_walk_C(forest_rows, 1, 1) *
          forest_walk_C(forest_rows, 3, 1) *
          forest_walk_C(forest_rows, 5, 1) *
          forest_walk_C(forest_rows, 7, 1) *
          forest_walk_C(forest_rows, 1, 2))

day_03_C()
