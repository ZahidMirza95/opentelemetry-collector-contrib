[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# elasticsearchreceiver

## Metrics

These are the metrics available for this scraper.

| Name | Description | Unit | Type | Attributes |
| ---- | ----------- | ---- | ---- | ---------- |
| **elasticsearch.breaker.memory.estimated** | Estimated memory used for the operation. | By | Gauge(Int) | <ul> <li>circuit_breaker_name</li> </ul> |
| **elasticsearch.breaker.memory.limit** | Memory limit for the circuit breaker. | By | Sum(Int) | <ul> <li>circuit_breaker_name</li> </ul> |
| **elasticsearch.breaker.tripped** | Total number of times the circuit breaker has been triggered and prevented an out of memory error. | 1 | Sum(Int) | <ul> <li>circuit_breaker_name</li> </ul> |
| **elasticsearch.cluster.data_nodes** | The number of data nodes in the cluster. | {nodes} | Sum(Int) | <ul> </ul> |
| **elasticsearch.cluster.health** | The health status of the cluster. Health status is based on the state of its primary and replica shards. Green indicates all shards are assigned. Yellow indicates that one or more replica shards are unassigned. Red indicates that one or more primary shards are unassigned, making some data unavailable. | {status} | Sum(Int) | <ul> <li>health_status</li> </ul> |
| **elasticsearch.cluster.nodes** | The total number of nodes in the cluster. | {nodes} | Sum(Int) | <ul> </ul> |
| **elasticsearch.cluster.shards** | The number of shards in the cluster. | {shards} | Sum(Int) | <ul> <li>shard_state</li> </ul> |
| **elasticsearch.node.cache.evictions** | The number of evictions from the cache. | {evictions} | Sum(Int) | <ul> <li>cache_name</li> </ul> |
| **elasticsearch.node.cache.memory.usage** | The size in bytes of the cache. | By | Sum(Int) | <ul> <li>cache_name</li> </ul> |
| **elasticsearch.node.cluster.connections** | The number of open tcp connections for internal cluster communication. | {connections} | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.cluster.io** | The number of bytes sent and received on the network for internal cluster communication. | By | Sum(Int) | <ul> <li>direction</li> </ul> |
| **elasticsearch.node.cluster.io.received** | The number of bytes received on the network for internal cluster communication. | By | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.cluster.io.sent** | The number of bytes sent on the network for internal cluster communication. | By | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.disk.io.read** | The total number of kilobytes read across all file stores for this node. | By | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.disk.io.write** | The total number of kilobytes written across all file stores for this node. | By | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.documents** | The number of documents on the node. | {documents} | Sum(Int) | <ul> <li>document_state</li> </ul> |
| **elasticsearch.node.fs.disk.available** | The amount of disk space available across all file stores for this node. | By | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.http.connections** | The number of HTTP connections to the node. | {connections} | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.open_files** | The number of open file descriptors held by the node. | {files} | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.operations.completed** | The number of operations completed. | {operations} | Sum(Int) | <ul> <li>operation</li> </ul> |
| **elasticsearch.node.operations.time** | Time spent on operations. | ms | Sum(Int) | <ul> <li>operation</li> </ul> |
| **elasticsearch.node.shards.data_set.size** | Total data set size of all shards assigned to the node. This includes the size of shards not stored fully on the node, such as the cache for partially mounted indices. | By | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.shards.reserved.size** | A prediction of how much larger the shard stores on this node will eventually grow due to ongoing peer recoveries, restoring snapshots, and similar activities. A value of -1 indicates that this is not available. | By | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.shards.size** | The size of the shards assigned to this node. | By | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.thread_pool.tasks.finished** | The number of tasks finished by the thread pool. | {tasks} | Sum(Int) | <ul> <li>thread_pool_name</li> <li>task_state</li> </ul> |
| **elasticsearch.node.thread_pool.tasks.queued** | The number of queued tasks in the thread pool. | {tasks} | Sum(Int) | <ul> <li>thread_pool_name</li> </ul> |
| **elasticsearch.node.thread_pool.threads** | The number of threads in the thread pool. | {threads} | Sum(Int) | <ul> <li>thread_pool_name</li> <li>thread_state</li> </ul> |
| **elasticsearch.node.translog.operations** | Number of transaction log operations. | {operations} | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.translog.size** | Size of the transaction log. | By | Sum(Int) | <ul> </ul> |
| **elasticsearch.node.translog.uncommitted.size** | Size of uncommitted transaction log operations. | By | Sum(Int) | <ul> </ul> |
| **elasticsearch.os.cpu.load_avg.15m** | Fifteen-minute load average on the system (field is not present if fifteen-minute load average is not available). | 1 | Gauge(Double) | <ul> </ul> |
| **elasticsearch.os.cpu.load_avg.1m** | One-minute load average on the system (field is not present if one-minute load average is not available). | 1 | Gauge(Double) | <ul> </ul> |
| **elasticsearch.os.cpu.load_avg.5m** | Five-minute load average on the system (field is not present if five-minute load average is not available). | 1 | Gauge(Double) | <ul> </ul> |
| **elasticsearch.os.cpu.usage** | Recent CPU usage for the whole system, or -1 if not supported. | % | Gauge(Int) | <ul> </ul> |
| **elasticsearch.os.memory** | Amount of physical memory. | By | Sum(Int) | <ul> <li>memory_state</li> </ul> |
| **jvm.classes.loaded** | The number of loaded classes | 1 | Gauge(Int) | <ul> </ul> |
| **jvm.gc.collections.count** | The total number of garbage collections that have occurred | 1 | Sum(Int) | <ul> <li>collector_name</li> </ul> |
| **jvm.gc.collections.elapsed** | The approximate accumulated collection elapsed time | ms | Sum(Int) | <ul> <li>collector_name</li> </ul> |
| **jvm.memory.heap.committed** | The amount of memory that is guaranteed to be available for the heap | By | Gauge(Int) | <ul> </ul> |
| **jvm.memory.heap.max** | The maximum amount of memory can be used for the heap | By | Gauge(Int) | <ul> </ul> |
| **jvm.memory.heap.used** | The current heap memory usage | By | Gauge(Int) | <ul> </ul> |
| **jvm.memory.nonheap.committed** | The amount of memory that is guaranteed to be available for non-heap purposes | By | Gauge(Int) | <ul> </ul> |
| **jvm.memory.nonheap.used** | The current non-heap memory usage | By | Gauge(Int) | <ul> </ul> |
| **jvm.memory.pool.max** | The maximum amount of memory can be used for the memory pool | By | Gauge(Int) | <ul> <li>memory_pool_name</li> </ul> |
| **jvm.memory.pool.used** | The current memory pool memory usage | By | Gauge(Int) | <ul> <li>memory_pool_name</li> </ul> |
| **jvm.threads.count** | The current number of threads | 1 | Gauge(Int) | <ul> </ul> |

**Highlighted metrics** are emitted by default. Other metrics are optional and not emitted by default.
Any metric can be enabled or disabled with the following scraper configuration:

```yaml
metrics:
  <metric_name>:
    enabled: <true|false>
```

## Resource attributes

| Name | Description | Type |
| ---- | ----------- | ---- |
| elasticsearch.cluster.name | The name of the elasticsearch cluster. | String |
| elasticsearch.node.name | The name of the elasticsearch node. | String |

## Metric attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| cache_name | The name of cache. | fielddata, query |
| circuit_breaker_name (name) | The name of circuit breaker. |  |
| collector_name (name) | The name of the garbage collector. |  |
| direction | The direction of network data. | received, sent |
| document_state (state) | The state of the document. | active, deleted |
| fs_direction (direction) | The direction of filesystem IO. | read, write |
| health_status (status) | The health status of the cluster. | green, yellow, red |
| memory_pool_name (name) | The name of the JVM memory pool. |  |
| memory_state (state) | State of the memory | free, used |
| operation (operation) | The type of operation. | index, delete, get, query, fetch, scroll, suggest, merge, refresh, flush, warmer |
| shard_state (state) | The state of the shard. | active, relocating, initializing, unassigned |
| task_state (state) | The state of the task. | rejected, completed |
| thread_pool_name | The name of the thread pool. |  |
| thread_state (state) | The state of the thread. | active, idle |