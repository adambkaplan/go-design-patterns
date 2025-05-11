# Decorator Pattern

Developers often call the Decorator pattern "object wrapping," where a class/object extends
functionality by including a field with the same interface/supertype. In this contrived example,
a hypothetical "beverage" is decorated with the extras a customer may add (milk, whipped cream,
etc.).

The implementation takes advantage of golang struct embedding to extend the capabilities of the
"Condiment" decorators. Unlike Java or other object-oriented languages, effective struct embedding
requires the "base" struct to be exported and useful.
