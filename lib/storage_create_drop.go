package lib

import (
	"log"
	"os"
	"time"
)

const DEFAULT_CONNECTION_STRING = "host=localhost user=hco password=hco dbname=mydb"

// GetDBURL returns database connection string
func GetDBURL() string {
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		log.Printf("env variable 'DATABASE_URL' found => %s", dbURL)
		return dbURL
	}

	log.Printf(
		"env variable 'DATABASE_URL' not found, using default connection string (%s)",
		DEFAULT_CONNECTION_STRING)

	return DEFAULT_CONNECTION_STRING
}

func CreateDatabaseTables(db *Database) (err error) {
	log.Printf("* * * * * * * * * * * * * * * * * * * * * * * *")
	log.Printf("Creating HCO database tables...")
	log.Printf("")

	time.Sleep(1 * time.Second)

	err = CreateHealthcareOrganisation(db)
	if err != nil {
		return err
	}

	err = CreateHcoEmployee(db)
	if err != nil {
		return err
	}

	err = CreateHcoEmployeeProfession(db)
	if err != nil {
		return err
	}

	err = CreateHealthcareOrganisationLicense(db)
	if err != nil {
		return err
	}

	err = CreateHcoLicenseResidence(db)
	if err != nil {
		return err
	}

	err = CreateHcoLicenseResidenceService(db)
	if err != nil {
		return err
	}

	err = CreateHcoService(db)
	if err != nil {
		return err
	}

	err = CreateHcoProfession(db)

	log.Printf("")
	log.Printf("HCO database tables created!")
	time.Sleep(1 * time.Second)
	log.Printf("* * * * * * * * * * * * * * * * * * * * * * * * ")

	return err
}

func DropDatabaseTables(db *Database) (err error) {
	log.Printf("* * * * * * * * * * * * * * * * * * * * * * * * \n")
	log.Printf("Dropping HCO database tables...")
	log.Printf("")

	time.Sleep(1 * time.Second)

	err = DropHealthcareOrganisation(db)
	if err != nil {
		return err
	}

	err = DropHcoEmployee(db)
	if err != nil {
		return err
	}

	err = DropHcoEmployeeProfession(db)
	if err != nil {
		return err
	}

	err = DropHealthcareOrganisationLicense(db)
	if err != nil {
		return err
	}

	err = DropHcoLicenseResidence(db)
	if err != nil {
		return err
	}

	err = DropHcoLicenseResidenceService(db)
	if err != nil {
		return err
	}

	err = DropHcoService(db)
	if err != nil {
		return err
	}

	err = DropHcoProfession(db)

	log.Printf("")
	log.Printf("HCO database tables dropped!\n")
	time.Sleep(1 * time.Second)
	log.Printf("* * * * * * * * * * * * * * * * * * * * * * * * ")

	return err
}
