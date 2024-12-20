# Search engine

[![Matcher Test](https://github.com/lukaszfabia/search-engine/actions/workflows/go-test.yml/badge.svg)](https://github.com/lukaszfabia/search-engine/actions/workflows/go-test.yml)

Quick Friday project. I've tried to implement own search functions and show it in action on a simple page build with `templ` and `htmx`.

I've built it using **Trie data structure**. Implementation of the `matcher` can be found [here](/internal/matcher/). For `Suggestion` function I've used Levenshtein's distance.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.
