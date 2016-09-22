# High-performance API example built with Revel and Golang

### Start the web server:

    revel run github.com/jongensvantechniek/revel-example

   Run with <tt>--help</tt> for options.

### Configuration
The configuration is based on environment variables:
```
#!/usr/bin/bash
export APP_DATABASE_HOST="" #PostgreSQL Database Host
export APP_DATABASE_USER="" #PostgreSQL Database User
export APP_DATABASE_PASS="" #PostgreSQL Database Pass
export APP_DATABASE_NAME="" #PostgreSQL Database Name
```

### Go to http://localhost:9000/ and you'll see:

The results of a query on the user table of a PostgreSQL database.

```
CREATE TABLE users (
    username VARCHAR(60)
)

INSERT INTO users(username) VALUES('user1'),('user2'),('user3')
```

### Description of Contents

The default directory structure of a generated Revel application:

    revel-example       App root
      app               App sources
        controllers     App controllers
          init.go       Interceptor registration
        models          App domain models
        routes          Reverse routes (generated code)
        views           Templates
      tests             Test suites
      conf              Configuration files
        app.conf        Main configuration file
        routes          Routes definition
      messages          Message files

app

    The app directory contains the source code and templates for your application.

conf

    The conf directory contains the application’s configuration files. There are two main configuration files:

    * app.conf, the main configuration file for the application, which contains standard configuration parameters
    * routes, the routes definition file.

messages

    The messages directory contains all localized message files.

test

    Tests are kept in the tests directory. Revel provides a testing framework that makes it easy to write and run functional tests against your application.
