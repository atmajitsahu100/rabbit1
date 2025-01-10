#include <iostream>
#include <hiredis/hiredis.h>

using namespace std;


/**
 * @brief Establishes a connection to a Redis server with a specified timeout.
 *
 * Attempts to connect to a Redis server using the provided hostname and port.
 * Sets a 2-second timeout for the connection attempt to prevent indefinite blocking.
 *
 * @param hostname The hostname or IP address of the Redis server
 * @param port The port number of the Redis server
 *
 * @return A pointer to the initialized redisContext if connection is successful,
 *         nullptr if connection fails or context cannot be allocated
 *
 * @note If connection fails, the function logs an error message to cerr
 * @note Any allocated context is properly freed in case of connection failure
 *
 * @exception None Throws no exceptions, returns nullptr on connection failure
 */
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


/**
 * @brief Creates a fallback Redis client when the primary connection fails.
 *
 * This function is called when the primary Redis connection cannot be established.
 * It logs a message indicating that a fallback client is being used and returns a null pointer.
 *
 * @return nullptr to indicate no valid Redis connection is available
 */
redisContext* createFallbackClient() {
  
    cout << "Using fallback Redis client." << endl;
    return nullptr; 
}


/**
 * @brief Executes a Redis command on the provided context.
 *
 * @param context Pointer to the Redis connection context
 * @param command Redis command to be executed as a string
 *
 * @details Attempts to execute a Redis command and handle the response.
 * Checks for a valid Redis context before command execution.
 * Logs an error if the context is null or command execution fails.
 * Prints the command result and frees the reply object after successful execution.
 *
 * @note Requires a valid Redis connection context
 * @note Outputs command result or error to console
 * @warning Does not throw exceptions, uses error logging instead
 */
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

/**
 * @brief Main entry point for Redis client demonstration.
 *
 * Establishes a connection to a local Redis server and performs a basic connectivity test.
 * 
 * @details The function attempts to:
 * 1. Connect to a Redis server at localhost (127.0.0.1) on default port 6379
 * 2. If connection fails, create a fallback client
 * 3. Execute a "PING" command to test server connectivity
 * 4. Properly free the Redis client resources
 *
 * @note Uses predefined local Redis server configuration
 * @note Handles potential connection failures gracefully
 * @return int Exit status of the program (0 indicates successful execution)
 */
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
