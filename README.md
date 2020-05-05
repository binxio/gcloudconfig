# gcloud-config-helper
Google Cloud SDK Credentials helper

The standard Google Cloud Platform SDK for Go ignores the gcloud configuration for 
  default project and account to use the authentication with.
  
The function `GetCredentials` returns a `google.Credentials` object based upon the active
  configuration. This allows you to create utilities that will work in the same context as gcloud.

The function `GetConfig` will return the properties of the active configuration in gcloud. 

The implementation entirely depends on the output of `gcloud config config-helper`, which is intended to 
  utilities to use the gcloud configuration. Oddly enough, Google also promises it to be a volatile and
  unstable interface. So if Google broke it or changed it, let me know perhaps we can fix the implementation
  of this library :-)
