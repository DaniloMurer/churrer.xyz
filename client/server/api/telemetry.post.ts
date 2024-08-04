import {readBody} from "#imports";

/**
 * Handle call to backend for telemetry from client (BFF principle)
 */
export default defineEventHandler( async (event) => {
    const body = await readBody(event);
    $fetch('http://localhost:8080/api/telemetry', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        },
        body: JSON.stringify(body)
    }).then((telemetry) => {
        console.trace('successfully created telemetry entry', telemetry);
    }).catch((error) => {
        console.error(error, JSON.stringify(body));
    })
})