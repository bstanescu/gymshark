# Software Engineering Challenge

Imagine for a moment that one of our product lines ships in various pack sizes:
- 250 Items
- 500 Items
- 1000 Items
- 2000 Items
- 5000 Items

Our customers can order any number of these items through our website, but they will
always only be given complete packs.

1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out no more items than necessary to
fulfil the order.
3. Within the constraints of Rules 1 &amp; 2 above, send out as few packs as possible to
fulfil each order.

Examples:

| Items ordered | Correct number of packs         | Incorrect number of packs                                                 |
|---------------|------------------------------|---------------------------------------------------------------------------|
| 1             | 1 x 250                      | 1 x 500 – more items than necessary                                       |
| 250           | 1 x 250                      | 1 x 500 – more items than necessary                                       |
| 251           | 1 x 500                      | 2 x 250 – more packs than necessary                                       |
| 501           | 1 x 500<br /> 1 x 250            | 1 x 1000 – more items than necessary <br /> 3 x 250 – more packs than necessary |
| 12001         | 2 x 5000<br /> 1 x 2000<br /> 1 x 250 | 3 x 5000 – more items than necessary                                      |

Write an application that can calculate the number of packs we need to ship to the customer.
The API must be written in Golang &amp; be usable by a HTTP API (by whichever method you
choose).

## Optional:
- Keep your application flexible so that pack sizes can be changed and added and
removed without having to change the code.
- Create a UI to interact with your API

# Current implementation

The [current implementation](main.go) uses dynamic programming to reduce the complexity to O(k*n)

It doesnt implement the API endpoints, in favour of:
- writing a quick interface that allows updating the packages too (they will persist on the server side)
- writing [tests](main_test.go)
