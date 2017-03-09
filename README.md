# pmbot

A) Context aware search

a.1. User profile information, user slack id, name, age, gender other
a.2. Support following queries,

/pmbot what is my fscore?
/pmbot what is my teams fscore?
/pmbot what is my last 2 months fscore?

B) Read & push channel data, audio, images to external api
b.1 Pull data from a slack channel. 
/pmbot read my_msgs , 
/pmbot read all_msgs from this channel
/pmbot read all my audios from this channel
/pmbot read all my images from this channel

C) Start & Record Skype audio call
c.1 Start skype call ex. /pmbot start skype call with conf_id #1234

```
{
    "data": [{
        "type": "my_score",
        "requestor": "slack_bot",
        "requestor_id": "@jeffboss",
        "usr_name": "John",
        "usr_age": 80,
        "usr_gender": "male",
        "requestor_destination": {
            "slack_channel_name": "math101"
        },
        "values": [{
            "score": 5,
            "score_description": "5 means average.",
            "created": "2015-05-22T14:56:29.000Z",
            "updated": "2015-05-22T14:56:28.000Z"
        }],
        "id": "1"
    }]
}
```

```
fscore is by channel ... so i dont want to type /pmbot my fscore for channel dog  or /pmbot my fscore for channel cat ...

if we can pass the channel name to bot or bot knows the command is executed from the context (channel = dog) than it will pull dog fscore and not cat fscore
```
