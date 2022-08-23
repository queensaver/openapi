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


export interface VarroaScanMetadata { 
    /**
     * classification of the object - we currently support: bee_leg,bee,mite,bee_wing,ant,wax_moth_droppings,wax_platelets,pollen,bee_droppings,cell_cover_grist
     */
    _class?: string;
    /**
     * how confident the AI is regarding the result
     */
    confidence?: number;
    /**
     * center of the object on the x axis
     */
    xCenter?: number;
    /**
     * center of the object on the y axis
     */
    yCenter?: number;
    /**
     * width of the object
     */
    width?: number;
    /**
     * height of the object
     */
    height?: number;
}

