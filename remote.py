# use_remote_class.py
import ray
import time

# Initialize Ray (connect to the Ray cluster)
ray.init(address='auto')

# Get the named actor
remote_instance = ray.get_actor("RemoteSSIServiceActor")

# Call methods on the remote instance
while True:
    value = ray.get(remote_instance.get_current_index_from_stream.remote("VN30"))
    print(f"VN30 value: {value}")
    time.sleep(1)
