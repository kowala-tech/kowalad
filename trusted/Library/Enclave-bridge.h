#pragma once
#ifdef __cplusplus
extern "C" {
#endif

void *ENCLAVE_NewEnclave(int value);
void ENCLAVE_DestroyEnclave(void *enclave);
int ENCLAVE_Request(void *enclave);

#ifdef __cplusplus
} // extern "C"
#endif