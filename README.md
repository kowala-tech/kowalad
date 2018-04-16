[![Gitter chat](https://badges.gitter.im/kowala/kcoin.png)](https://gitter.im/kowala-tech/lobby)

# kowalad

kowalad is an experimental middleware/deamon that simplifies communication with Kowala's blockchains.  
kowalad is currently an underlying part of the Kowala's Oracle - authenticated data feed system.  
Account management is not covered in this module as we will use an Intel SGX Enclave for key management.

## Build

* Generate an experimental library which provides a C++ wrapper over a light client node.
  * dynamic library - `make kowalad-dynamic`
  * static library - `make kowalad-static`
* Standalone server - `make kowalad`

## Core Contributors

[Core Team Members](https://github.com/orgs/kowala-tech/people)

## Contact us

Feel free to email us at support@kowala.tech or talk to us on [Gitter](https://gitter.im/kowala-tech/lobby).
