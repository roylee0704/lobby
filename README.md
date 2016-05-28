# lobby
lobby, load balancer in Go


## Simulations

Simulating lobby with `requester:100` , `worker:10`:
- each requester sends one request at a time, waits for response, and rest for 1-4 seconds before another new request.
- each job takes about 1-5seconds to process.
- each worker maintains a job queue of `r requesters`.

Note:
- currently each print shows average workload in the worker pool.


[![asciicast](https://asciinema.org/a/4dc1js4vtdfehxuo6sz8gpdx6.png)](https://asciinema.org/a/4dc1js4vtdfehxuo6sz8gpdx6)
