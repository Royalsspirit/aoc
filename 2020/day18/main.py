import time
import re
start=time.time()
a=open('input.txt').read().split("\n")[:-1]
exp=[]
for x in a:
    expr=[]
    for l in x:
        if l!=' ':
            expr+= [l]
    exp+=[expr]

def brace(e):
    br=0
    for i in range(len(e)):
        if e[i]=='(': br+=1
        elif e[i]==')': br-=1
        if br==0:
            return calc2(e[1:i])[0],i+1

def calc2(e):
    if e[0]!='(': r=int(e[0])
    else: r=0
    i=0
    while i <len(e)-1:
        if e[i]=='(':
            a,b=brace(e)
            r=a
            i+=b
        elif e[i]=='+':
            if e[i+1]!='(':
                r+=int(e[i+1])
                i+=1
            else:
                a,b=brace(e[i+1:])
                r+=a
                i+=b+1
        elif e[i]=='*':
                a,b=calc2(e[i+1:])
                r*=a
                i+=b+1
        elif e[i]==')':
            i+=1
            return r,i
        else: i+=1
    return r,i

def calc(e):
    if e[0]!='(': r=int(e[0])
    else: r=0
    i=0
    while i <len(e)-1:
        if e[i]=='(':
            a,b= calc(e[i+1:])
            r=a
            i+=b+1
        elif e[i]=='+':
            if e[i+1]!='(':
                #print('here')
                r+=int(e[i+1])
                i+=1
            else:
                a,b=calc(e[i+2:])
                r+=a
                i+=b+2
        elif e[i]=='*':
            if e[i+1]!='(':
                r*=int(e[i+1])
                i+=1
            else:
                a,b=calc(e[i+2:])
                r*=a
                i+=b+2
        elif e[i]==')':
            i+=1
            return r,i
        else: i+=1
    return r,i

sum,sum2=0,0
for e in exp:
   sum+=calc(e)[0]
   print("sum",sum)
print('part1',sum)
for e in exp:
   sum2+=calc2(e)[0]
print('part2',sum2)
end=time.time()
print(round(end-start,6))
