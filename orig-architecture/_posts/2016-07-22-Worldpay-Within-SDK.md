---
layout: post
title:  "4. Worldpay Within SDK"
date:   2016-07-27 12:13:03 +0100
categories: jekyll update
---
Worldpay Within SDK
===================

Overview
--------

As a complement to this architecture document, Worldpay will release a
Worldpay Within SDK. The intention of the SDK is to encapsulate
implementation and therefore assist third party vendors in integration
payments into their IoT solutions.

Scope
-----

The core of the SDK will be developed using the Go programming language
with wrappers created for the Java, Node.JS and Python, C/C++ and C\#
(tbc). Service discovery and broadcast will be implemented using TCP/IP
networking. Transport Security needs to be reviewed and implemented
before going into production.

Architecture
------------

The SDK is composed of four areas of concern:

1.  The IoT Thing, a functional device developed by a third party and
    consumes the Worldpay Within SDK

2.  SDK Wrapper – A wrapper for the SDK written in a variety of
    languages to support calling the SDK from those languages. This
    wrapper code interfaces with the SDK core via RPC.

3.  The SDK RPC interface allows for other languages to call core
    functions

4.  SDK Core – The actual implementation of the Worldpay
    Within specification. Developed using Go programming language

Sample Code
-----------

TO DO – Update this section with reference to sample code and possibly
move to appendix

Ongoing concerns – For Slide Deck??
-----------------------------------

There are a number of concerns that need to be discussed, solved and
implemented in the SDK. These include but are not limited to:

-   Persistence and provision of secure data such as API keys and
    consumer payment credentials

-   Transport level security between clients and RPC host in local
    environment