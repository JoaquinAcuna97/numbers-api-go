# NUMBERS API GO
API that categorizes and stores numbers based on their divisibility


### Objective
> [!TIP]
> Implement an API that categorizes and stores numbers based on their divisibility by
3 and 5, mimicking the logic described below:
* Numbers divisible by 3: Categorized as "Type 1"
* Numbers divisible by 5: Categorized as "Type 2"
* Numbers divisible by both 3 and 5: Categorized as "Type 3"
* Other numbers: Stored as their original value

### Constraints
> [!CAUTION]
> Utilize only one if statement within the core logic. Avoid else, else if, ternary
operators, or switch statements.

### Functional Requirements
> [!NOTE]
>  The API shall store numbers and their classifications in memory or any
other support you consider appropriate.

#### Classifications include:
+ Numbers divisible by 3: Categorized as "Type 1"
+ Numbers divisible by 5: Categorized as "Type 2"
+ Numbers divisible by both 3 and 5: Categorized as "Type 3"
+ Other numbers: Stored as their original value

#### Number Addition:
○ The API shall provide a mechanism to save a number into its collection.
#### Number Retrieval:
○ The API shall allow the retrieval of a specific number's classification. If
the number is not found, the API should indicate this.
#### Collection Retrieval:
○ The API shall enable the retrieval of all stored numbers and their
classifications.