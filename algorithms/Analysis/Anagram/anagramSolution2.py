def anagramSolution2(s1, s2):
    alist1 = list(s1)
    alist2 = list(s2)

    alist1.sort()
    alist2.sort()

    pos = 0
    ok = True

    while pos < len(s1) and ok:
        if alist1[pos] == alist2[pos]:
            pos += 1
        else:
            ok = False

    return ok

print(anagramSolution2("abcd", "bcda"))