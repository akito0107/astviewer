#!/bin/sh

tools/goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("./public")))'
