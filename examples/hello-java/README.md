# mcandre/java/hello

Hello World in Java

# EXAMPLE

```
$ gradle shadowJar
$ bin/hello
Hello World
```

# REQUIREMENTS

* [JDK](http://www.oracle.com/technetwork/java/javase/downloads/index.html) 1.7+
* [Gradle](http://gradle.org/) 2.1+

## Optional

* [Sonar](http://www.sonarqube.org/)
* [Infer](http://fbinfer.com/)

E.g., `brew install gradle sonar sonar-runner infer`

# JAVADOCS

```
$ gradle javadoc
$ open build/docs/javadoc/index.html
```

# TEST + CODE COVERAGE

```
$ gradle test jacoco
$ open build/reports/jacoco/test/html/index.html
```

# LINTING

```
$ gradle check
```

## Optional: FindBugs

```
$ gradle check
$ open build/reports/findbugs/main.html
```

## Optional: Sonar

```
$ sonar start
$ gradle check sonar
$ open http://localhost:9000/
```

## Optional: Infer

```
$ infer -- gradle clean build
```
