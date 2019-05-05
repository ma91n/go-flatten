# go-flatten
go-flatten is utility library to convert JSON to CSV by own logic

## Spec

### Basic1

```sh
echo '{"A":1, "B":1, "C":1}' | flatten -format ltsv
A:1, B:1, C:1
```

### Basic2: Object

```sh
echo '{"A":1, "B":1, "C":{"X":11, "Y":11, "Z":11}}' | flatten -format ltsv
A:1, B:1, C.X:11, C.Y:11, C.Z:11
```

### Basic3: Array

```sh
# array is not empty
echo '{"A":1, "B":1, "C":[1,2,3]}' | flatten -format ltsv
A:1, B:1, C:1
A:1, B:1, C:2
A:1, B:1, C:3

# array is empty. result will not be empty record.
echo '{"A":1, "B":1, "C":[]}' | flatten -format ltsv
A:1, B:1, C:nuil 

// array length is only one.
echo '{"A":1, "B":1, "C":[1]}' | flatten -format ltsv
A:1, B:1, C:1
```

### Basic4: Array of Object

```sh
# only one flatten
echo '{"A" : [{"a":11, "b":11, "c":11}, {"a":12, "b":12, "c":12}, {"a":13, "b":13, "c":13}]}' \
| flatten -format json
{ A : {a:11, b:11, c:11}}
{ A : {a:12, b:12, c:12}}
{ A : {a:13, b:13, c:13}}

# flatten recursive
echo '{"A" : [{"a":11, "b":11, "c":11}, {"a":12, "b":12, "c":12}, {"a":13, "b":13, "c":13}]}' \
| flatten -r -format ltsv
A.a:11, A.b:11, A.c:11
A.a:12, A.b:12, A.c:12
A.a:13, A.b:13, A.c:13
```

### Basic5: One Element is Array of Object

```sh
# flatten recursive
echo '{"A":1, "B":1, "C":[{"a":11, "b":11, "c":11}, {"a":12, "b":12, "c":12}, {"a":13, "b":13, "c":13}]}' \
| flatten -format json
{A:1, B:1, C:{a:11, b:11, c:11}}
{A:1, B:1, C:{a:12, b:12, c:12}}
{A:1, B:1, C:{a:13, b:13, c:13}}

# flatten recursive
echo '{"A":1, "B":1, "C":[{"a":11, "b":11, "c":11}, {"a":12, "b":12, "c":12}, {"a":13, "b":13, "c":13}]}' \
| flatten -r -format ltsv
A:1, B:1, C.a:11, C.b:11, C.c:11
A:1, B:1, C.a:12, C.b:12, C.c:12
A:1, B:1, C.a:13, C.b:13, C.c:13
```

### Basic5': Object key size is not correspond

```sh
echo '{"A":1, "B":1, "C":[{"a":11, "b":11, "c":11}, {"a":12, "b":12}, {}]}' \
| flatten -format json
{A:1, B:1, C:{a:11, b:11, c:11}}
{A:1, B:1, C:{a:12, b:12}}
{A:1, B:1, C:{}}

# The lack of keys is implicitly empty. That is, the above result is synonymous with.
echo '{"A":1, "B":1, "C":[{"a":11, "b":11, "c":11}, {"a":12, "b":12}, {}]}' \
| flatten -r -format ltsv
A:1, B:1, C.a:11, C.b:11, C.c:11
A:1, B:1, C.a:12, C.b:12
A:1, B:1, C.null
```

### Basic6: Double Array of Object

```
echo '{"A":{"a":1, "b":2, "c":3}, B:[{"d":11, "e":11, "f":11}, {"d":12, "e":12, "f":12}, {"d":13, "e":13, "f":13}]}' \
| flatten -format json
{A:{a:1, b:2, c:3}, B:{a:11, b:11, c:11}}
{A:{a:1, b:2, c:3}, B:{a:12, b:12, c:12}}
{A:{a:1, b:2, c:3}, B:{a:13, b:13, c:13}}

# flatten recursive
echo '{"A":{"a":1, "b":2, "c":3}, B:[{"d":11, "e":11, "f":11}, {"d":12, "e":12, "f":12}, {"d":13, "e":13, "f":13}]}' \
| flatten -r  -format ltsv
A.a:1, A.b:2, A.c:3, B.a:11, B.b:11, B.c:11
A.a:1, A.b:2, A.c:3, B.a:12, B.b:12, B.c:12
A.a:1, A.b:2, A.c:3, B.a:13, B.b:13, B.c:13
```


