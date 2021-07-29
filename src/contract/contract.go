package contract

func GetContractAndSetGwidAndDeviceIdInCookie() error {

	contractURL := config.AquareaSmartCloudURL + "/remote/contract"

	_, err := PostREQ(contractURL)
	if err != nil {
		return err
	}
	return nil
}
