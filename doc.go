/*
Please is a utility for making and receiving web requests and parsing and reformatting the common data formats that are sent over them.

Examples:

To make a web request, parse the response as json and reformat it in yaml:

	please get http://my.api.com/things | please parse -i json -o yaml

To listen for a web request on port 8001 and reply "Hello, world" with a 200 status:

	echo "Hello, world" | please respond 200 :8001
*/
package please
