from threading import Thread, Semaphore, Lock

buffer = []
buffer_size = 10
filled = Semaphore(0)
empty = Semaphore(buffer_size)
lock = Lock()

class Producer(Thread):
    def run(self):
        num = 0
        while True:
            empty.acquire()
            lock.acquire()
            buffer.append(num)
            lock.release()
            num = num + 1
            filled.release()


class Consumer(Thread):
    def run(self):
        num = 0
        while True:
            filled.acquire()
            lock.acquire()
            num = buffer.pop(0)
            lock.release()
            print("Consumed %i" % num)
            empty.release()


prod = Producer()
cons = Consumer()
prod.start()
cons.start()
    