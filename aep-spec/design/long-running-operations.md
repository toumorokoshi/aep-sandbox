# Long Running Operations

The following document outlines some thoughts around the use of long-running
operations.

## When should LROs be used?

LROs are fundamentally a different protocol in communicating the status of an
operation: in contrast to the respone being returned synchronously, instead a
token is polled, eventually returning back a response (successful or otherwise).

LROs are more complex (and therefore expensive to implement) than synchronous
operations:

- They require the server to implement a mechanism to store tokens with
  operation information attached.
- The client must end multiple subsequest requests to retrieve the state of the
  operation, until it reaches a terminal state.

Therefore, synchronous operations should be the defaul, with an LRO being used
in extraordinary cases.

## Alternate pattern: synchronous RPC + state enum

For resource standard methods, LROs with long operation times (beyond 10 minute)
can still be problematic. For example, declarative clients such as Terraform
providers have an much better user experience if the resource mutations return
in a reasonable period of time: long operations increase the amount of time the
apply runs.

In these situations, an alternative pattern would be having a mutation return back
quickly, and instead use a state enum to indicate if the resource is ready to be used.
However, this forces the user to be aware that the resource is actually not yet ready.
