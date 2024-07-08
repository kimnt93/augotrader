"""
Create Remote class for data service and trading service
Also create a gRPC server to handle the request from client to access trading service and data service
"""
import logging
from concurrent import futures

import grpc

import ray
from dotenv import load_dotenv
import json
from ssi_trading.server import SSIServices
from ssi_trading.utils import run_forever
from ssi_trading.config import TradingServiceConfig, DataServiceConfig
import os
import redis
from ray import serve
from ray.serve.config import gRPCOptions

from app.config.const import (
    DATA_SERVICE_NAME,
    MARKET_DATA_STREAM_NAME,
    INDEX_DATA_STREAM_NAME,
    FOREIGN_DATA_STREAM_NAME, BAR_DATA_STREAM_NAME
)
from app.factory import (
    create_trading_service,
    create_trading_stream,
    create_data_service,
    create_market_data_stream,
    create_index_data_stream,
    create_fr_data_stream, create_bar_data_stream
)
from mproto import service_pb2_grpc

load_dotenv()


REDIS_CONFIG = {
    "host": os.environ['REDIS_HOST']
}
ACCOUNT_CONFIG_PREFIX = os.environ["SV_ACCOUNT_CONFIG_PREFIX"]
DATA_CONFIG_PREFIX = os.environ["SV_DATA_CONFIG_PREFIX"]


@ray.remote
class RemoteSSIService(SSIServices):
    # Do not need to implement this class as it is just a placeholder
    pass


class SSIServiceDeployment:
    pass


@ray.deployment(num_cpus=1)
class GrpcSSIServiceDeployment:
    def __init__(self):
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        service_pb2_grpc.add_TradingServiceServicer_to_server(SSIServiceDeployment(), self.server)
        self.server.add_insecure_port('[::]:50051')
        self.server.start()

    async def __call__(self, starlette_request):
        pass  # Not used since this is for gRPC


def main():
    ray.init(include_dashboard=True)
    serve.start()

    rd_client = redis.Redis(**REDIS_CONFIG)

    # load config
    # data_configs = json.loads(rd_client.get(DATA_CONFIG_PREFIX))
    # RemoteSSIService.remote()

    ssis = RemoteSSIService.options(name="RemoteSSIServiceActor").remote()

    # load all account config
    for account_config_key in rd_client.scan_iter(f"{ACCOUNT_CONFIG_PREFIX}.*"):
        trading_config = json.loads(rd_client.get(account_config_key))
        trading_config = TradingServiceConfig(
            consumer_id=trading_config['consumer_id'],
            consumer_secret=trading_config['consumer_secret'],
            account_id=trading_config['account_id'],
            auth_token=trading_config.get('auth_token', ''),
            paper_trading=trading_config.get('paper_trading', False),
            account_type=trading_config['account_type'],
        )
        # create service
        service = create_trading_service(config=trading_config)
        ssis.add_trading_service.remote(service=service)
        # create stream
        stream = create_trading_stream(config=trading_config)
        ssis.add_trading_steam.remote(stream=stream)

    # add data server
    cfg: bytearray = rd_client.get(f"{DATA_CONFIG_PREFIX}.{DATA_SERVICE_NAME}")
    if cfg:
        data_service_config = json.loads(cfg.decode())
        data_service_config = DataServiceConfig(
            consumer_id=data_service_config['consumer_id'],
            consumer_secret=data_service_config['consumer_secret'],
            symbols=None
        )
        service = create_data_service(config=data_service_config)
        ssis.add_data_service.remote(service=service)

    # add data streams
    # market stream
    cfg: bytearray = rd_client.get(f"{DATA_CONFIG_PREFIX}.{MARKET_DATA_STREAM_NAME}")
    if cfg:
        market_stream_config = json.loads(cfg.decode())
        market_stream_config = DataServiceConfig(
            consumer_id=market_stream_config['consumer_id'],
            consumer_secret=market_stream_config['consumer_secret'],
            symbols=market_stream_config['symbols']
        )
        stream = create_market_data_stream(config=market_stream_config)
        ssis.add_data_stream.remote(stream=stream)

    # add index stream
    cfg: bytearray = rd_client.get(f"{DATA_CONFIG_PREFIX}.{INDEX_DATA_STREAM_NAME}")
    if cfg:
        index_stream_config = json.loads(cfg.decode())
        index_stream_config = DataServiceConfig(
            consumer_id=index_stream_config['consumer_id'],
            consumer_secret=index_stream_config['consumer_secret'],
            symbols=index_stream_config['symbols']
        )
        stream = create_index_data_stream(config=index_stream_config)
        ssis.add_data_stream.remote(stream=stream)

    # add fr stream
    cfg: bytearray = rd_client.get(f"{DATA_CONFIG_PREFIX}.{FOREIGN_DATA_STREAM_NAME}")
    if cfg:
        fr_stream_config = json.loads(cfg.decode())
        fr_stream_config = DataServiceConfig(
            consumer_id=fr_stream_config['consumer_id'],
            consumer_secret=fr_stream_config['consumer_secret'],
            symbols=fr_stream_config['symbols']
        )
        stream = create_fr_data_stream(config=fr_stream_config)
        ssis.add_data_stream.remote(stream=stream)

    # start bar
    cfg: bytearray = rd_client.get(f"{DATA_CONFIG_PREFIX}.{BAR_DATA_STREAM_NAME}")
    if cfg:
        bar_stream_config = json.loads(cfg.decode())
        bar_stream_config = DataServiceConfig(
            consumer_id=bar_stream_config['consumer_id'],
            consumer_secret=bar_stream_config['consumer_secret'],
            symbols=bar_stream_config['symbols']
        )
        stream = create_bar_data_stream(config=bar_stream_config)
        ssis.add_data_stream.remote(stream=stream)

    logging.info("Starting data stream")
    # start stream nè
    ssis.start_data_stream.remote()
    ssis.start_trading_stream.remote()

    g = GrpcSSIServiceDeployment.bind()
    app1 = "app1"
    serve.run(target=g, name=app1, route_prefix=f"/{app1}")
    run_forever(15)


# Main program to use the remote class
if __name__ == "__main__":
    # Create an instance of RemoteSSIService
    main()
