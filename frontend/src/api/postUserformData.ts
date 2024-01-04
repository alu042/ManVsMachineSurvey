interface postUserformDataReponse {
	respondentID: number;
}

export const postUserformData = (
	age: string,
	education: string,
	healthcare_personnel: string,
	gender: string
): Promise<postUserformDataReponse> => {
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
	})
		.then((response) => {
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			return response.json();
		})
		.then((data) => {
			console.log(data);
			localStorage.setItem("RespondentId", data);
			return data;
		})
		.catch((error) => {
			console.log(error);
		});
	return response;
};
