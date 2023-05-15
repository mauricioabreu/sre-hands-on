# SRE Hands On

This tutorials is a hands on to [Site Reliability Engineering](https://en.wikipedia.org/wiki/Site_reliability_engineering).

## Running

```
go run main.go
```

## Routes

Our app has only three features:

* Reading todos (GET)
* Creating a todo (POST)
* Deleting a todo (DELETE)

The _sample.http_ file describes some examples of requests that can be performed.

## Performance

Observing how a system performs is not something new, but it changed drastically over the years.

I remember the time we had no idea if a system was running under good conditions. How did we know something was failing? We didn't. Customers used to be our alert system. If everything was silent, it was a sign that everything was working fine. Otherwise, customers would be flooding the project manager's inbox. At that time, we had scripts reading the last lines of log files looking for words like "ERROR", "Exception" and so on.

Nowadays, we have plenty of tools and strategies to parse, collect and use all the information we can extract from systems. This wide and diverse range of tools may confuse even experiencied developers. What tools do we need? Which metrics we must collect? Will these numbers show how buggy is my code?

If you are working in a healthy environment, you should not worry whether these numbers will impact your reasoning about your team's code quality. Measuring should serve for one thing only: understanding systems. And when I say _understand_, I'm talking about awaraness. We should not need to keep eyes on log parsers, charts or email inbox to check if all systems are up and running. Our job, as developers, is to automate these processes.
