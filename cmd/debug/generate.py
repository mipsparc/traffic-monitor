import requests
import random
import time
import uuid
import datetime

# This is a demo script to periodically generate random events and send them to the API

API_URI = 'https://localhost:8443/api/v1/report'

class RandomEvent:
    report_types = {
        'crash': {
            'severity': 80,
            'text': 'a car crashed to the side wall',
        },
        'collision': {
            'severity': 90,
            'text': '3 car collided',
        },
        'stop_in_live_lane': {
            'severity': 50,
            'text': 'stop in live lane',
        },
        'stop_in_shoulder': {
            'severity': 30,
            'text': 'stop in shoulder',
        },
        'wrong_way': {
            'severity': 80,
            'text': '1 car running on wrong way(opposite)',
        },
        'pedestrian': {
            'severity': 25,
            'text': '1 pedestrian on the road',
        },
        'animal': {
            'severity': 20,
            'text': '2 large animals on the road',
        },
        'smoke': {
            'severity': 80,
            'text': 'smoke detected',
        },
        'fire': {
            'severity': 90,
            'text': 'fire detected in the lane 1',
        },
        'too_low_speed': {
            'severity': 10,
            'text': '1 car is running in too low speed',
        },
        'too_fast_speed': {
            'severity': 10,
            'text': '1 car is running in too fast speed',
        },
        'debris': {
            'severity': 25,
            'text': '2 large debris on the road',
        },
        'facility_damage': {
            'severity': 30,
            'text': '1 pole is damaged',
        },
    }

    def generate(self):
        reason = random.choice(list(self.report_types.keys()))
        event = self.report_types[reason]
        # make severity random
        event['severity'] += random.randint(-5, 5)
        return {
            'report_type': reason,
            'severity': event['severity'],
            'text': event['text'],
        }

random_event = RandomEvent()

print("This is a demo script to generate and send random events to the API.")
input("Press Enter to send random events...")

while True:
    time.sleep(random.randint(5,15))

    camera_id = random.randrange(1000,1500)
    data_length = random.randint(1,3)

    send_data = {
        'camera_id': camera_id,
        'report': []
    }
    for i in range(data_length):
        unique_id = str(uuid.uuid4())
        now = datetime.datetime.now(datetime.timezone.utc).isoformat()
        video_id = 'video-' + unique_id + '.mp4'
        lat = 35.6895 + random.uniform(-0.05, 0.05)
        long = 139.7000 + random.uniform(-0.05, 0.05)

        e = random_event.generate()

        report = {
            'uuid': unique_id,
            'time': now,
            'video_id': video_id,
            'lat': lat,
            'long': long,
            'severity': e['severity'],
            'report_type': e['report_type'],
            'text': e['text'],
        }

        send_data['report'].append(report)

    try:
        result = requests.post(
            API_URI,
            json=send_data,
            cert=('client-test.crt', 'client-test.key'),
            verify='rootCA.crt'
        )
    except requests.exceptions.SSLError as e:
        raise Exception(f"TLS error occurred: {e}")
    except requests.exceptions.RequestException as e:
        raise Exception(f"Request failed with error: {e}")

    if result.status_code != 200:
        print(result.text)
        raise Exception(f"Request failed with status code {result.status_code}")
    elif result.status_code == 200:
        print(f"{time.time()} :Request sent.")
