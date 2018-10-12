from threading import Thread, Semaphore, Lock

savages = 12
M = 10
servings = 0
filled = Semaphore(0)
empty = Semaphore(0)
lock = Lock()

class Cook(Thread):
    def run(self):
        while True:
            empty.acquire()
            print("Filling pot")
            filled.release()


class Savage(Thread):
    def __init__(self, i):
        self.i = i
        super(Savage, self).__init__()
        print("Savage " + str(i) + "made")
    
    def run(self):
        global servings
        while True:
            lock.acquire()
            if servings == 0:
                empty.release()
                filled.acquire()
                servings = M
            servings -= 1
            print("Savage " + str(i) + " has ate")
            lock.release()


cook = Cook()
cook.start()
for i in range(savages):
    sav = Savage(i)
    sav.start()

cook.join()
    