### Basic6': Object has array element

```sh
echo '{"A":1, "B":{"d":11, "e":11, "f":[1, 2, 3]}}' | flatten -format json
{A:1, B.d:11, B.e:11, B.f:[1, 2, 3]}

echo '{"A":1, "B":{"d":11, "e":11, "f":[1, 2, 3]}}' | flatten -r -format ltsv
A:1, B.d:11, B.e:11, B.f:1
A:1, B.d:11, B.e:11, B.f:2
A:1, B.d:11, B.e:11, B.f:2
```


### Basic7 Nest Array

```sh
echo {A:1, B:1, C:[{a:11, b:11, c:[111, 112, 113]}, {a:12, b:12, c:[121, 122]}, {a:13, b:13, c:[]}]} \
| flatten -format json
{A:1, B:1, C:{a:11, b:11, c:[111, 112, 113]}}
{A:1, B:1, C:{a:12, b:12, c:[121, 122]}}
{A:1, B:1, C:{a:13, b:13, c:[]}]}

echo {A:1, B:1, C:[{a:11, b:11, c:[111, 112, 113]}, {a:12, b:12, c:[121, 122]}, {a:13, b:13, c:[]}]} \
| flatten -r 1 -format json
{A:1, B:1, C:{a:11, b:11, c:111}}
{A:1, B:1, C:{a:11, b:11, c:112}}
{A:1, B:1, C:{a:11, b:11, c:113}}
{A:1, B:1, C:{a:12, b:12, c:121}}
{A:1, B:1, C:{a:12, b:12, c:122}}
{A:1, B:1, C:{a:13, b:13, c:null}}

echo {A:1, B:1, C:[{a:11, b:11, c:[111, 112, 113]}, {a:12, b:12, c:[121, 122]}, {a:13, b:13, c:[]}]} \
| flatten -r 2 -format ltsv
A:1, B:1, C.a:11, C.b:11, C.c:111
A:1, B:1, C.a:11, C.b:11, C.c:112
A:1, B:1, C.a:11, C.b:11, C.c:113
A:1, B:1, C.a:12, C.b:12, C.c:121
A:1, B:1, C.a:12, C.b:12, C.c:122
A:1, B:1, C.a:13, C.b:13, C.c:null
```


## Advanced Spec

### Multiple object element

```sh
echo '{A:1, B:[1, 2], C:[1, 2, 3]}' | flatten -format ltsv
A:1, B:1, C:1
A:1, B:2, C:2
A:1,      C:3
```

This work process is ...

```sh
# initial step
{A:1, B:[1, 2], C:[1, 2, 3]}

# second step
<-- A*B --> <-- A*C -->
{A:1, B:1}  {A:1, C:1}
{A:1, B:2}  {A:1, C:2}
            {A:1, C:3}

# last step
A:1, B:1, C:1
A:1, B:2, C:2
A:1,      C:3
```

### Multiple object element has array

```sh
echo '{"A":1, "B":{"d":11, "e":[1, 2], "f":[1, 2, 3]}, "C":{"g":12, "h":[1, 2, 3, 4]}}' \
| flatten -fromat ltsv
A:1, B.d:11, B.e:1, B.f:1 C.g:12, C.h:1
A:1, B.d:11, B.e:2, B.f:2 C.g:12, C.h:2
A:1, B.d:11,        B.f:3 C.g:12, C.h:3
A:1, B.d:11,              C.g:12, C.h:4

This work process is...

```sh
# initial step
{"A":1, "B":{"d":11, "e":[1, 2], "f":[1, 2, 3]}, "C":{"g":12, "h":[1, 2, 3, 4]}}

# second step
{A:1, B.d:11, B.e:[1, 2], B.f:[1, 2, 3], C.g:12, C.h:[1, 2, 3, 4]}

# third step
<------------------------>  <------------------------>  <------------------------>
A:1, B.d:11, B.e:1, C.g:12  A:1, B.d:11, B.f:1, C.g:12  A:1, B.d:11, C.g:12, C.h:1
A:1, B.d:11, B.e:2, C.g:12  A:1, B.d:11, B.f:2, C.g:12  A:1, B.d:11, C.g:12, C.h:2
                            A:1, B.d:11, B.f:3, C.g:12  A:1, B.d:11, C.g:12, C.h:3
                                                        A:1, B.d:11, C.g:12, C.h:4

