class Car:
    color: str 
    model: str 
    def __init__(self,color:str,model:str,engine_state:bool=False):
        self.color = color 
        self.model = model 
        self.engine_state = engine_state

    def start_engine(self):
        if self.engine_state:
            print("Engine is already started")
        else:
            self.engine_state = True
            print("Engine started")
    def stop_engine(self):
        if(self.engine_state):
            self.engine_state = False
            print("Engine stopped")
        else:
            print("Engine is already stopped") 
car1 = Car("White","Hyundai")
car1.start_engine()
car1.stop_engine()
car1.stop_engine()

print(car1.color)
print(car1.model)

class Stack:
    def __init__(self,n:int):
        self.stack = [0] * n  
        self.top = -1
    def push(self,element:int)->bool:
        if self.top == len(self.stack)-1:
            print("Stack Overflow")
            return False
    
        self.top +=1
        self.stack[self.top] = element 
        return True 

    def pop(self)->int:
        if self.top == -1:
            print("Stack underflow")
            return -1 
        self.top -=1
        return self.stack[self.top+1]

    def display(self):
        if self.top == -1:
            print("stack is empty")
            return 
        for i in range(self.top,-1,-1):
            print(self.stack[i],end=" ")

if __name__ == "__main__":
    s = Stack(5)
    s.push(1)
    s.push(2)
    s.push(3)   
    s.display()
    popped = s.pop()
    print("popped element:" + str(popped))
    s.display()

