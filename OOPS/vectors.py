import math


class Vector:
    def __init__(self,x :float, y: float):
        self.x = x
        self.y = y 
    def __eq__(self,other:Vector)->bool:
        return (self.x == other.x) and (self.y == other.y)
    def __add__(self,other:Vector) -> Vector:
        res = Vector(0,0)
        res.x = self.x + other.x
        res.y = self.y + other.y 
        return res 

    def __len__(self)->int:
        return int(math.sqrt(self.x**2 + self.y**2))

    def __repr__(vec:Vector)->str:
        return f"x: {vec.x} and y {vec.y}"


v1 = Vector(1,2)
v2 = Vector(2,4)
print(v1 == v2)
l2 = len(v2)
print(l2)
add = v1 + v2
print(add)