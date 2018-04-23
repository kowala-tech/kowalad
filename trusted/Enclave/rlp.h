
class RLPStream
{
  public:
    RLPStream() {}

    /// Read the byte stream.
    bytes const &out() const
    {
        if (!m_listStack.empty())
            BOOST_THROW_EXCEPTION(RLPException() << errinfo_comment("listStack is not empty"));
        return m_out;
    }

  private:
    bytes m_out;
};