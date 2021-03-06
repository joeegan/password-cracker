Display the time taken to crack the supplied password

```bash
> ./password-cracker -password=1234
Searched 435 of 3000 surnames
Searched 299 of 3000 firstnames
Found password '1234' at index 4 of 3000 common passwords
Time taken to crack 244.288µs
```
**PLAN**

Simplistic approach:
Range to be hardcoded to between 5 and 10 chars for now, alphanumeric, upper and lower cases permitted.

* A. Looks up a list of most common passwords. (e.g Password1)
* B. Looks up common names list. (e.g. joe, bloggs)
* C. Looks up english dictionary. (e.g. a, at)
* D. Try first letter uppercase on A, B, C lists.
* E. Try numbers 1-9 then 0 appended to lists
* F. Try E with D
* G. Use combination-lock style approach with numbers and letters

**Next steps**
* Run concurrent list searches
* Expose as a service
* Build web client which displays data and uses 'time to crack' ranges to convey password strength
* Host using AWS
