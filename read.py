import pickle
import zlib

import cv2
import numpy as np
from tqdm import tqdm
import redis
import datetime

# Connect to the Redis server
redis_host = '192.168.0.119'
redis_port = 6379  # Default Redis port
redis_db = 0       # Default Redis database

client = redis.StrictRedis(host=redis_host, port=redis_port, db=redis_db)

# Define the time range
check_ts = "07:23:03"
time_format = "%H:%M:%S"
current_date = datetime.date.today()
check_time = datetime.datetime.combine(current_date, datetime.datetime.strptime(check_ts, time_format).time())
time_from = (check_time - datetime.timedelta(seconds=4)).timestamp() * 1000
time_to = (check_time + datetime.timedelta(seconds=4)).timestamp()  * 1000


def read_keys_in_range(start, end):
    start = int(start)
    end = int(end)
    for key in tqdm(range(start, end + 1)):
        key = str(int(key))
        value = client.get(key) # or client.get("2" + key)
        if value:
            frame = pickle.loads(zlib.decompress(value))  # opencv numpy (60*60*3) array
            yield key, frame


target_size = (300, 300)
# sr = cv2.dnn_superres.DnnSuperResImpl_create()
# sr.readModel("ESPCN_x4.pb")
# sr.setModel("espcn", 4)

for key, frame in read_keys_in_range(time_from, time_to):
    # Assuming the frame is an image
    frame_image = np.array(frame)
    # Convert the key back to a datetime object
    key_datetime = datetime.datetime.fromtimestamp(int(key) / 1000)
    key_datetime = str(key_datetime).split(" ")[-1]
    # Upscale the image to the target size
    upscaled_frame = cv2.resize(frame_image, target_size, interpolation=cv2.INTER_LINEAR)
    # upscaled_frame = sr.upsample(upscaled_frame)

    # Display the frame with annotation
    annotated_frame = cv2.putText(upscaled_frame.copy(), key_datetime, (10, 30), cv2.FONT_HERSHEY_SIMPLEX, 1, (255, 0, 0), 2, cv2.LINE_AA)
    cv2.imshow('Annotated Frame', annotated_frame)
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

cv2.destroyAllWindows()
