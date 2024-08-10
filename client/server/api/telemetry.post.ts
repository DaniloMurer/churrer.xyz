import {readBody} from "#imports";
import {base64urlEncode} from "iron-webcrypto";

/**
 * Handle call to backend for telemetry from client (BFF principle)
 */
export default defineEventHandler(async (event) : Promise<void> => {
	const apiHost = process.env.API_HOST || 'http://localhost:8080';
	const body = await readBody(event);
	$fetch(`${apiHost}/api/telemetry`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
		},
		body: JSON.stringify(body)
	}).then((telemetry) => {
		console.trace('successfully created telemetry entry', telemetry);
	}).catch((error) => {
		console.error(error, JSON.stringify(body));
	})
})
