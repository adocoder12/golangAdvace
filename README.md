# golangAdvace

<h3> Embedding </h2>

<p>
    In Go, embedding means putting one struct (a container for variables), inside another struct without giving it a field name.
</p>

<p>
   Example from your code:

    type App struct {
    dbmock.Database // This is embedding!
    }

</p>
