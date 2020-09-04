# The Farmer’s Market

From: https://gist.github.com/jbartels/d75a9f5282abebe071694723a5f25f0e

We sell the following four products at the local farmer’s market every week.

```
+--------------|--------------|---------+
| Product Code |     Name     |  Price  |
+--------------|--------------|---------+
|     CH1      |   Chai       |  $3.11  |
|     AP1      |   Apples     |  $6.00  |
|     CF1      |   Coffee     | $11.23  |
|     MK1      |   Milk       |  $4.75  |
|     OM1      |   Oatmeal    |  $3.69  |
+--------------|--------------|---------+
```

This week, we’re celebrating our one year anniversary and would like to offer the
following specials. To do so, we’ll need to update our checkout system to apply
the following rules.

1. BOGO -- Buy-One-Get-One-Free Special on Coffee. (Unlimited)
2. APPL -- If you buy 3 or more bags of Apples, the price drops to \$4.50.
3. CHMK -- Purchase a box of Chai and get milk free. (Limit 1)
4. APOM -- Purchase a bag of Oatmeal and get 50% off a bag of Apples

At any time, we should be able to print out the current register to see what the state of
the basket is. This should include the price and the applied discount or special, if any.

For example, given the following basket:

CH1, AP1, AP1, AP1, MK1

After CH1 and AP1, it should yield:

```
Item                          Price
----                          -----
CH1                            3.11
AP1                            6.00
-----------------------------------
                               9.11
```

After the entire basket is added, it would yield:

```
Item                          Price
----                          -----
CH1                            3.11
AP1                            6.00
            APPL              -1.50
AP1                            6.00
            APPL              -1.50
AP1                            6.00
            APPl              -1.50
MK1                            4.75
            CHMK              -4.75
-----------------------------------
                              16.61
```

Use Python, Ruby, Javascript, or Golang to implement a checkout system that allows us to fulfill the above requirements. Submit the code via a Dockerfile which builds a Docker container to run the code sample in. A readme should be included which describes how to execute the code. The source code needs to be publicly accessible for code review.

Here’s some test data for your specs:

```
Basket: CH1, AP1, CF1, MK1
Total price expected: $20.34
```

```
Basket: MK1, AP1
Total price expected: $10.75
```

```
Basket: CF1, CF1
Total price expected: $11.23
```

```
Basket: AP1, AP1, CH1, AP1
Total price expected: $16.61
```

---

The scope of this project is meant to be fairly narrow, but there are a lot of details to consider.
We’re looking for the following things when reviewing your code:

1. Design
2. Testing
3. Accuracy
4. Flexibility
5. Containerization

Please feel free to ask questions as they arise via email at or-devleads@objectrocket.com - do not leave comments here.
