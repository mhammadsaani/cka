Who can access can be configured via following [authentication]
- Static Token File
- Certificates
- External Auth Providers - LDAP


Static Token is like generating a Bearer token and then using that to get access.

NOTE
Note: This is not recommended in a production environment. 
This is only for learning purposes. Also note that this approach is deprecated in Kubernetes version 1.19
and is no longer available in later releases



kubectl get pods --token-auth-file=user-token-details.csv. Not recommended as it is insecure.



