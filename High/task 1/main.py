import grpc
from proto import booking_pb2, booking_pb2_grpc

def run_grpc_client():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = booking_pb2_grpc.BookingServiceStub(channel)
        response = stub.GetStatus(booking_pb2.StatusRequest(table_id=5))
        print("Go gRPC Server says:", response.status)