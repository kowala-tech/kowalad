#include "rlp.h"
#include "common.h"

enum IncludeSignature
{
    WithoutSignature = 0,
    WithSignature = 1,
};

class Transaction
{
  public:
    /// Constructs a signed message-call transaction.
    Transaction(u256 const &_value, u256 const &_gasPrice, u256 const &_gas, Address const &_dest, bytes const &_data, u256 const &_nonce, Secret const &_secret) : m_type(MessageCall), m_nonce(_nonce), m_value(_value), m_receiveAddress(_dest), m_gasPrice(_gasPrice), m_gas(_gas), m_data(_data) { sign(_secret); }

    /// @returns the RLP serialisation of this transaction.
    bytes rlp(IncludeSignature _sig = WithSignature) const
    {
        RLPStream s;
        streamRLP(s, _sig);
        return s.out();
    }

    /// Serialises this transaction to an RLPStream.
    /// @throws TransactionIsUnsigned if including signature was requested but it was not initialized
    void streamRLP(RLPStream &_s, IncludeSignature _sig = WithSignature) const;
};