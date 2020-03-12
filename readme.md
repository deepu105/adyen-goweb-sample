## Adyen integration for a Golang application

### Positives

1. Lot of documentation
2. API explorer
3. Overall documentation seems solid
4. API is straightforward and easy to use mostly
5. Providing test cards and banks are nice

### Notes

1. Missing story when new to payments domain. I like the sample projects offered by by Stripe, something like that could be helpful for people who are new to payments to get an idea using a working example app - _I see that we are building these_
2. what is difference between Checkout APIs and Payment APIs, they seem to have some shared functionality but API explorer doesn't talk about any differences, Is payments legacy API? This part is bit confusing when deciding what to choose
3. In the guides for the checkout API, there is no direct instruction on where to get the SDK/Libs from, its mentioned but it would be nice to just provide that info as first step to make it easy for the devs like how stripe does it
4. The merchant account to use is not clear enough, since there were two accounts and only one worked for the API
5. In "Step 2: Add Components to your payments form" of "Web Components integration guide" I find the flow would be easier if example is followed by explanation
6. Would be nice to give example of using reference for mounting Components in Vue/React
7. The redirect behaviors from `checkout.createFromAction` is not very clear for 3DS2 when POST redirect is involved, this is where I spent the most time, so might be nice to provide some simple samples here else you would have to dig up https://github.com/adyen-examples/ and go through it to see how it can be done
8. I don't see the purpose of `onAdditionalDetails` on the configuration and its not demonstrated why its needed?
9. The amount gets messed up during iDEAL redirection when there are decimal points
10. The country dropdown in the billing address of component doesn't work with standard list filtering
11. The state field in the billing address of component resets when selecting country
12. The "Web Components integration guide" doesn't mention that 3DS payments require slightly different flow, one might not realize this when doing the integration. I had to do a lot of rework when I realized this
13. The test environment feels a bit buggy and unpolished and hence may not give a lot of confidence to Dev who are evaluating, trying out the API
14. How is the live account API URL obtained?
15. Payment types in paymentMethods response doesn't always match with component id used for `checkout.create("card")` which is bit weird. For example card/scheme
16. Finding the correct id to use for component is not straight forward, you have to drill into multiple levels of docs for each type to find that. It would be nice to have a table of all components with their ids may here https://docs.adyen.com/checkout/supported-payment-methods
17. The amount value field in API spec is int64, shouldn't this be float?
18. Error messages are not very helpful as they don't give a lot of info on what is wrong with the request. This results in developer having to do a lot of trial and error and ultimately being stuck here

### Website

1. getting started: seems bit involved as it requires you to get advice from Adyen team - for mid market it would be nicer if this directs more towards self service
2. payment fundamentals:
   1. payments glossary: terms are in alphabetical order which means you jump around a lot if you are new to the domain(which could be the case for mid market), might be nicer to explain in logical sequence with an example

### Good to have

1. Go Lib for API, I tried to generate with openAPI generator, its looks fine, but with some effort we can make it less verbose and more user friendly as well, else we would need a little effort to package and validate
   1. its quite a lot of effort to do the manual integration sine the APIs have pretty complex response/req objects
   2. manual integration would be hard to maintain when API changes
   3. Companies using Go might prefer Stripe since they provide a Go lib
2. Golang sample app would be great, I can work further on this to add to to examples repo, but this would be much nicer if done after above
