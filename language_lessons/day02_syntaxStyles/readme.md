![](https://github.com/irisida/golang/blob/master/assets/freegopher.png)

# Day 2

Today we're going to cover some more variables and in particular syntactic style, what does Go code look like? Essentially we're going to cover longform, shortform & inferred variable declaration and how to do multiple variable declaration as well as concise declaration with datatypes.

What the hell is with lots of different ways to declare variables? Fair question, not one I can answer, something that can be seen in several languages really. At a guess, key influences at concept and design stage couldn't agree on a single unifying method because it's not an argument easily won.

## variables declaration styles

The longform - This is like hearing Hugh Grant speak, it's formal, proper, articulated. Feels a bit snooty.
![](https://github.com/irisida/golang/blob/master/assets/day2/day02.001.png)

The short form and inferred form - this is loose and casual, might raise an eyebrow at a code review depending on the place you're at and the rules about this type of thing. It's not wrong, but it just gets on som people's goat and you risk a PR rejection.

![](https://github.com/irisida/golang/blob/master/assets/day2/day02.002.png)
1[](https://github.com/irisida/golang/blob/master/assets/day2/day02.004.png)

In Go we can make shortcuts to keep code elegant and tidy when declaring multiple variables, these can even be grouped which might seem like a double edged sword because it can read/imply like there is a relationship that is not actyually present.
![](https://github.com/irisida/golang/blob/master/assets/day2/day02.003.png)

Here we see the grouping example that looks like a struct (think class if you're coming from other languages, real world object or thing if you've never programmed before)
![](https://github.com/irisida/golang/blob/master/assets/day2/day02.006.png)

finally we can declare multiple variables with a single type on the same line. Again the primary purpose is concise code blocks and this is both idiomatic to Go and often referenced as an elegant modern language touch.
![](https://github.com/irisida/golang/blob/master/assets/day2/day02.005.png)
