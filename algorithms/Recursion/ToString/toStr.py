def toStr(number, base):
    convertString = "0123456789ABCDEF"
    if number < base:
        return convertString[number]
    else:
        return toStr(number // base, base) + convertString[number % base]

print(toStr(1453, 16))