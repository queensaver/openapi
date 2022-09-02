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
import { Hive } from './hive';


export interface Stand { 
    /**
     * A stand can have many hives. However, when sending POST or PUT requests you can only update the stand metadata, any hives attached to this array will be ignored. You will have to do the /hives API call instead to create / update hives with the correct stand UUID. This array will only be populated with GET requests.
     */
    hives?: Array<Hive>;
    /**
     * Name of the stand. Can be chosen by the user. A stand is a collection of bee hives.
     */
    name: string;
    /**
     * latitude of the stand
     */
    latitude?: number;
    /**
     * longitude of the stand
     */
    longitude?: number;
    /**
     * HTTP response code. Used for internal purposes, will be let out at the API level.
     */
    httpReponseCode?: number;
    /**
     * Epoch when the data was last updated. This will be set internally, no need to add this with PUT or POST calls.
     */
    epoch?: number;
    /**
     * Unique Identifier for this stand
     */
    uuid?: string;
    /**
     * if set to true, the hive has been deleted at this epoch.
     */
    deleted?: boolean;
}
