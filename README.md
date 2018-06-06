# go-dark

A preprocessor for Go code that adds functional programming features.

## Usage

Name your source files `*.dark` and run go-dark. This will produce corresponding Go code which can be built with the go tool as usual. The dark magic mostly happens in m4, so you need that installed first.

## Typed lambdas

In certain contexts, go-dark lets you express function literals using lambda expression of the form `args_list => return_type = expr`. This gets expanded to `func(args_list) return_type { return expr }`.

    s := lambda(x string => string = x + "!")("hello")
    -> s = "hello!"

## Fmap, filter, and friends

go-dark lets you apply fmap, filter, etc to arrays of any type. This sorta kinda adds generics to Go, but only within the context of these few methods. Think of them as magic functions like append, len, etc, only darker.

You can construct an iterator via `iter(any_type).from(your_array)`. To "force" an Iterator back into an array, use the `array(any_type)` method.

    ints := []int{1, 2, 3}
    strings := iter(int).from(ints).
      fmap(x int => string = strconv.Itoa(x)).
      fmap(x string => string = x + "!").
      array(string)
    -> strings = {"1!", "2!", "3!"}

## But why?

Go is so low-level it feels like writing in C or assembly. Might as well goober it up with macros.

Ryanne
