This makes generating lorem ipsum a snap for your project.

=============

Usage
-----
import "lorem"


Exposed functions
-----------------
Ranged generators


These will generate a string with a variable number 
of elements specified by a range you provide

* lorem.Word(min, max int) string
* lorem.Sentence(min, max int) string
* lorem.Paragraph(min, max int) string

Convenience functions
---------------------
* lorem.Host() string
* lorem.Email() string
* lorem.Url() string


