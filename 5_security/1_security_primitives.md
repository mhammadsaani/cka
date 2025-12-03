Basics
- All access must be secured
- root access disabled
- password based auth disabled
- only ssh key based

In k8s, as api server is the point of contact so we need to secure it. Here, securing means
- Who can access it
- What actions can it perform

Who can access can be configured via following [authentication]
- Static Token File
- Certificates
- External Auth Providers - LDAP
- Service Accounts

What can. they do is defined by [authorization]
- RBAC
- ABAC
- Node authorization
- Webhook

All communication between various components of  clusters is secured by tls

By default, all pods can access other pods but. access can be restricted via network policies.



