export const postFormData = async (
	respondentID: number,
	allFormAnswers: string
) => {
	let url = "https://helseundersokelsen.no/submitanswers";

	const response = fetch(url, {
		method: "POST",
		body: JSON.stringify({
			respondentID: respondentID,
			allFormAnswers: allFormAnswers,
		}),
	})
		.then((response) => {
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			return response.ok;
		})
		.catch((error) => {
			console.log(error);
		});
	return response;
};
