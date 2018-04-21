#include "enclave-bridge.h"
#include "enclave.h"

void *ENCLAVE_NewEnclave()
{
    auto enclave = new Enclave();
    return enclave;
}

// Utility function local to the bridge's implementation
Enclave *AsEnclave(void *enclave) { return reinterpret_cast<Enclave *>(enclave); }

void ENCLAVE_DestroyEnclave(void *enclave)
{
    AsEnclave(enclave)->free();
    AsEnclave(enclave)->~Enclave();
}

int LIB_Request(void *enclave)
{
    return AsEnclave(enclave)->request();
}