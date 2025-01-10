#include <iostream>
#include <hiredis/hiredis.h>

using namespace std;


redisContext* initializeRedis(const string& hostname, int port) {
    struct timeval timeout = {2, 0}; // 2 seconds timeout
    redisContext* context = redisConnectWithTimeout(hostname.c_str(), port, timeout);

    if (context == nullptr || context->err) {
        if (context) {
            cerr << "Redis initialization failed: " << context->errstr << endl;
            redisFree(context);
        } else {
            cerr << "Redis initialization failed: Can't allocate redis context" << endl;
        }
        return nullptr;
    }
    cout << "Redis initialized successfully." << endl;
    return context;
}


redisContext* createFallbackClient() {
  
    cout << "Using fallback Redis client." << endl;
    return nullptr; 
}


void executeCommand(redisContext* context, const string& command) {
    if (!context) {
        cout << "Cannot execute command on null client." << endl;
        return;
    }

    redisReply* reply = (redisReply*)redisCommand(context, command.c_str());
    if (!reply) {
        cerr << "Failed to execute command: " << context->errstr << endl;
        return;
    }

    cout << "Command executed successfully: " << reply->str << endl;
    freeReplyObject(reply);
}

int main() {
    const string redisHost = "127.0.0.1";
    const int redisPort = 6379;

  
    redisContext* redisClient = initializeRedis(redisHost, redisPort);

   
    if (!redisClient) {
        redisClient = createFallbackClient();
    }

    
    executeCommand(redisClient, "PING");


    if (redisClient) {
        redisFree(redisClient);
    }

    return 0;
}
