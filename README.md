# lobby
lobby, load balancer in Go


## Design Specs

![1-h1jmjyq0y0k2hlplrojmbq](https://cloud.githubusercontent.com/assets/3850661/15625145/eb14e986-24ce-11e6-969d-a27f5eefab53.png)

https://medium.com/@roylee0704/load-balancer-93118673f31a#.hy2o1x15z


## Simulations

Simulating lobby with `requester:100` , `worker:10`:
- each requester sends one request at a time, waits for response, and rest for 1-4 seconds before another new request.
- load-balancer maintains a priority-based pool of workers, where most lightly-loaded worker will be assigned to serve the request.
- each worker maintains a job queue of `r requesters`.
- each job takes about 1-5seconds to process.
- upon job completion, worker reports itself to load-balancer and requester (request-response model).


Note:
- currently each print shows average workload in the worker pool.


[![asciicast](https://asciinema.org/a/4dc1js4vtdfehxuo6sz8gpdx6.png)](https://asciinema.org/a/4dc1js4vtdfehxuo6sz8gpdx6)
