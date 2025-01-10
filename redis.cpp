#include <iostream>
#include <hiredis/hiredis.h>

using namespace std;


/**
 * @brief Establishes a connection to a Redis server with a specified timeout.
 *
 * Attempts to connect to a Redis server using the provided hostname and port.
 * Sets a 2-second timeout for the connection attempt. Handles connection errors
 * by logging appropriate error messages and freeing resources if necessary.
 *
 * @param hostname The hostname or IP address of the Redis server
 * @param port The port number of the Redis server
 *
 * @return A pointer to the Redis context if connection is successful, 
 *         nullptr if connection fails or cannot be established
 *
 * @note Logs error messages to cerr if connection initialization fails
 * @note Frees the context if connection cannot be established
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
 * @brief Executes a Redis command on the provided Redis context.
 *
 * This function attempts to execute a given command on a Redis server using the provided Redis context.
 * It performs basic error checking and handles the command execution and response.
 *
 * @param context Pointer to the Redis context for executing the command. Must not be null.
 * @param command The Redis command to be executed as a string.
 *
 * @note If the context is null or command execution fails, appropriate error messages are logged.
 * @note The reply object is automatically freed after processing.
 *
 * @warning Requires a valid Redis connection context to function correctly.
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
 * @brief Main entry point for Redis client demonstration and testing.
 *
 * Establishes a connection to a local Redis server running on the default port.
 * Attempts to initialize a Redis client and execute a PING command.
 * Handles connection failures by using a fallback client mechanism.
 * 
 * @details The function performs the following steps:
 * 1. Sets default Redis server host (localhost) and port
 * 2. Attempts to initialize Redis connection
 * 3. Creates a fallback client if primary connection fails
 * 4. Executes a PING command to verify connectivity
 * 5. Properly frees Redis client resources
 *
 * @return int Exit status of the program (0 indicates successful execution)
 * @note Demonstrates basic Redis client initialization and error handling
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
