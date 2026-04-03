import pytest
import grpc
from proto import booking_pb2, booking_pb2_grpc

GRPC_ADDR = "localhost:50051"

def test_grpc_connection():
    try:
        with grpc.insecure_channel(GRPC_ADDR) as channel:
            stub = booking_pb2_grpc.BookingServiceStub(channel)
            
            response = stub.CheckTable(
                booking_pb2.TableRequest(id=5),
                timeout=2 
            )
            
            assert response is not None
            assert isinstance(response.available, bool)
            
    except grpc.RpcError as e:
        pytest.fail(f"gRPC сервер недоступен на {GRPC_ADDR}: {e.details()}")

def test_grpc_invalid_data():
    with grpc.insecure_channel(GRPC_ADDR) as channel:
        stub = booking_pb2_grpc.BookingServiceStub(channel)
        response = stub.CheckTable(booking_pb2.TableRequest(id=-1))
        assert response is not None