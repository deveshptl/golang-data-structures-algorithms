## Contributing Guide:


### Prerequisites

The only prerequisites is to have go installed on you computer in order to contribute. Follow how to install instructions from [golang download instructions](https://golang.org/doc/install) if you don't already have.


### Issues

If you find anything wrong about any data structures or algorithm not working the way they should then feel free to create a new issue, but before doing that make sure that no one else has created the same issue.

It would be better if you create a issue first before implementing anything, so that other people don't accidently duplicate your effort.


### Pull Requests

Before submitting a pull request, please make sure the following is done:

1. Fork the repository.
2. Keep the implementations in proper directory. For e.g all data structures are present in the root directory, while the implementations of algorithms are in its own directory. Let's suppose say that you are implementing any sorting algorithm then you have to create a new folder under `algorithms` directory named `sorting` and as there are many sorting algorithms, you have to create a new folder under the `algorithms/sorting` as well and give a proper name to it like `quick_sort` and then create a new file under that directory. So your structure would look something like `algorithms/sorting/quick_sort/`. Also make sure to give proper descriptive file names as well.
3. Also include your implementation in the README.md file as well. For e.g the `bubble_sort` algorithm is under `algorithms/sorting/bubble_sort`. See how it has been written in README.md file in an alphabetical order.

```
sorting -
    bubble_sort - Bubble Sort
```

4. It would be great if you provide your tests inputs (CLI inputs) as comments at the end of the file, so before merging your pull request we can easily test your code and see if it's working correctly as expected or not.
5. Also make sure that all of your code is indented proper using `go fmt` command.

We look forward to your contributions. :joy:

```
HAPPY CODING
HAPPY LEARNING
```