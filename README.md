# SRE Hands On

This tutorials is a hands on to [Site Reliability Engineering](https://en.wikipedia.org/wiki/Site_reliability_engineering).

First, we will focus on learning some basic aspects of SRE. Then, we will learn how to collect **Service Level Indicators (SLIs)** and define **Service Level Objectives (SLOs)**.

If you wish to learn more about the teory behind SRE, I recommend you read the [SRE Workbook from Google](https://sre.google/workbook/table-of-contents/)

## Performance

Observing how a system performs has drastically changed over the years.

I remember a time when we had no idea if a system was running under good conditions. How did we know if something was failing? We didn't. Customers used to be our alert system. If everything was silent, it was a sign that everything was working fine. Otherwise, customers would flood the project manager's inbox. At that time, we had scripts reading the last lines of log files looking for words like _ERROR_ and _Exception_.

Nowadays, we have plenty of tools and strategies to **parse**, **collect**, and **analyze** all the information we can extract from systems. This wide and diverse range of tools may confuse even experienced developers. Which tools do we need? Which metrics must we collect? Will these numbers show how buggy my code is?

If you are working in a healthy environment, you should not worry about whether these numbers will impact your reasoning about your team's code quality. Measuring should serve one purpose only: _understanding_ systems. And when I say understand, I'm talking about awareness. We should not need to keep our eyes on log parsers, charts, or email inboxes to check if all systems are up and running. Our job, as developers, is to automate these processes.

## Understanding SLOs and SLIs

**SLIs** are metrics that quantify the performance and behavior of your system or service.

**SLOs** define the acceptable levels of performance or reliability based on SLIs.

SLOs are a tool to:
* Help you decide what the next steps are in evolving your systems
* Identify which parts of the system need more attention
* Negotiate future improvements with stakeholders

All of this is driven by data.

A Service Level Objective is a _target_, and as such, it must be designed and defined with all the people responsible for the target, from the infrastructure teams to stakeholders. It is an important commitment, and transparency is mandatory.

### Availability

**Availability** is the word used to define whether a system is able to complete a request from a client. We use the word _client_ because we don't care whether the request comes from humans or tools.

If a GET request to the `/todos` route works, it means the route works, and the request completes. But does it mean that the application is available?

Availability is often measured using the following formula:

$successful / (successful + failed)$

Let's check some examples:
* Our app received *1000* requests in the past minute
* *100* out of 1000 requests failed
* 900 / (900 + 100) = 0.
* Our app has 90% availability

### Latency

**Latency** measures how long a request takes to complete. How long does it take to process a job?

We use latency to measure the overall user experience.

Slow responses will force users to leave our websites, and that's something we don't want. Availability is important. We want our applications to respond with **HTTP 200 OK**, but a good status can't come with high latency. We don't want our applications to respond fast with **HTTP 5xx errors**, just as we don't want **HTTP 2xx** responses taking 10 seconds to complete. Keep in mind that slow errors are even worse than fast errors.

Measuring latency is a real challenge. At first glance, it may seem like a good idea to store every single response time and then calculate the average time of all the requests. However, it turns out that this is not a good idea. Let's take a closer look at the following example:

The _Just Code_ company launches a fresh new app. On the first day, it only has a few users navigating the site. Upset about the poor numbers, the marketing team decides to invest in advertising. And then, like magic, the app starts to receive a tremendous amount of requests. A month later, the CTO becomes concerned about the costs in the cloud. All the metrics are working fine, but they seem to be very expensive.

Why do you think they are not as cheap as some people might wonder?

One of the best ways to collect latency metrics without sacrificing too much information and paying less for it is by storing the information in buckets and calculating percentiles of latency.
