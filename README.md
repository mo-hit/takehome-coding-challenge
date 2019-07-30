# Gopher Gold

Gopher Gold is set to be the next great unicorn company. It will be the world's leading joke sharing platform and we need your help to build it! Webscale!

## Setup

To run this project you must have [Go](https://golang.org/dl/) installed.

To get started, run `make test`.

You should see three failing tests. Your job is to get these tests passing!

Remember, we'd much rather see an incomplete response than nothing at all. Wherever you get to in the time limit, commit and push it.

When implementing the features give some thought to the following criteria:

* Is your code idiomatic?
* Are you providing a good end user experience?

## Feature One

It's great that we can efficiently distribute our excellent jokes to the world. If we're going to raise our next round, though, we need to crowdsource this stuff. Implement the joke submission handler to receive jokes submitted by our users and store them in our proprietary NoSQL database.

## Feature Two

We want to ensure that we only accept the best jokes into our database. Make sure we never store any jokes rated below a five!

Also, our joke rating scale is strictly out of ten. Make sure we never ever accept a joke with a rating above ten. That's just silly.

## Running the app

`make run` and point a web browser to [http://localhost:3000](http://localhost:3000)

## Extra credit

If you end up with extra time feel free to improve the codebase or add any features you think will be handy. Remember not to exceed the time limit, though.
