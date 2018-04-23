#ifndef TLS_CLIENT_H
#define TLS_CLIENT_H

#include <string>
#include <vector>

#include "mbedtls/net.h"
#include "mbedtls/ssl.h"
#include "mbedtls/entropy.h"
#include "mbedtls/ctr_drbg.h"
#include "mbedtls/debug.h"

namespace https
{

const string MethodGet = "GET";

class Client
{
public:
  Client();
  ~Client();

  void get(string url) : Do(Request(MethodGet, url));

private:
  void Do(Request &request);

  mbedtls_net_context server_fd;
  mbedtls_entropy_context entropy;
  mbedtls_ctr_drbg_context ctr_drbg;
  mbedtls_ssl_context ssl;
  mbedtls_ssl_config conf;
  mbedtls_x509_crt cacert;
};

class Request
{
public:
  Request(string method, string url);
  string AsMessage();

private:
  const string method;

  // @TODO (rgeraldes) replace with url::URL url;

  const string url;
  const string hostname;

  // @TODO (rgeraldes) - add
  //const Header header;
};

/*
class Header
{
private:
  const vector<string> values;

public:
  void add();
};
*/

class Response
{
private:
  const string content;

public:
  const string &getContent() const
  {
    return content;
  }
};
}

#endif
