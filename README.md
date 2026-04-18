# traffic-monitor

<img height="500" alt="Screenshot" src="https://github.com/user-attachments/assets/ad75d455-bf0c-45c0-b075-fb520226a497" />

## IMPORTANT
This is a **demo project**.

There are some credentials for test purpose in the codebase,
but there is no security risk here.

Please don't report me about it.

## How to run
Follow these steps:

### Clone the repo
Nothing to say here.

### Install the cert into your browser
I ensured IoT security using mTLS.

Console also requests client certification.

In the `cmd/debug` directory, there are two files:
- `client-test.p12`
- `rootCA.crt`

Install them into your browser.

Never forget to remove them after the test.

### Run the system
You just need to run `docker-compose up --build` in the root directory.

Everything should be fine.

### Access the system
You can access the console at `https://localhost:8443`

At this time there are no data available, but keep the page open.

### Put random data
There is a script `cmd/debug/generate.py` to send randomized data.

Use your new Python environment to run it.

After that, you can see the report on the console.

I designed the console frontend "ping" the backend every second
without heavy queries to the database.

Every second, it acquires the newest ID from the server.
If the ID is newer than the local one,
finally, the frontend requests the information from the server.

## Stop the system
Run `docker-compose down`.

Beware, **never forget to remove the CA certificate** from your browser.

It makes the browser insecure.

## Trade-offs
- There is no test code.
- No ORM used.


