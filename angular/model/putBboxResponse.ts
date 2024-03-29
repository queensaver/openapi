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
import { Bbox } from './bbox';


export interface PutBboxResponse { 
    /**
     * HTTP response code. Used for internal purposes, will be sent out at the API.
     */
    httpResponseCode?: number;
    bbox?: Bbox;
}

