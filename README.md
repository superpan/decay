#decay
[![Build Status](https://travis-ci.org/superpan/decay.png?branh=master)](https://travis-ci.org/superpan/decay)

sorting algorithms implemented for go

### decay.Reddit(float64) fn(float64, float64, Time) float64
Popularity calcuation based on published Reddit rankings.  [http://amix.dk/blog/post/19588](http://amix.dk/blog/post/19588)
```
  import (
    "decay"
    "time"
  )

  func main() {
    redditScore := decay.Reddit(45000)
    fmt.println(redditScore(10, 3, time.now()))
  }
```

### decay.HackerNews(float64) fn(int, int, Time) float64
Popularity calcuation based on published HackerNews rankings.  [http://amix.dk/blog/post/19574](http://amix.dk/blog/post/19574)
```
  import (
    "decay"
    "time"
  )

  func main() {
    hackerScore := decay.HackerNews(1.9)
    fmt.println(hackerScore(10, 3, time.now()))
  }
```

### decay.WilsonScore(float64) fn(float64, float64) float64
Non-decay based calculation based on statistical confidence of data.

```
  import (
    "decay"
    "time"
  )

  func main() {
    wilsonScore := decay.WilsonScore(1.96)
    fmt.println(wilsonScore(10, 3))
  }
```