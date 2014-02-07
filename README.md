GoMobileDetect
==============

[![Build Status](https://travis-ci.org/Shaked/gomobiledetect.png?branch=master)](https://travis-ci.org/Shaked/gomobiledetect)

### Description

GoMobileDetect is a lightweight package imported from PHP for detecting mobile devices including tablets. 

The package is imported from [MobileDetect](http://www.mobiledetect.net) which was originally written in PHP.

### Installation 

    $ go get https://github.com/Shaked/gomobiledetect 

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
        deviceProperty := "iPhone"
        if detect.VersionFloat(deviceProperty) > 6 { 
            // do something with iPhone v6 
        } 
    }

### License

Go Mobile Detect is an open-source script released under [MIT License](http://www.opensource.org/licenses/mit-license.php). 
