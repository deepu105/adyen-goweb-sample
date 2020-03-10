## Adyen integration for a Golang application

### Sign up

**notes:**

1. Missing story when new to payments domainI like the sample projects offered by by Stripe, something like that could be helpful for people who are new to payments to get an idea using a working example app
2. what is difference between Checkout APIs and payment APIs, they seem to have some shared functionality but API explorer doesn't talk about any differences, Is payments legacy API? This part is bit confusing when deciding what to choose
3. In the guides for the checkout API, there is no direct instruction on where to get the SDK/Libs from, its mentioned but it would be nice to just provide that info as first step to make it easy for the devs like how stripe does it
4. The merchant account to use is not clear enough, since there were two accounts and only one worked for the API
5. In "Step 2: Add Components to your payments form" I find the flow would be easier if example is followed by explanation
6. Would be nice to give example of using reference for mounting Components in Vue/React
7. The redirect behaviors from `checkout.createFromAction` is not very clear for 3DS2 when POST redirect is involved, this is where I spent the most time, so might be nice to provide some simple samples here else you would have to dig up https://github.com/adyen-examples/ and go through it
8. The amount gets messed up during redirection when there are decimal points
9. The country dropdown in the billing address of component doesn't work with standard list filtering
10. The state field in the billing address of component resets when selecting country

### Website

1. getting started: seems bit involved as it requires you to get advice from Adyen team - for mid market it would be nicer if this directs more towards self service
2. payment fundamentals:
   1. payments glossary: terms are in alphabetical order which means you jump around a lot if you are new to the domain(which could be the case for mid market), might be nicer to explain in logical sequence with an example
3.

### Positives

1. Lot of documentation
2. API explorer
