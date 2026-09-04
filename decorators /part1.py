"""
Decorator methods alter how class methods interact with class and instance state
Types:
    1. Instance Methods 
    2. Class Methods 
    3. Static Methods 
"""

"""
Class Methods @classmethod 
    Receives class itself (cls) as it's first argument instead of an instance (self)
Use case:
    1. Alternative constructors (Factory Methods): Instantiating objects using alternative 
    data formats like json strings, dates or databases.
    2. Class-Level State Management: Accessing or Modifying attributes shared across all instances.
"""

class User:
    def __init__(self,username:str,email:str):
        self.username = username 
        self.email = email 
    
    # Alternative Constructor 1: Parse from CSV string
    @classmethod 
    def from_csv(cls,csv_string:str)-> "User":
        username,email = csv_string.split(",")
        return cls(username.strip(),email.strip())

    # Alternative constructor 2: Parse from Dictionary 
    @classmethod  
    def from_dict(cls,data:dict) -> "User":
        return cls(username=data["username"],email=data["email"])

#standard tarika
user1 = User("alice","alice@example.com")

# Instantiation using factory method @classmethods 
user2 = User.from_csv("bob_dev, bobdev@exampl.com")
user3 = User.from_dict({"username":"bob_dev","email":"bobdev@example.com"})

print(user2.username,user3.username)

"""
Static Methods:
    A @staticmethod is a plain function bound inside a class namespace. 
    It receives neither self nor cls directly or automatically.

Primary Usecase:
    1. Utility or Helper function:
        Pure logic related to the domain of the class that does not read or mutate instance of class.
    2. Code organization: Grouping standalone helper logic inside the class where it logically belongs.
"""

class PassswordValidator:
    MIN_LENGTH = 8 
    def __init__(self,password:str):
        self.password = password 

    @staticmethod 
    def is_valid_email(email:str)->bool:
        return "@" in email and "." in email.split("@")[-1]

    @staticmethod 
    def calculate_entropy(text:str) -> float:
        import math 
        return len(text) * math.log2(len(set(text))) if text else 0.0

print(PassswordValidator.is_valid_email("test@com"))
print(PassswordValidator.calculate_entropy("P@ssword123"))



"""
Task 1: 
"""

class Date:
    def __init__(self,day:int,month:int,year:int):
        self.day = day 
        self.month = month 
        self.year = year 

    @classmethod
    def from_string(cls,date:str) ->list[int]:
        day,month,year = date.split(":")
        return cls(day=day,month=month,year=year)

    @staticmethod 
    def is_leap_year(year:int) ->bool:
        return (year%4==0 and year % 100 != 0) or (year % 400 == 0) 

date1 = Date.from_string("04:08:2003")
print(date1.day)
print(Date.is_leap_year(2028))