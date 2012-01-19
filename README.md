it's not ready yet.

This is here primarily for backup.  It may turn into something useful.
I'm exploring how django-like models might be crafted in against mongo

=============

Pillars
-------

* No Magic
* No Reflect
* No Unsafe
* As functional as feasible, i.e. defering mutable state until optimization phase


Dependecies
-----------

The dependencies listed are required if you wish to use SansMagic

* [.golang](http://golang.org/doc/install.html) 
* goinstall launchpad.net/gobson/bson
* goinstall launchpad.net/gobson/mgo



### Getting Started

-not yet-

$ goinstall sansmagic
$ sans new-proj sans-hello
$ cd sans-hello
$ make run


Installation
-----------

    goinstall sansmagic

if this doesn't work then you probably have an old go install (that is of course
if my code didn't break it!)

Example
-----


Testing
-------

Go to the projects root directory and

    $ ./test

To add tests see the `Commands` section earlier in this
README.

Contributing
------------

1. Fork it.
2. Create a branch (`git checkout -b my_branch`)
3. Commit your changes (`git commit -am "recoded the whole thing"`)
4. Push to the branch (`git push origin my_branch`)
5. Create an [Issue][1] with a link to your branch
6. Prepare for eternal fame and fortune.

[1]: http://github.com/drhodes/sansmagic/issues
