"""
Key Inheritance      | PatternsPatternDescription                                  | Syntax / Behavior
Single Inheritance   | A child class inherits from one parent class.               | class Child(Parent):
Multiple Inheritance | A child class inherits from two or more parent classes      | .classChild(ParentA, ParentB):
Constructor Chaining | Initializes parent attributes inside the child constructor. | super().__init__(args)
"""

class Animal:
    def __init__(self,name: str,species: str):
        self.name = name 
        self.species = species 
    def make_sound(self):
        return "A sound"

class Dog(Animal):
    def __init__(self,name:str,breed: str):
        super().__init__(name=name,species="Dog")
        self.breed = breed 

    def make_sound(self):
        return "Woof"

    def display(self):
        print(f"{self.name} is a {self.species} of breed {self.breed} and it says {self.make_sound()}")

dog = Dog("Moku","Golden Retriever")
dog.display()

"""
Python allows a class to derive from multiple parent classes. Method Resolution Order (MRO) determines the order in which parent classes are searched for methods and attributes.
"""

class Walker:
    def __init__(self, walk_speed: float):
        self.walk_speed = walk_speed

    def move(self) -> str:
        return f"Walking at {self.walk_speed} m/s"

class Swimmer:
    def __init__(self, swim_speed: float):  # Fixed __int__ -> __init__
        self.swim_speed = swim_speed 

    def move(self) -> str:
        return f"Swimming at {self.swim_speed} m/s"

class Amphibian(Walker, Swimmer):
    def __init__(self, name: str, walk_speed: float, swim_speed: float):
        Walker.__init__(self, walk_speed)
        Swimmer.__init__(self, swim_speed)
        self.name = name

frog = Amphibian("Froggy", 1.5, 2.0)

# Calling specific parent methods explicitly if both actions are needed
print(f"{frog.name} can {Walker.move(frog)} and {Swimmer.move(frog)}")

# Default move() resolves to Walker via MRO
print(frog.move())

# MRO Output
print([c.__name__ for c in Amphibian.__mro__])