from ssi_trading.services.client import BaseTradingService, BaseDataService

from ssi_trading.services.client.fundamental import FundamentalTradingService
from ssi_trading.services.client.futures import FutureTradingService
from ssi_trading.services.paper.fundamental import PaperFundamentalTradingService
from ssi_trading.services.paper.futures import PaperFutureTradingService
from ssi_trading.services.client.data import MarketDataService

from ssi_trading.config import TradingServiceConfig, DataServiceConfig

# import all streams
from ssi_trading.services.stream import BaseDataStream, BaseTradingStream

# data stream
from ssi_trading.services.stream.fr import ForeignRoomDataStream
from ssi_trading.services.stream.market import MarketDataStream
from ssi_trading.services.stream.bar import BarDataStream
from ssi_trading.services.stream.index import IndexDataStream

# trading stream: account, order, portfolio, trading, etc.
from ssi_trading.services.stream.trading import TradingStream


def create_trading_service(config: TradingServiceConfig) -> BaseTradingService:
    if config.account_type == 'fundamental':
        return FundamentalTradingService(config) if not config.paper_trading else PaperFundamentalTradingService(config)
    elif config.account_type == 'future':
        return FutureTradingService(config) if not config.paper_trading else PaperFutureTradingService(config)
    else:
        raise ValueError('Invalid account type')


def create_data_service(config: DataServiceConfig) -> BaseDataService:
    return MarketDataService(config)


def create_market_data_stream(config: DataServiceConfig) -> BaseDataStream:
    return MarketDataStream(config)


def create_index_data_stream(config: DataServiceConfig) -> BaseDataStream:
    return IndexDataStream(config)


def create_fr_data_stream(config: DataServiceConfig) -> BaseDataStream:
    return ForeignRoomDataStream(config)


def create_bar_data_stream(config: DataServiceConfig) -> BaseDataStream:
    return BarDataStream(config)


def create_trading_stream(config: TradingServiceConfig) -> BaseTradingStream:
    return TradingStream(config)
