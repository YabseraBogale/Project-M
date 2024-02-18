a=1
b=2
for i in range(1,60):
	b=a+b
	a=b-a
print(len(str(b)))

