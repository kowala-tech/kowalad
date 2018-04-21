#pragma once

class Enclave
{
public:
  Enclave();
  ~Enclave();
  void request() const;

private:
  sgx_enclave_id_t global_eid = 0;
};