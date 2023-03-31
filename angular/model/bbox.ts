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
import { Bhive } from './bhive';


export interface Bbox { 
    /**
     * The ID of the bbox electronix (QBox). Is usually a mac address of a network interface.
     */
    bboxId?: string;
    /**
     * A cron type of description of when the sensore measurements are supposed to be done.
     */
    schedule?: string;
    /**
     * Unique Identifier for this bbox
     */
    uuid?: string;
    bhives?: Array<Bhive>;
    /**
     * If the bbox turns the power off after a successful run and wakes up later according to the schedule.
     */
    powerSave?: boolean;
    /**
     * The registration ID of the bbox. The user needs to put this into the interface so that the bbox can then register via the /configs/bbox/register API call to retrieve the token.
     */
    registrationId?: string;
    /**
     * Hardware type of the bbox - could be a varroa-scanner or a scale, etc.
     */
    hardwareType?: string;
    /**
     * Hardware revision - newer revisions might have different features which are important to know.
     */
    hardwareRevision?: number;
    /**
     * Name of the bbox - optional
     */
    name?: string;
    /**
     * Associated Hive with this bbox
     */
    hiveUuid?: string;
}

