"""
Public self.val
Protected self._val
Private self.__val 
"""

class BankAccount:
    owner: str 
    balance: float
    pin: int 
    accType: str 
    def __init__(self,owner:str, balance: float,pin: int):
        self.owner = owner
        self._accType = "Savings"
        self.__pin = pin
    def display(self):
        print(f"Owner: {self.owner}, Account Type: {self._accType}, Pin: {self.__pin}")

acc1 = BankAccount("John",1000,2345)
print(acc1._accType)

try:
    print(acc1.__pin)
except AttributeError as e:
    print("direct access failed:" ,e)
print(acc1._BankAccount__pin)


"""
    Controlled Access
"""

class SmartAccount:
    def __init__(self,balance: int):
        self.__balance = balance 
    @property
    def balance(self)->int:
        """Read-Only access to balance"""
        return self.__balance        

    @balance.setter 
    def balance(self,value: int):
        """Validates input before updating internal state"""
        if value < 0:
            raise ValueError("Balance cannot be negative")
        self.__balance = value

    @balance.deleter
    def balance(self) -> None:
        """Deletes the balance attribute"""
        print("Resetting balance to 0")
        self.__balance = 0

acc = SmartAccount(1000)
print(acc.balance)

acc.balance = 2000

try:
    acc.balance = -1
except ValueError as e:
    print("Error: ",e)