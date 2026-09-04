"""
Python allows custom classes to intercept built-in operations (like +, ==, <, len(), or []) by implementing special methods
(also called dunder methods for double underscore).

By default, Python compares objects by memory location (id()). 
Implementing __eq__ and __lt__ allows value-based comparisons.
"""

class Product:
    def __init__(self,name:str,price:str):
        self.name = name 
        self.price = price 

    # equality check
    def __eq__(self,other:object) -> bool:
        if not isinstance(other,Product):
            return NotImplemented
        return self.price == other.price

    # less than check 

    def __lt__(self,other:object) -> bool:
        if not isinstance(other,Product):
            return NotImplemented
        return self.price < other.price 

p1 = Product("Acer laptop",80000)
p2 = Product("Dell laptop",90000)

print(p1<p2)


class Inventory:
    def __init__(self,items:list[str]):
        self._items = items 

    #control len 
    def __len__(self)->int:
        return len(self._items)

    def __getitem__(self,index:int)->str:
        return self._items[index]

inv = Inventory(["CPU","GPU","TPU","RAM"])

print(len(inv))
print(inv[1])