def anagramSolution3(s1, s2):
    c1 = [0]*26
    c2 = [0]*26

    for i in range(len(s1)):
        pos = ord(s1[i]) - ord('a')
        c1[pos] = c1[pos] + 1

    for i in range(len(s2)):
        pos = ord(s2[i]) - ord('a')
        c2[pos] = c2[pos] + 1

    j = 0;
    ok = True
    while j < 26 and ok:
        if (c1[j] == c2[j]):
            j += 1
        else:
            ok = False

    return ok

print(anagramSolution3('abcd', 'bcda'))