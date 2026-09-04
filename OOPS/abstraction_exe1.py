from abc import ABC,abstractmethod 

class Shape(ABC):
    def __init__(self):
        pass 
    @abstractmethod
    def area(self,**kwargs)->float:
        print("Printing area for the shape")
        pass 

    @abstractmethod
    def perimeter(self,**kwargs)->float:
        print("Printing perimeter for the shape")
        pass 

    def log_info(self) -> None:
        """Concrete method: Inherited by all subclasses."""
        print(f"[{self.__class__.__name__} Info] Area: {self.area():.2f} | Perimeter: {self.perimeter():.2f}")

class Rectangle(Shape):
    def __init__(self,length:float,breadth:float,**kwargs):
        self.length = length
        self.breadth = breadth 

    def area(self)->float:
        print("Printing area for Rectangle",self.length * self.breadth)

    def perimeter(self)->float:
        print("Printing perimeter for Rectangle", 2 * (self.length + self.breadth))

    def log_info(self, **kwargs):
        print(f"Rectangle: length: {self.length}, and breadth: {self.breadth}")

class Circle(Shape):
    def __init__(self,rad:float):
        self.rad = rad 

    def area(self)->float:
        print("Printing area for circle", 3.14 * self.rad * self.rad)

    def perimeter(self)->float:
            print("Printing perimeter for circle", 2 * 3.14  * self.rad)
    
    def log_info(self, **kwargs):
        print(f"Circle: Radis: {self.rad}")

def print_shape_info(shape: Shape) -> None:
    print(f"Type: {shape.__class__.__name__}")
    print(f"Area: {shape.area():.2f}")
    print(f"Perimeter: {shape.perimeter():.2f}")
    shape.log_info()
    print("-" * 35)

if __name__ == "__main__":
    try:
        shape = Shape()
    except TypeError as e:
        print("cannot instantiate class ",e)

    rect = Rectangle(length=10,breadth=20)
    rect.area()
    rect.perimeter()
    rect.log_info()

    cir = Circle(rad=7)
    cir.area()
    cir.perimeter()
    cir.log_info()

    shapes: list[Shape] = [
        Rectangle(length=10.0, breadth=20.0),
        Circle(rad=7.0)
    ]

    for s in shapes:
        print_shape_info(s)
