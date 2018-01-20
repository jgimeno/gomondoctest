Normally it is annoying to do unit test when mongo is involved, mocking
all the dependencies involved is practically imposible. The other solution
is using a real mongodb, this comes to solve that.

From your own test you can download a mongo image, launch it and it will
be ready in your localhost.

Prerequisites

- Install docker.

To use it is as easy as in your tests do:

```
func TestMyMongoCode(t *testing.T) {
    gm := gomondoctest.NewGomondoc(t)

    gm.RunMongo()
    defer gm.StopMongo()

    mgo.Dial("mongodb://localhost")

    ....
}
```

This is a little library inspired by this post https://developers.almamedia.fi/2014/painless-mongodb-testing-with-docker-and-golang/

Feel free to contribute!
