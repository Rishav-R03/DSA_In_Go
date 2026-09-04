class Solution:
    def __init__(self,arr:list[int]):
        self.arr = arr 
    def contains_duplicate(self,arr:list[int])->bool:
        map1 = {}
        for a in arr:
            if a in map1:
                map1[a] +=1
            else :
                map1[a] = 1

        for _,v in map1.items():
            if v > 1:
                return True 
        return False

    def two_sum(self,arr:list[int],target:int)->list[int]:
        map1 = {}
        for k,v in enumerate(arr):
            diff = target-v 
            if diff in map1:
                return [map1[diff],k]
            map1[diff] = k 
        return []


if __name__ == "__main__":
    arr = [1,3,1]
    if(Solution(arr).contains_duplicate(arr)):
        print("contains duplicate")
    else:
        print("no it doesnot")

    indices = Solution(arr).two_sum(arr,2)
    print(indices)