[![Gitter chat](https://badges.gitter.im/kowala/kcoin.png)](https://gitter.im/kowala-tech/lobby)

# oracled

oracled is currently an underlying part of the Kowala's Oracle System. It simplifies communication with Kowala's blockchain and management of the authenticated data feed system. oracled's architecture is as follows:

                           +---------ORACLED---------+
                           |                         |
                           |                         |
                           | +---------+ +---------+ |
                           | |         | |  CRON   | |
                           | |  LIGHT<----+   +    | |  

SUBMITS AUTHENTICATED DATA | | CLIENT | | +--v--+ | | (TLS) REQUESTS DATA  
 <-----------------+ | | | SGX +------------------>
| | | | +-----+ | |
| +---------+ +---------+ |
+-------------------------+

## Build

* Generate a native node.js module for [Kowala's oracle app](https://github.com/kowala-tech/oracle).
* Standalone deamon:
  * source $SGX_SDK/environment # not needed when you already have it
  * `make oracled`

## Core Contributors

[Core Team Members](https://github.com/orgs/kowala-tech/people)

## Contact us

Feel free to email us at support@kowala.tech or talk to us on [Gitter](https://gitter.im/kowala-tech/lobby).
