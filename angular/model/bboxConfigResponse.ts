/**
 * Queensaver API
 * Queensaver API to send in sensor data and retrieve it. It\'s also used for user management.
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


export interface BboxConfigResponse { 
    /**
     * HTTP response code.
     */
    httpResponseCode?: number;
    /**
     * How often the bbox needs to take and send measurements from the scale. The unit is in seconds.
     */
    scaleMeasureInterval?: number;
    /**
     * How often the bbox needs to send the measurements (in batches)
     */
    batchInterval?: number;
    /**
     * time in ms to wait for GPS, 0 deactivates GPS
     */
    gps?: number;
}

