package cmd

const CHAINDATA_ROOT_DIR = 1

func GetValue(command uint8) (*string, error) {
	var value *string
	var err error

	switch command {
	case CHAINDATA_ROOT_DIR:
		{
			*value = "db"
		}
	}

	return value, err
}
