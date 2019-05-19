# Keypair - Go convenience wrapper for a private/public RSA keypair.

A Go convenience wrapper which provides access to a private/public RSA keypair, which is located in either a Google Cloud Platform Storage bucket or local storage.

## What?
This provides synchronised access to private and public RSA keys for encrypt/decrypt operations.

## Why?
Written to provide convenient encryption/decryption, and synchronised access to both keys so that they can be passed around.

## How?
The best place to start is with the tests. If running locally, then ensure that  [Google Application Credentials] have been created. If running from a [GCP] virtual machine, then ensure that the relevant service account (compute, appengine etc.) has the following IAM scopes: 'Storage Object Viewer' and 'Storage Object Creator', or 'Storage Object Admin'. See [GCP service accounts] for further details.

## Examples
See the tests for usage examples.

## Dependencies and services
This utilises the following fine pieces of work:
* [Dave Grijalva]'s [jwt-go] Go implementation of JSON Web Tokens (JWT)
* [GCP]'s [Storage Go client]

## Installation
Install using go get.

```sh
$ go get -u github.com/lidstromberg/keypair
```
#### Environment Variables
You will also need to export (linux/macOS) or create (Windows) some environment variables.
```sh
################################
# KEYPAIR
################################
export KP_TYPE="bucket"
export KP_GCP_BUCKET="{{BUCKETNAME}}"
export KP_PRIKEY="jwt.key"
export KP_PRIKEYP="{{PRIVATEKEYPASSWORD}}"
export KP_PUBKEY="jwt.key.pub"
```
```sh
################################
# GCP CREDENTIALS
################################
export GOOGLE_APPLICATION_CREDENTIALS="/PATH/TO/GCPCREDENTIALS.JSON"
```
(See [Google Application Credentials])

#### Private/Public Certs
The following will generate RSA private/public keys (assuming you have openssl installed). You should place a password on the private key when prompted.

```sh
$ ssh-keygen -t rsa -b 4096 -f jwt.key
$ openssl rsa -in jwt.key -pubout -outform PEM -pubout -out jwt.key.pub
```

#### Google Cloud Platform Requirements
If you intend to use GCP datastore as your backend, then you will require:
* A GCP project
* A GCP storage bucket (private) to store the RSA private/public keys (in the root of the bucket)
* Your GOOGLE_APPLICATION_CREDENTIALS json credentials key should be created with the following IAM scopes: 'Storage Object Viewer' and 'Storage Object Creator', or 'Storage Object Admin'.


### Main Files
| File | Purpose |
| ------ | ------ |
| keypair.go | Logic manager |
| keypair_test.go | Tests |

### Ancillary Files
| File | Purpose |
| ------ | ------ |
| config.go | Boot package parameters, environment var collection |
| errors.go | Package error definitions |
| env | Package environment variables for local/dev installation |
| gogets | Statements for go-getting required packages |


   [Dave Grijalva]: <https://github.com/dgrijalva>
   [jwt-go]: <https://github.com/dgrijalva/jwt-go>
   [GCP]: <https://cloud.google.com/>
   [Storage Go client]: <https://cloud.google.com/storage/docs/reference/libraries#client-libraries-install-go>
   [Google Application Credentials]: <https://cloud.google.com/docs/authentication/production#auth-cloud-implicit-go>
   [examples]: <https://github.com/lidstromberg/examples>
