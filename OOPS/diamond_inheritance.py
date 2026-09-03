class Walker:
    def __init__(self,walk_speed: float, **kwargs):
        super().__init__(**kwargs)
        self.walk_speed = walk_speed

    def move(self)->str:
        return f"Walking at {self.walk_speed} m/s"


class Swimmer:
    def __init__(self,swim_speed: float, **kwargs):
        super().__init__(**kwargs)
        self.swim_speed = swim_speed 

    def move(self)->str:
        return f"Swimming at {self.swim_speed} m/s"

class Amphibian(Walker, Swimmer):
    def __init__(self,name:str,**kwargs):
        super().__init__(**kwargs)
        self.name = name 

frog = Amphibian(name="Froggy", walk_speed=1.5, swim_speed=2.0)

print(f"Name: {frog.name}")
print(f"Walk Speed: {frog.walk_speed} m/s")
print(f"Swim Speed: {frog.swim_speed} m/s")

print([c.__name__ for c in Amphibian.__mro__])
