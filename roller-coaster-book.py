from threading import Thread, Semaphore, Lock
import time
import os

C = 10
n = 15

mutex = Lock()
mutex2 = Lock()
boarders = 0
unboarders = 0
boardQueue = Semaphore(0)
unboardQueue = Semaphore(0)
allAboard = Semaphore(0)
allAshore = Semaphore(0)
total_rides = 0

def signal(sem, n):
    for _ in range(n):
        sem.release()

class Train(Thread):
    def run(self):
        while True:
            global total_rides
            print("Train loading")
            signal(boardQueue, C)
            allAboard.acquire()

            print("Train leaving")
            time.sleep(1)
            
            print("Train unboarding")
            signal(unboardQueue, C)
            total_rides+=C
            print(total_rides)
            if total_rides >= 1000:
                os._exit(1)
            allAshore.acquire()              
        
class Passenger(Thread):
    def run(self):
        global boarders, unboarders
        boardQueue.acquire()

        mutex.acquire()
        boarders += 1
        if boarders == C :
            allAboard.release()
            boarders = 0
        mutex.release()

        unboardQueue.acquire()
        
        mutex2.acquire()
        unboarders += 1
        if unboarders == C :
            allAshore.release()
            unboarders = 0
        mutex2.release()

train = Train()
train.start()
while True:
    passenger = Passenger()
    passenger.start()