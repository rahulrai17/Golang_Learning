Here's a set of interview-style questions on maps in Go, organized from beginner to expert levels. Each question includes a sample solution and an explanation of the approach.

---

### Beginner Level

1. **Counting Character Occurrences**

   - **Question**: Write a function that takes a string and returns a map where the keys are characters, and values are the number of times each character appears in the string.
   - **Solution**:

     ```go
     package main

     import "fmt"

     func countChars(s string) map[rune]int {
         charCount := make(map[rune]int)
         for _, char := range s {
             charCount[char]++
         }
         return charCount
     }

     func main() {
         text := "hello"
         fmt.Println(countChars(text)) // Output: map[h:1 e:1 l:2 o:1]
     }
     ```

   - **Explanation**: This function iterates over each character in the string. It uses a map to count occurrences, incrementing the count each time a character appears. Maps are well-suited for this due to their average constant time complexity for insertions and lookups.

2. **Mapping IDs to Names**

   - **Question**: Write a function that takes a list of IDs and names as input and returns a map where each ID is mapped to the corresponding name.
   - **Solution**:

     ```go
     package main

     import "fmt"

     func mapIDsToNames(ids []int, names []string) map[int]string {
         idNameMap := make(map[int]string)
         for i, id := range ids {
             if i < len(names) {
                 idNameMap[id] = names[i]
             }
         }
         return idNameMap
     }

     func main() {
         ids := []int{101, 102, 103}
         names := []string{"Alice", "Bob", "Charlie"}
         fmt.Println(mapIDsToNames(ids, names)) // Output: map[101:Alice 102:Bob 103:Charlie]
     }
     ```

   - **Explanation**: Here, we map each `id` to a `name` by indexing them in order. Maps allow efficient lookups, enabling us to retrieve a name by its ID in constant time.

---

### Intermediate Level

3. **Finding Unique Elements**

   - **Question**: Write a function that takes a slice of integers and returns a slice of unique elements.
   - **Solution**:

     ```go
     package main

     import "fmt"

     func uniqueElements(nums []int) []int {
         uniqueMap := make(map[int]bool)
         result := []int{}

         for _, num := range nums {
             if !uniqueMap[num] {
                 uniqueMap[num] = true
                 result = append(result, num)
             }
         }

         return result
     }

     func main() {
         nums := []int{1, 2, 2, 3, 4, 4, 5}
         fmt.Println(uniqueElements(nums)) // Output: [1 2 3 4 5]
     }
     ```

   - **Explanation**: We use a map to track elements we've already seen by marking them as `true`. This approach is efficient for large inputs, as checking and inserting in maps is generally constant time.

4. **Find Missing Element**

   - **Question**: Given two slices `a` and `b` where `b` contains all elements of `a` except one, find the missing element from `a`.
   - **Solution**:

     ```go
     package main

     import "fmt"

     func findMissing(a, b []int) int {
         elementMap := make(map[int]int)

         for _, num := range a {
             elementMap[num]++
         }
         for _, num := range b {
             elementMap[num]--
         }

         for key, value := range elementMap {
             if value > 0 {
                 return key
             }
         }
         return -1
     }

     func main() {
         a := []int{1, 2, 3, 4, 5}
         b := []int{1, 2, 3, 5}
         fmt.Println(findMissing(a, b)) // Output: 4
     }
     ```

   - **Explanation**: This approach leverages a map to count the occurrences of each element in `a` and `b`. If a number has a positive count after processing both slices, it's the missing number. This solution is efficient and avoids sorting or nested loops.

---

### Advanced Level

5. **Finding Common Elements in Two Lists**

   - **Question**: Write a function that finds the common elements in two slices of integers.
   - **Solution**:

     ```go
     package main

     import "fmt"

     func commonElements(a, b []int) []int {
         countMap := make(map[int]bool)
         result := []int{}

         for _, num := range a {
             countMap[num] = true
         }

         for _, num := range b {
             if countMap[num] {
                 result = append(result, num)
                 countMap[num] = false
             }
         }

         return result
     }

     func main() {
         a := []int{1, 2, 3, 4, 5}
         b := []int{3, 4, 5, 6, 7}
         fmt.Println(commonElements(a, b)) // Output: [3 4 5]
     }
     ```

   - **Explanation**: Using a map, we store elements from the first slice. We then check if each element of the second slice exists in the map, which allows us to find common elements in linear time without sorting or nested loops.

