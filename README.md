# aws-docker-login

A small binary to get the ECR credentials for AWS.

Use the normal aws configuration settings for this.

## Why?

Using the python API in a docker image is nice, but it's *really* large, especially
when you only need a single part of the API. Compared to a virtual env/python install
coming in at ~400+mb of extra space for a virtual container, this adds \~7.6mb, and with
`upx` it gets down to around 3mb. Yeah, I could just use a shell script or whatever but
even that would require cURL and other dependencies the container may not want.
