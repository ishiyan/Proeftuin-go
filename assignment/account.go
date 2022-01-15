package assignment

// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`   // for CoP
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"` // for CoP
	AccountNumber           string   `json:"account_number,omitempty"`           // generated if ommited
	AlternativeNames        []string `json:"alternative_names,omitempty"`        // ALWAYS
	BankID                  string   `json:"bank_id,omitempty"`                  // depends on country
	BankIDCode              string   `json:"bank_id_code,omitempty"`             // depends on country
	BaseCurrency            string   `json:"base_currency,omitempty"`            // ALWAYS
	Bic                     string   `json:"bic,omitempty"`                      // ALWAYS
	Country                 *string  `json:"country,omitempty"`                  // ALWAYS
	Iban                    string   `json:"iban,omitempty"`                     // generated if ommited
	JointAccount            *bool    `json:"joint_account,omitempty"`            // for CoP
	Name                    []string `json:"name,omitempty"`                     // ALWAYS
	SecondaryIdentification string   `json:"secondary_identification,omitempty"` // for CoP
	Status                  *string  `json:"status,omitempty"`                   // ALWAYS
	Switched                *bool    `json:"switched,omitempty"`                 // for CoP
}

/*
Country code: GB
Bank ID: required, 6 characters, UK sort code
BIC: required
Bank ID Code: required, has to be GBDSC
Account Number: optional, 8 characters, generated if not provided
IBAN: Generated if not provided

Australia
Country code: AU
Bank ID: optional, 6 characters, Australian Bank State Branch (BSB) code
BIC: required
Bank ID Code: required, has to be AUBSB
Account Number: optional, 6-10 characters, first character cannot be 0, generated if not provided
IBAN: has to be empty

Belgium
Country code: BE
Bank ID: required, 3 characters
BIC: optional
Bank ID Code: required, has to be BE
Account Number: optional, 7 characters, generated if not provided
IBAN: generated if not provided

Canada
Country code: CA
Bank ID: optional, 9 characters starting with zero, Routing Number for Electronic Funds Transfers
BIC: required
Bank ID Code: optional, if provided, has to be CACPA
Account Number: optional, 7-12 characters, generated if not provided
IBAN: not supported, has to be empty

Estonia
Country code: EE
Bank ID: required, 4 characters, bank code + branch code
BIC: required
Bank ID Code: required, has to be EE
Account Number: optional, 12 characters, generated if not provided
IBAN: generated if not provided

France
Country code: FR
Bank ID: required, 10 characters, national bank code + branch code (code guichet)
BIC: optional
Bank ID Code: required, has to be FR
Account Number: optional, 10 characters, generated if not provided
IBAN: generated if not provided

Germany
Country code: DE
Bank ID: required, 8 characters, Bankleitzahl (BLZ)
BIC: optional
Bank ID Code: required, has to be DEBLZ
Account Number: optional, 7 characters, generated if not provided
IBAN: generated if not provided

Greece
Country code: GR
Bank ID: required, 7 characters, HEBIC (Hellenic Bank Identification Code)
BIC: optional
Bank ID Code: required, has to be GRBIC
Account Number: optional, 16 characters, generated if not provided
IBAN: generated if not provided

Hong Kong
Country code: HK
Bank ID: optional, 3 characters, Bank Code or Institution ID
BIC: required
Bank ID Code: optional, if provided, has to be HKNCC
Account Number: optional, 9-12 characters, generated if not provided
IBAN: not supported, has to be empty

Ireland
Country code: IE
Bank ID: required, 6 characters, Irish sort code (NSC)
BIC: required
Bank ID Code: optional, if provided, has to be IENCC
Account Number: optional, 8 characters, generated if not provided
IBAN: generated if not provided

Italy
Country code: IT
Bank ID: required, national bank code (ABI) + branch code (CAB), 10 characters if account number is not present, 11 characters with added check digit as first character if account number is present
BIC: optional
Bank ID Code: required, has to be ITNCC
Account Number: optional, 12 characters, generated if not provided
IBAN: generated if not provided

Luxembourg
Country code: LU
Bank ID: required, 3 characters, IBAN Bank Identifier
BIC: optional
Bank ID Code: required, has to be LULUX
Account Number: optional, 13 characters, generated if not provided
IBAN: generated if not provided

Netherlands
Country code: NL
Bank ID: not supported, has to be empty
BIC: required
Bank ID Code: not supported, has to be empty
Account Number: optional, 10 characters, generated if not provided
IBAN: generated if not provided

Poland
Country code: PL
Bank ID: required, 8 characters, national bank code + branch code + national check digit
BIC: optional
Bank ID Code: required, has to be PLKNR
Account Number: optional, 16 characters, generated if not provided
IBAN: generated if not provided

Portugal
Country code: PT
Bank ID: required, 8 characters, bank identifier + PSP reference number
BIC: optional
Bank ID Code: required, has to be PTNCC
Account Number: optional, 11 characters, generated if not provided
IBAN: generated if not provided

Spain
Country code: ES
Bank ID: required, 8 characters, Código de entidad + Código de oficina
BIC: optional
Bank ID Code: required, has to be ESNCC
Account Number: optional, 10 characters, generated if not provided
IBAN: generated if not provided

Switzerland
Country code: CH
Bank ID: required, 5 characters
BIC: optional
Bank ID Code: required, has to be CHBCC
Account Number: optional, 12 characters, generated if not provided
IBAN: generated if not provided

United States
Country code: US
Bank ID: required, 9 characters, ABA routing number
BIC: required
Bank ID Code: required, has to be USABA
Account Number: optional, 6-17 characters, generated if not provided
IBAN: not supported, has to be empty
*/
