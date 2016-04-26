from deque import Deque

def palChecker(aString):
    # Create new deque
    deque = Deque()

    # add to deque
    for ch in aString:
        deque.addRear(ch)

    stillEqual = True
    # check
    while deque.size() > 1 and stillEqual:
        first = deque.removeFront()
        last = deque.removeRear()
        if first != last:
            stillEqual = False

    return stillEqual

print(palChecker("radar"))