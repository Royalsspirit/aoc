with open("input.txt") as fin:
    data = fin.read()

program = data.split("\n")
print(len(program))
def run(program):
    accumulator = 0
    pointer = 0

    def execute(instruction):
        if len(instruction) == 0:
            return 0, 0
        op, arg = instruction.split(" ")
        if op == "acc":
            return int(arg), 1

        elif op == "jmp":
            return 0, int(arg)

        else:
            return 0, 1
    
    
    seen = set() # Set of pointers we've seen so far
    accumulator = 0
    pointer = 0
    while pointer not in seen and pointer < len(program):
        seen.add(pointer)
        print(program[pointer])
        acc, jmp = execute(program[pointer])
        accumulator += acc
        print("accumulator",accumulator)
        pointer += jmp

    if pointer >= len(program):
        return accumulator

    return None


# BASH BASH BASH
for i in range(len(program)):
    print("next",program[i])
    if program[i].startswith("acc"):
        continue

    if program[i].startswith("nop"):
        copy = program.copy()
        copy[i] = "jmp" + copy[i][3:]

    else:
        copy = program.copy()
        copy[i] = "nop" + copy[i][3:]
    
    x = run(copy)
    if x:
        print(x)
        break
