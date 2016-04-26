from queue import Queue

def hotPotato(namelist, num):
    queue = Queue()
    for name in namelist:
        queue.enqueue(name)

    while queue.size() > 1:
        for i in range(num):
            queue.enqueue(queue.dequeue())

        queue.dequeue()

    return queue.dequeue()

print(hotPotato(["A", "B", "C", "D", "E", "F"], 7))
