1. Only the specific length of items in cache
2. LRU - only most recently used items are stored in the cache


For a true LRU cache - 
1. If an item already exist in the cache we need to remove it and add it to the beginning of the cache
2. An order of the items in the cache needs to be maintained
3. Deletion happens at the tail of the cache and addition happens at the head of the cache


Sample output of code:

```
Start LRU Cache with size:  5
Add ravi
{ravi}
Add shruti
{shruti}<-->{ravi}
Add kalpna
{kalpna}<-->{shruti}<-->{ravi}
Add mukta
{mukta}<-->{kalpna}<-->{shruti}<-->{ravi}
Add rohan
{rohan}<-->{mukta}<-->{kalpna}<-->{shruti}<-->{ravi}
Add varsha
Remove ravi
{varsha}<-->{rohan}<-->{mukta}<-->{kalpna}<-->{shruti}
```