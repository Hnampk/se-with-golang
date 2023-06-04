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

## A popular Go proverb to avoid the extra import that triggers a circular dependency
> "A little copying is better than a little dependency."

## Tips and tools for writing lean and easy-tomaintain Go code
1. Readability & performance.
- If you find yourself navigating through a lengthy function that contains deeply nested if/else blocks, repeated blocks of code, or its implementation tackles several seemingly unrelated concerns, it would be a great opportunity to apply some drive-by refactoring and extract any potential self-contained blocks of logic into separate functions.
- There is always a trade-off between code readability and performance. Sometimes, splitting the business logic across functions takes a toll on performance: keep the implementation neatly tucked within a single function to avoid the extra Go runtime overhead that's incurred when making function calls (for example, pushing arguments to the stack, checking that the stack is large enough for the callee, and popping things off the stack when the function call returns).
> "There is always a trade-off between code readability and performance.
When dealing with complex systems, readability is oftentimes preferred,
but at the end of the day, it's up to you and your team to figure out which
mix of readability versus performance works best for your particular use
case."

2. Variable naming conventions

put some love while naming your variables

    for dateIdx := 0; dateIdx < len(tickers); dateIdx++ {
        for stockIdx := 0; stockIdx < len(tickers[dateIdx]); stockIdx++ {
            value := tickers[dateIdx][stockIdx]
            // ...
        }
    }

instead of

    for i := 0; i < len(s); i++ {
        for j := 0; j < len(s[i]); j++ {
            value := s[i][j]
            // ...
        }
    } // other developers/or your future self might wonder what is s? What is the value we're trying to access?

> "At the end of the day, each engineer has their own preferred variable
naming approach and philosophy. When deciding on which approach to
adopt, try to take a few minutes to consider how your variable naming
choices affect other engineers that collaborate with you on shared code
bases."

3. Using Go interfaces effectively
