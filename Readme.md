
# Table of Contents

1.  [Overview](#orgf0e91ed)
2.  [CLI usage](#orga7e5f16)
    1.  [Installation](#org6a69f21)
    2.  [Usage](#org08c9dae)
        1.  [Command line options](#org29fd117)
3.  [Library usage](#orgb435ef2)
    1.  [Getting started](#orgec6fe55)


<a id="orgf0e91ed"></a>

# Overview

This respository provides both a CLI tool for querying Dehashed, and a library which may be by other projects to query the Dehashed API natively within Go. 


<a id="orga7e5f16"></a>

# CLI usage


<a id="org6a69f21"></a>

## Installation

To install the CLI tool, simply run the following command. Note that you must have a $GOPATH set. 

    go get gitlab.com/cgboal/dehashed


<a id="org08c9dae"></a>

## Usage

The CLI tool requires both a username and an API key to function. Both of these pieces of information are specified via environment variables.  

These environment variables can be set by exporting the appropriate key-value pairs, e.g. 

    export DEHASHED_USERNAME=example.user@example.org
    export DEHASHED_API_KEY=XXXXXXXXXXXXXXXXXXXXXX

These lines can be added to your ~/.zshrc, or ~/.bashrc files for persistence. 


<a id="org29fd117"></a>

### Command line options

By default the CLI tool will output in the format  email:password. Supplying the flag `-oJ` will output in JSON instead. Furthermore, the CLI tool will only output results where a cleartext password was found by defualt, to view all results, use the `-all` flag.


<a id="orgb435ef2"></a>

# Library usage


<a id="orgec6fe55"></a>

## Getting started

To import the library, add the following url to your imports:

`gitlab.com/cgboal/dehashed/lib`

Once imported, you can query the dehashed API as follows: 

    query := "onsecurity.co.uk"
    
    results := dehashed.FetchAll(query)
    
    for _, result := range results {
        fmt.Printf("%s:%s", result.Email, result.Password)
    }

The struct used to represent dehashed results is as follows: 

    type Entry struct {
      Email string `json:"email"`
      Username string `json:"username"`
      Password string `json:"password"`
      Hash string `json:"hashed_password"`
      Name string `json:"name"`
      Source string `json:"obtained_from"`
    }

