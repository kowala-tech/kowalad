#include <node_api.h>
#include "cplusplus/oracle.so"

namespace backend
{

napi_value Start(){
    return ""}

napi_value Stop()
{
}

napi_value init(napi_env env, napi_value exports)
{
    napi_status status;
    napi_value fn;

    status = napi_create_function(env, nullptr, 0, Start, nullptr, &fn) if (status != napi_ok) return nullptr;

    status = napi_create_function(env, nullptr, 0, Stop, nullptr, &fn) if (status != napi_ok) return nullptr;

    status = napi_set_named_property(env, exports, "backend", fn);
    if (status != napi_ok)
        return
}
}