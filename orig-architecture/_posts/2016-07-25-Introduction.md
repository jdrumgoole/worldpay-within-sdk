---
layout: post
title:  "1. Introduction"
date:   2016-07-27 12:13:03 +0100
categories: jekyll update
---
Introduction
============

Purpose
-------

The purpose of this document is to define a reference architecture and
set of APIs for providing secure payment facilities for machine to
machine payments in the Internet of Things (IoT) via the Worldpay
Online.worldpay.com platform.

Online.worldpay.com is the Worldpay name for the Worldpay Online
Payments API, <https://online.worldpay.com/>.

Scope
-----

This document will not specify the communication layer (i.e. Bluetooth
Low Energy, NFC, GPRS, etc) of service discovery, but rather the steps
necessary independent of communication layer.

This document shall consider online payments only.

It shall consider all aspects of an IoT payments relating to the
provision of a service, including:

-   Service Discovery;

-   Service Negotiation;

-   Payment;

-   Service Delivery.

Worldpay Within provides a Thing with the ability to make and / or
receive payments from other Things, as such a Thing can be both a
service consumer and service provider.

A Worldpay Within enabled Thing can be both the merchant and the
consumer. This document considers the interaction between two things,
for the purpose of clarity, Thing A shall be the consumer of the
service, Thing B being the service provider, however it shall be equally
possible for Thing B to consume services from Thing A.

Audience
--------

Public external consumption.