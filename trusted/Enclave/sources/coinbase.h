#ifndef TRUSTED_ENCLAVE_SCRAPERS_COINBASE_H_
#define TRUSTED_ENCLAVE_SCRAPERS_COINBASE_H_

#include "Scraper.h"

class CoinbaseScraper : Scraper
{
public:
  void run();

private:
  // TODO (rgeraldes) : apikey and secret
};

#endif