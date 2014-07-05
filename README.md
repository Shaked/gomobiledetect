GoMobileDetect
==============

[![GoDoc](https://godoc.org/github.com/Shaked/gomobiledetect?status.png)](https://godoc.org/github.com/Shaked/gomobiledetect)
[![Build Status](https://travis-ci.org/Shaked/gomobiledetect.png?branch=master)](https://travis-ci.org/Shaked/gomobiledetect)
[![Coverage Status](https://coveralls.io/repos/Shaked/gomobiledetect/badge.png)](https://coveralls.io/r/Shaked/gomobiledetect)

### Description

GoMobileDetect is a lightweight package imported from PHP for detecting mobile devices including tablets. 

The package is imported from [MobileDetect](http://www.mobiledetect.net) which was originally written in PHP.

### Installation 

    $ go get github.com/Shaked/gomobiledetect 

### Updates 

- This version supports a new method ```IsKey(key int)```. This method will replace the ```Is(key string)``` method in the future. All keys can be found [here](https://github.com/Shaked/gomobiledetect/blob/maps-to-lists/rules.go#L4)

- The ```Is(key string)``` has been changed to support both ```string``` and ```int``` using the ```interface{}```. This support **will be removed in the future**

### Usage

    import "github.com/Shaked/gomobiledetect"
    //code here 
    func handler(w http.ResponseWriter, r *http.Request) {
        detect := gomobiledetect.NewMobileDetect(r, nil)
        if detect.IsMobile() { 
            // do some mobile stuff 
        }
        if detect.IsTablet() {
            // do some tablet stuff 
        }
        
        if detect.IsKey(gomobiledetect.IPHONE) { 
            // do something with iPhone
        }
        
        //Deprecated, still works but will be removed in the near future
        if detect.Is("iphone") { 
            // do something with iPhone
        }
        
        deviceProperty := "iPhone"
        if detect.VersionFloat(deviceProperty) > 6 { 
            // do something with iPhone v6 
        } 
    }

### License

Go Mobile Detect is an open-source script released under [MIT License](http://www.opensource.org/licenses/mit-license.php). 
