#include "coinbase.h"
#include "http.h"
#include "external/picojson/picojson.h"

void CoinbaseScraper::run(u256 *resp)
{
    http::Client client;
    u256 final_price = 0;

    try
    {
        http::Response ret = client.Get("https://api.coinbase.com/v2/prices/BTC-USD/spot");
        picojson::value resp_data;
        string err = picojson::parse(resp_data, ret.getContent());
        if (!err.empty())
        {
            // @TODO (rgeraldes)
        }

        // @TODO (rgeraldes) - add check for the currency .contains("currency") and is == to "USD"
        if (resp_data.is<picojson::array>() && resp_data.get<picojson::array>().size() == 1 && resp_data.get<picojson::array>()[0].contains("amount") && resp_data.get<picojson::array>()[0].get("amount").is<string>())
        {
            final_price = resp_Data.get<picojson::array>()[0].get("amount").get<string>();
        }
        else
        {
            // @TODO (rgeraldes)
        }
        catch ()
        {
        }
    }

    *resp = final_price;
}