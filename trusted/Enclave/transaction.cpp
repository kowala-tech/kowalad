#include "transaction.h"

void Transaction::streamRLP(RLPStream &_s, IncludeSignature _sig) const
{
    if (m_type == NullTransaction)
        return;

    _s.appendList((_sig || _forEip155hash ? 3 : 0) + 6);
    _s << m_nonce << m_gasPrice << m_gas;
    if (m_type == MessageCall)
        _s << m_receiveAddress;
    else
        _s << "";
    _s << m_value << m_data;

    if (_sig)
    {
        if (!m_vrs)
            BOOST_THROW_EXCEPTION(TransactionIsUnsigned());

        if (hasZeroSignature())
            _s << m_chainId;
        else
        {
            int const vOffset = m_chainId * 2 + 35;
            _s << (m_vrs->v + vOffset);
        }
        _s << (u256)m_vrs->r << (u256)m_vrs->s;
    }
    else if (_forEip155hash)
        _s << m_chainId << 0 << 0;
}