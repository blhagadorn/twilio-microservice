
/* Exchange a hello world with twillio API, get it going, hardcode API keys and stuff
    do springboot version - start with GET, then move to a full POST
    then dockerize it
*/

package com.twilio.microservice;

// Install the Java helper library from twillio.com/docs/libraries/java
import com.twilio.Twilio;
import com.twilio.rest.api.v2010.account.Message;
import com.twilio.type.PhoneNumber;
import static java.lang.System.*;

public class Main
{
    // Find your Account Sid and Token at twilio.com/console
    public static final String ACCOUNT_SID = getenv("ACCOUNT_SID");
    public static final String AUTH_TOKEN = getenv("AUTH_TOKEN");

    public static void main(String[] args) {

        // Initialize the session and authenticate
        Twilio.init(ACCOUNT_SID, AUTH_TOKEN);

        // Create a new message object called "message"
        Message message = Message
                .creator(new PhoneNumber(getenv("TO_NUMBER")), // to
                        new PhoneNumber(getenv("FROM_NUMBER")), // from
                        "What you aim at determines what you see.")
                .create();
    }
}