6. **Inverting a Map**

   - **Question**: Write a function that takes a map of integers to strings and returns a new map that inverts the original, mapping strings to integers.
   - **Solution**:

     ```go
     package main

     import "fmt"

     func invertMap(original map[int]string) map[string]int {
         inverted := make(map[string]int)
         for key, value := range original {
             inverted[value] = key
         }
         return inverted
     }

     func main() {
         m := map[int]string{1: "one", 2: "two", 3: "three"}
         fmt.Println(invertMap(m)) // Output: map[one:1 two:2 three:3]
     }
     ```

   - **Explanation**: This function iterates over the original map and inverts the key-value pairs by setting `value` as the key in the new map. This technique is often used when needing fast lookups in both directions.

---

### Expert Level

7. **Detecting a Cycle in a Directed Graph**

   - **Question**: Implement a function to detect a cycle in a directed graph represented as an adjacency list using a map.
   - **Solution**:

     ```go
     package main

     import "fmt"

     func hasCycle(graph map[int][]int) bool {
         visited := make(map[int]bool)
         recursionStack := make(map[int]bool)

         var dfs func(int) bool
         dfs = func(node int) bool {
             if recursionStack[node] {
                 return true
             }
             if visited[node] {
                 return false
             }

             visited[node] = true
             recursionStack[node] = true
             for _, neighbor := range graph[node] {
                 if dfs(neighbor) {
                     return true
                 }
             }
             recursionStack[node] = false
             return false
         }

         for node := range graph {
             if dfs(node) {
                 return true
             }
         }
         return false
     }

     func main() {
         graph := map[int][]int{
             1: {2},
             2: {3},
             3: {4},
             4: {2},
         }
         fmt.Println(hasCycle(graph)) // Output: true
     }
     ```

   - **Explanation**: We use DFS to detect cycles by maintaining a recursion stack that keeps track of nodes currently in the recursive call chain. If we revisit a node in the recursion stack, there’s a cycle. Maps help efficiently store graph structure and visited state.

8. **Implementing a Cache with Expiry**

   - **Question**: Design a cache that can store key-value pairs with an expiry time. If a key is requested after its expiry time, return a not-found error.
   - **Solution**:

     ```go
     package main

     import (
         "fmt"
         "time"
     )

     type CacheItem struct {
         value     interface{}
         expiry    time.Time
     }

     type ExpiringCache struct {
         items map[string]CacheItem
     }

     func NewExpiringCache() *ExpiringCache {
         return &ExpiringCache{
             items: make(map[string]CacheItem),
         }
     }

     func (c *ExpiringCache) Set(key string, value interface{}, duration time.Duration) {
         c.items[key] = CacheItem{
             value:  value,
             expiry: time.Now().Add(duration),
         }
     }

     func (c *ExpiringCache) Get(key string) (interface{}, bool) {
         item, exists := c.items[key]
         if !exists || time.Now().After(item.expiry) {
             delete(c.items, key)
             return nil, false
         }
         return item.value, true
     }

     func main() {
         cache := NewExpiringCache()
         cache.Set("key1", "value1", 5*time.Second
     ```

)
val, found := cache.Get("key1")
fmt.Println(val, found) // Output: value1 true

         time.Sleep(6 * time.Second)
         val, found = cache.Get("key1")
         fmt.Println(val, found) // Output: <nil> false
     }
     ```

- **Explanation**: This design uses a struct to store each cache item’s value and expiry time, and a map to store items by key. When `Get` is called, the function checks if the item exists and if it’s expired. Expired items are removed to prevent memory leaks.

---

These questions cover various Go map use cases and test candidates on common interview scenarios involving data structure manipulation, efficient searching, and edge cases.
