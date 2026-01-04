


### Closures
- Function bundled with surrounding variables
- Why name closure -> it closes over the free variables(out of scope -> not an parameter or local)
```
function makeCounter() {
    let count = 0;
    return function() {
        count++;
        return count;
    };
}

const counter = makeCounter();
console.log(counter()); // 1
console.log(counter()); // 2

```