#include "Enclave_t.h"

int ecall_currency_value(uint8_t* raw_tx, size_t* raw_tx_len) {
    return 0;
}

int ecall_create_sealed_account(uint8_t* sealed_account, uint32_t sealed_account_size){
    return 0;
}

int ecall_delete_sealed_account(const uint8_t* sealed_log, uint32_t sealed_account_size){
    return 0;
}

int ecall_select_sealed_account(const uint8_t* sealed_account, uint32_t sealed_account_size){
    return 0;
}

int ecall_register_oracle([user_check] uint8_t* raw_tx, [out] size_t* raw_tx_size){
    return 0;
}

int ecall_deregister_oracle([user_check] uint8_t* raw_tx, [out] size_t* raw_tx_size){
    return 0;
}
        
int ecall_redeem_oracle_funds([user_check] uint8_t* raw_tx, [out] size_t* raw_tx_size){
    return 0;
}


/*


#include <vector>
#include <thread>

using namespace std;

vector<Scraper> exchanges{Coinbase(), CoinMarketCap()};

#include "scraper.h"
#include "sources/coinbase.h"
#include "sources/coinmarketcap.h"
#include "kowala.h"

try {
        return currency_value(raw_tx, len);
    }
    catch(){}

    // @TODO (rgeraldes) - return error
    return 0

// @TODO (rgeraldes) - valid price intervals
int currency_value(uint8_t* raw_tx, size_t* len) {
    len = exchanges.size();
    vector<u256> priceRegistry(len, 0);
    spawnRequests(len, priceRegistry);
    
    // average value based on valid inputs 
    int valid = 0;
    u256 sum = 0;
    int i = 0;
    for (it = prices.begin(); it != exchanges.end(); it++, i++) {
        u256 price = priceRegistry[i];
        
        if (price == 0) {
            continue;
        }

        valid++;
        sum += price;
    }
    u256 final_price = sum/valid;

    // ABI encoding data


    // create transaction
    kowala::Transaction tx(u256 const &_value, u256 const &_gasPrice, u256 const &_gas, Address const &_dest, bytes const &_data, u256 const &_nonce, Secret const &_secret)
    bytes tx_rlp = tx.rlp();
    memcpy(tx_raw, &tx_rlp[0], rx_rlp.size());
    *raw_tx_len = out.size();

}

func spawnRequests(int len, vector<u256> priceRegistry) {
    thread threads[len];
    for (int = 0; i < len; i++) {
        threads[i] = thread(exchanges[i].run(&priceRegistry[i]));
    }
    for (auto& th : threads) {
        th.join();
    }
}




int request()
{
    try
    {
        // fetch prices // one thread per exchange

        // validate prices

        //

        // create a transaction with the average price depending on the volume per exchange
    }
    catch (const std::exception &e)
    {
    }
}
*/