# last step
A:1, B.d:11, B.e:1, B.f:1 C.g:12, C.h:1
A:1, B.d:11, B.e:2, B.f:2 C.g:12, C.h:2
A:1, B.d:11,        B.f:3 C.g:12, C.h:3
A:1, B.d:11,              C.g:12, C.h:4
```

### Array exists in parallel

```sh
echo '{A:1, B:[{a:11, b:11}, {a:12, b:12}, {a:13, b:13}], C:[{c:21, d:21}, {c:22, d:22}]}' \
| flatten -format ltsv
A:1, B.a:11, B.b:11, C.c:21, C.d:21
A:1, B.a:12, B.b:12, C.c:22, C.d:22
A:1, B.a:13, B.b:13, C.c:null, D.e:null
```

This process works is ...

```sh
# first step
{A:1, B:[{a:11, b:11}, {a:12, b:12}, {a:13, b:13}], C:[{c:21, d:21}, {c:22, d:22}]}

# second step
<------------------->  <------------------->
{A:1, B:{a:11, b:11}}  {A:1, C:{c:21, d:21}}
{A:1, B:{a:12, b:12}}  {A:1, C:{c:22, d:22}}
{A:1, B:{a:13, b:13}}  

# last step
A:1, B.a:11, B.b:11, C.c:21, C.d:21
A:1, B.a:12, B.b:12, C.c:22, C.d:22
A:1, B.a:13, B.b:13, C.c:null, D.e:null
```

### Nested array object has array

```sh
echo '{"A":1, "B":[{"a":11, "b":11}, {"a":12, "b":12}, {"a":13, "b":13}], "C":[{"d":21, "f":[1,2,3]}, {"d":22, "f":[1,2]}], "D":[{"d":21, "f":[1,2]}, {"d":22, "f":[1]}]}
' \
| flatten -format ltsv
A:1, B.a:11, B.b:11, C.d:21, C.f:1, D.d:21, D.f:1
A:1, B.a:12, B.b:12, C.d:21, C.f:2, D.d:21, D.f:2
A:1, B.a:13, B.b:13, C.d:21, C.f:3, D.d:22, D.f:1
A:1, B:null,         C.d:22, C.f:1, D:null
A:1, B:null,         C.d:22, C.f:2, D:null
```

This process work is...
```sh
# first step
{
   "A":1,
   "B":[
      {"a":11, "b":11},
      {"a":12, "b":12},
      {"a":13, "b":13}
   ],
   "C":[
      {"d":21, "f":[1, 2, 3 ]},
      {"d":22, "f":[1, 2]}
   ],
   "D":[
      {"d":21, "f":[1, 2]},
      {"d":22, "f":[1]}
   ]
}

# second step
<--------A*B-------->  <------A*C--------->  <------A*D--------->
{A:1, B:{a:11, b:11}}  {A:1, C:{d:21, f:1}}  {A:1, D:{d:21, f:1}}
{A:1, B:{a:12, b:12}}  {A:1, C:{d:21, f:2}}  {A:1, D:{d:21, f:2}}
{A:1, B:{a:13, b:13}}  {A:1, C:{d:21, f:3}}  {A:1, D:{d:22, f:1}}
                       {A:1, C:{d:22, f:2}}

# third step
{A:1, B:{a:11, b:11}, C:{d:21, f:1}, D:{d:21, f:1}}
{A:1, B:{a:12, b:12}, C:{d:21, f:2}, D:{d:21, f:2}}
{A:1, B:{a:13, b:13}, C:{d:21, f:3}, D:{d:22, f:1}}
{A:1, B:null,         C:{d:22, f:1}, D:null}
{A:1, B:null,         C:{d:22, f:2}, D:null}

# last step
A:1, B.a:11, B.b:11, C.d:21, C.f:1, D.d:21, D.f:1
A:1, B.a:12, B.b:12, C.d:21, C.f:2, D.d:21, D.f:2
A:1, B.a:13, B.b:13, C.d:21, C.f:3, D.d:22, D.f:1
A:1, B:null,         C.d:22, C.f:1, D:null
A:1, B:null,         C.d:22, C.f:2, D:null
```
