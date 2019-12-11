# Stringify

Stringifies json into an escaped string that can be put into json.
I found this useful for creating AWS lambda API Gateway Proxy tests.

By default, gets the string from the clipboard and replaces it with the new string.
Alternatively you can accept user input, i.e.:

```
$ stringify -i <<EOF
    {
        "field": "one"
    }
EOF
"{\"field\":\"one\"}"
```
(it will also copy the above to your clipboard)

This also works with vim selections :)
```
:'<,'>!stringify
```
