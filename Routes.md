## [sdkman-candidates](https://github.com/sdkman/sdkman-candidates/blob/master/conf/routes)

* /candidates
    * /alive
    * /default/:candidate
    * /validate/:candidate/:version/:platform
    * /all
    * /list
    * /:candidate/:platform/versions/all
    * /:candidate/:platform/versions/list
    
## [sdkman-broker](https://github.com/sdkman/sdkman-candidates/blob/master/conf/routes)

* /broker
    * /health/:name?
    * /version
    * /download/sdkman/version/:versionType
    * /download/sdkman/:command/:version/:platform
    * /download/:candidate/:version
    * /download/:candidate/:version/:platform

## [sdkman-broadcast](https://github.com/sdkman/sdkman-broadcast/blob/master/src/main/groovy/io/sdkman/controller/BroadcastController.groovy)

* /broadcast
    * /latest/id
    * /latest
    * /:id
    
## [sdkman-hooks](https://github.com/sdkman/sdkman-hooks/blob/master/conf/routes)

* /alive
* /install
* /selfupdate
* hooks/:phase/:candidate/:version/:platform