def anagramSolution1(s1, s2):
    alist = list(s2)

    pos1 = 0
    stillOk = True

    while pos1 < len(s1) and stillOk:
        pos2 = 0
        found = False
        while pos2 < len(alist) and not found:
            if s1[pos1] == alist[pos2]:
                found = True
            else:
                pos2 = pos2 + 1

        if found:
            alist[pos2] = None
        else:
            stillOk = False

        pos1 = pos1 + 1

    return stillOk

print(anagramSolution1('abcd', 'bcda'))

