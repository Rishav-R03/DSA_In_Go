"""
Notification class for the exercise 3 of the OOPS module.
"""

class ConsoleNotifier:
    """
    A base class for console notification.    
    """
    def __init__(self,**kwargs):
        super().__init__() # calls object.__init__() safely
    def send(self,message:str)->None:
        print("[Console]: ",message)

class EmailDecorator(ConsoleNotifier):
    """
    Email notifier
    """
    def __init__(self,email:str,**kwargs):
        super().__init__(**kwargs)
        self.email = email 
    def send(self,message:str)->None:
        print(f"[Email]: to {self.email}")
        super().send(message)


class SMSNotifier(ConsoleNotifier):
    """
    SMS Notifier
    """
    def __init__(self,phone:int,**kwargs):
        super().__init__(**kwargs)
        self.phone = phone 
    def send(self,message:str)->None:
        print(f"[SMS] message: {message} sent to phone: {self.phone}")

class MultiChannel(EmailDecorator,SMSNotifier):
    def __init__(self,email:str,phone:int,**kwargs):
        super().__init__(email=email,phone=phone,**kwargs)
        self.email = email 
        self.phone = phone 
    def send(self,message:str)->None:
        print(f"Sending message to phone and email")
        print(f"[Multi-Channel] message: {message} SMS to: {self.phone}, and email to: {self.email}") 


if __name__ == "__main__":
    console = ConsoleNotifier()
    console.send("a message for console")

    sms = SMSNotifier(8234)
    sms.send("a message for sms")

    email = EmailDecorator("rishav@example.com")
    email.send("a message for email")

    muti_chan = MultiChannel(email="rishav@example.com",phone=8234)
    muti_chan.send("a message for multi-channel")