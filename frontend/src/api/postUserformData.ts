interface postUserformDataReponse {
	respondentID: number;
}

export const postUserformData = async (
	age: string,
	education: string,
	healthcare_personnel: string,
	gender: string,
	has_answered_before: string,
	county: string,
	submit_date: string
): Promise<postUserformDataReponse> => {
	let url = "https://helseundersokelsen.online/submituserform";
	let personnel = healthcare_personnel == "Ja" ? true : false;
	let answered_before = has_answered_before == "Ja" ? true : false;

	const response = fetch(url, {
		method: "POST",
		body: JSON.stringify({
			age: age,
			education: education,
			healthcare_personnel: personnel,
			gender: gender,
			answered_before: answered_before,
			county: county,
			submit_date: submit_date,
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
			localStorage.setItem("RespondentId", data.respondentID);
			return data;
		})
		.catch((error) => {
			console.log(error);
		});
	return response;
};
