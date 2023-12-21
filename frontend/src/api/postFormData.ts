export const postFormData = (
	age: string,
	education: string,
	healthcare_personnel: string,
	gender: string
) => {
	let url = "http://localhost:8080/submitform";
	let personnel = healthcare_personnel == "Ja" ? true : false;

	const response = fetch(url, {
		method: "POST",
		body: JSON.stringify({
			age: age,
			education: education,
			healthcare_personnel: personnel,
			gender: gender,
		}),
	}).catch((error) => {
		console.log(error);
	});
	return response;
};
