"""
Abstraction: 
Hides implementation details and enforces a mandatory interface for subclasses.
Ex:= abc module (ABC, @abstractmethod)

Polymorphism:
Allows different classes to respond to the same method call in their own specific way.
Ex:= Method overriding & Duck Typing

"""

from abc import ABC,abstractmethod
class PaymentProcessor(ABC):
    """
    Abstraction prevents developers from instantiating generic base 
    classes and enforces that child classes implement required methods.
    """
    @abstractmethod
    def process_payment(self,amount:float)->bool:
        """Abstract method must be implemented by all concrete subclass"""
        pass 

    @abstractmethod
    def refund(self,amount:float)->bool:
        """Abstract method: Must be implemented by all concrete subclass"""

    def log_transaction(self,amount:float)->None:
        """Concrete method: Inherited directly by all subclasses."""
        print(f"[AUDIT LOG]: Transaction of ${amount:.2f} logged.")

try:
    processor = PaymentProcessor()
except TypeError as e:
    print("Cannot instantiate processor: ",e)



class CreditCardProcessor(PaymentProcessor):
    """
    CreditCard processor
    """

    def __init__(self,card_num:str):
        self.card_num = card_num

    def process_payment(self,amount:float)->bool:
        print(f"Charging ${amount:.2f} to credit card ending in : {self.card_num[-4]}")
        return True 

    def refund(self,amount:float)->bool:
        print(f"Refunding ${amount:.2f} to credit card ending in {self.card_num[-4]}")
        return True 

class PayPalProcessor(PaymentProcessor):
    """
    PayPal processor
    """

    def __init__(self,email:str):
        self.email =email 

    def process_payment(self, amount)->bool:
        print("processing for paypal")
        return super().process_payment(amount)

    def refund(self,amount:float)->bool:
        print(f"Refunding amount: ${amount} to email user: {self.email}")
        return True 

def checkout(processor: PaymentProcessor,amount:float)->None:
    processor.log_transaction(amount)
    if processor.process_payment(amount):
        print("Payment Successful!\n")

card_proc = CreditCardProcessor("4111")
paypal_proc = PayPalProcessor("user@example.com")

checkout(card_proc,150.00)
checkout(paypal_proc,12.00)