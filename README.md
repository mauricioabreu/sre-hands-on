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

## Performance

Observing how a system performs is not something new, but it changed drastically over the years.

I remember the time we had no idea if a system was running under good conditions. How did we know something was failing? We didn't. Customers used to be our alert system. If everything was silent, it was a sign that everything was working fine. Otherwise, customers would be flooding the project manager's inbox. At that time, we had scripts reading the last lines of log files looking for words like "ERROR", "Exception" and so on.

Nowadays, we have plenty of tools and strategies to parse, collect and use all the information we can extract from systems. This wide and diverse range of tools may confuse even experiencied developers. What tools do we need? Which metrics we must collect? Will these numbers show how buggy is my code?

If you are working in a healthy environment, you should not worry whether these numbers will impact your reasoning about your team's code quality. Measuring should serve for one thing only: understanding systems. And when I say _understand_, I'm talking about awaraness. We should not need to keep eyes on log parsers, charts or email inbox to check if all systems are up and running. Our job, as developers, is to automate these processes.

### Availability

**Availability** is the word used to define whether a system is able to complete a request from a client. We use the word _client_ because we don't care whether the request comes from humans or tools.

If a GET request to `/todos` route works, it means the route works and the requests completes. Does it mean that the application is available?

Availability if often measured by using the following formula:

$sucessful / (successful + failed)$

The result is a percentage like 99,9%.

Let's check some examples:

* Our app receives *1000* requests in the past minute.
* 100 out of 1000 requests fail
* 900 / (900 + 100) = 0.9
* Our app has 90% of availability
