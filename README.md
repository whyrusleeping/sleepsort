Sleepsort
=========

The super amazing sleep sort. For those of you who want an O(n) sort with no side effects!

It is almost guaranteed not to take longer than the largest number in your list of milliseconds.

The implementation is beautiful, Go supports easily spawning off coroutines that can
wait until it is their turn to enter the list, therefore it is ideal for making a super
multithreaded sorting algorithm such as this!

Unfortunately, it can only sort integers at the moment, but as soon as I acquire a new flux capacitor,
capabilities for sorting floating point numbers may be added.
