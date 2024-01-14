export const convertAnswerDataToNumeric = (value: string): number => {
	let result = 0;

	switch (value) {
		case "Veldig dårlig":
			result = 1;
			break;
		case "Dårlig":
			result = 2;
			break;
		case "Nøytral":
			result = 3;
			break;
		case "Bra":
			result = 4;
			break;
		case "Veldig bra":
			result = 5;
			break;
		case "Vet ikke":
			result = 6;
			break;
		default:
			result = 0;
			break;
	}

	return result;
};

export const convertAnswerDataToString = (value: number): string => {
	let result = "";

	switch (value) {
		case 1:
			result = "Veldig dårlig";
			break;
		case 2:
			result = "Dårlig";
			break;
		case 3:
			result = "Nøytral";
			break;
		case 4:
			result = "Bra";
			break;
		case 5:
			result = "Veldig bra";
			break;
		case 6:
			result = "Vet ikke";
			break;
		default:
			break;
	}

	return result;
};
