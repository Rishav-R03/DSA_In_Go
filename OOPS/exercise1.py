class BankAccount:
    def __init__(self,account_holder: str, balance: float = 0.0):
        self.account_holder = account_holder 
        self.__balance = balance

    @property
    def balance(self)->float:
        """Read only property"""
        return self.__balance 
    @balance.setter 
    def balance(self,value: float):
        if value < 0:
            raise ValueError("Balance cannot be negative")
        self.__balance = value 

    def deposit(self,amount: float):
        if amount <= 0:
            raise ValueError("Deposit cannot be negative or zero")
        self.__balance += amount 

    def withdraw(self,amount: float):
        if amount <= 0:
            raise ValueError("Withdrawal cannot be negative or zero")
        if amount > self.__balance:
            raise ValueError("Insufficient funds")
        self.__balance -= amount 

if __name__ == "__main__":
    name = input("Enter account holder name:")
    acc = BankAccount(account_holder=name)
    print(f"Account Created for {acc.account_holder} with initial balance: ${acc.balance:.2f}")

    while True:
        action = input("Enter action (balance, deposit, withdraw, exit): ").lower()
        if action == "balance":
            print(f"Current balance: ${acc.balance:.2f}")
        elif action == "deposit":
            amount = float(input("Enter deposit amount: "))
            acc.deposit(amount)
            print(f"Deposited ${amount:.2f}. New balance: ${acc.balance:.2f}")
        elif action == "withdraw":
            amount = float(input("Enter withdrawal amount: "))
            acc.withdraw(amount)
            print(f"Withdrew ${amount:.2f}. New balance: ${acc.balance:.2f}")
        elif action == "exit":
            print("Exiting...")
            break
        else:
            print("Invalid action. Please try again.")