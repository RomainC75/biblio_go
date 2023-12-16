package filter

import "errors"

func FilterISBN( isbn string )(string, error){
	res := ""
	ref :="0123456789"

	for _, c := range isbn{
		for _, r := range ref{
			if c == r{
				res += string(c)
				break
			}
		}
	}
	
	if len(res) != 10 && len(res) != 13{
		return "", errors.New("isbn not valid")
	}
	return res, nil
}