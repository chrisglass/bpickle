This is a work in progress library to serialize and deserialize bpickle.

Why bpickle?
============

You might ask, rightfully, why yet another serialization library is needed.
The simple answer is that I need to work with a system that implements
its own serialization library, called bpickle.

Should I used bpickle?
=======================

Probably not.

If you're looking for a network efficient transport format that is multi-language
and multi-platforms, I recommend having a look at protocol buffers instead.

But then why are you developing it?
===================================

It scratches my itch, and it's a fun and interesting way for me to learn more
about the inner workings of Go.
