import Control.Concurrent
import System.Random

fsthalf :: [a] -> [a]
fsthalf xs = take (length xs `div` 2) xs

sndhalf :: [a] -> [a]
sndhalf xs = drop (length xs `div` 2) xs

merge :: Ord a => [a] -> [a] -> [a]
merge xs [] = xs
merge [] ys = ys
merge (x:xs) (y:ys) 
          | (x <= y)  = x:(merge xs (y:ys)) 
          | otherwise = y:(merge (x:xs) ys)


mergesort :: Ord a => [a] -> MVar [a] -> IO()
mergesort [] ch = putMVar ch [] 
mergesort [x] ch = putMVar ch [x]

mergesort xs ch = do
    left <- newEmptyMVar
    right <- newEmptyMVar
    forkIO $ mergesort (fsthalf xs) left
    forkIO $ mergesort (sndhalf xs) right
    a <- takeMVar left
    b <- takeMVar right
    putMVar ch (merge a b)

main = do
    g <- getStdGen
    let list =  take 10000 (randoms g :: [Int])
    result <- newEmptyMVar
    forkIO $ mergesort list result
    sorted <- takeMVar result
    print(sorted)