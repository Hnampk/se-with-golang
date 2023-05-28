Section 2: Best Practices for Maintainable and Testable Go Code

## SOLID

1. S: The `single responsibility principle (SRP)` was described by Robert Martin - Uncle Bob
> "In any well-designed system, objects should only have a single responsibility."

2. O: The `open/closed principle` was coined by Bertrand Meyer
> "A software module should be open for extension but closed for modification."

- In other words, a package should not be able to modify the behavior of
the types that are exported by other packages

3. L: `Liskov substitution principle (LSP)` was introduced by Barbara Liskov
> "If, for each object, `O1` of type `S` there is an object `O2` of type `T` such that for all  programs `P` defined in terms of `T`, the behavior of `P` is unchanged when `O1` is substituted for `O2`, then `S` is a subtype of `T`"


4. I: the `interface segregation principle (ISP)` was also coined by Robert Martin
> "According to this principle, clients should not be forced to depend upon the interfaces that they do not use."


5. D: Yet another principle identified by Robert Martin is the `dependency inversion principle (DIP)`.
> "High-level modules should not depend on low-level modules. Both should depend on abstractions. 
Abstractions should not depend on details. Details should depend on abstractions."
