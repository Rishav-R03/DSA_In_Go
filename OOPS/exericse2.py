class Employee:
    def __init__(self, name: str, base_salary: float):
        self.name = name 
        self.base_salary = base_salary

    def calculate_pay(self, bonus: float = 0.0) -> float:
        return self.base_salary + bonus


class Manager(Employee):
    def __init__(self, name: str, base_salary: float, department: str, bonus: float = 0.0):
        super().__init__(name, base_salary)
        self.department = department
        self.bonus = bonus  # Flat monetary bonus attribute

    def calculate_pay(self) -> float:
        # Calls parent calculate_pay using self.bonus
        return super().calculate_pay(self.bonus)


if __name__ == "__main__":
    emp1 = Employee(name="Alice", base_salary=50000.0)
    emp2 = Manager(name="Bob", base_salary=70000.0, department="Sales", bonus=15000.0)

    print(f"{emp1.name} Pay: ${emp1.calculate_pay(bonus=2000.0):.2f}")  # $52,000.00
    print(f"{emp2.name} ({emp2.department}) Pay: ${emp2.calculate_pay():.2f}")  # $85,000.00