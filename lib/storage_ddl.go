package lib

import (
	"fmt"
	"log"
	"time"
)

const ddlSleepMilliseconds int = 200

func CreateHealthcareOrganisation(db *Database) (err error) {
	table_name := "hco_healthcare_organisation"
	var rawSQL string = fmt.Sprintf(
		"CREATE TABLE " + table_name +
			" (id SERIAL PRIMARY KEY, " +
			" name VARCHAR(200) NOT NULL, " +
			" code VARCHAR(20) NOT NULL, " +
			" address VARCHAR(120) NULL)")

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' created", table_name)
	} else {
		err = fmt.Errorf("unable to create table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func DropHealthcareOrganisation(db *Database) (err error) {
	table_name := "hco_healthcare_organisation"
	var rawSQL string = fmt.Sprintf("DROP TABLE " + table_name)

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' dropped", table_name)
	} else {
		err = fmt.Errorf("unable to drop table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func CreateHcoEmployee(db *Database) (err error) {
	table_name := "hco_employee"
	var rawSQL string = fmt.Sprintf(
		"CREATE TABLE " + table_name +
			" (id SERIAL PRIMARY KEY, " +
			" hco_id INT NOT NULL, " +
			" firstname VARCHAR(100) NOT NULL, " +
			" surname VARCHAR(100) NOT NULL, " +
			" code VARCHAR(20) NOT NULL) ")

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' created", table_name)
	} else {
		err = fmt.Errorf("unable to create table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func CreateHcoEmployeeProfession(db *Database) (err error) {
	table_name := "hco_employee_profession"
	var rawSQL string = fmt.Sprintf(
		"CREATE TABLE " + table_name +
			" (id SERIAL PRIMARY KEY, " +
			" hco_employee_id INT NOT NULL, " +
			" hco_profession_id INT NOT NULL) ")

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' created", table_name)
	} else {
		err = fmt.Errorf("unable to create table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func DropHcoEmployeeProfession(db *Database) (err error) {
	table_name := "hco_employee_profession"
	var rawSQL string = fmt.Sprintf("DROP TABLE " + table_name)

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' dropped", table_name)
	} else {
		err = fmt.Errorf("unable to drop table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func DropHcoEmployee(db *Database) (err error) {
	table_name := "hco_employee"
	var rawSQL string = fmt.Sprintf("DROP TABLE " + table_name)

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' dropped", table_name)
	} else {
		err = fmt.Errorf("unable to drop table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func CreateHealthcareOrganisationLicense(db *Database) (err error) {
	table_name := "hco_license"
	var rawSQL string = fmt.Sprintf(
		"CREATE TABLE " + table_name +
			" (id SERIAL PRIMARY KEY, " +
			" hco_id INT NOT NULL, " +
			" code VARCHAR(20) NOT NULL, " +
			" date_started TIMESTAMP NULL, " +
			" date_ended TIMESTAMP NULL, " +
			" license_type_name VARCHAR(30) NULL)")

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' created", table_name)
	} else {
		err = fmt.Errorf("unable to create table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func CreateHcoLicenseResidence(db *Database) (err error) {
	table_name := "hco_license_residence"
	var rawSQL string = fmt.Sprintf(
		"CREATE TABLE " + table_name +
			" (id SERIAL PRIMARY KEY, " +
			" hco_license_id INT NOT NULL, " +
			" address VARCHAR(200) NOT NULL)")

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' created", table_name)
	} else {
		err = fmt.Errorf("unable to create table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return
}

func CreateHcoLicenseResidenceService(db *Database) (err error) {
	table_name := "hco_license_residence_service"
	var rawSQL string = fmt.Sprintf(
		"CREATE TABLE " + table_name +
			" (id SERIAL PRIMARY KEY, " +
			" hco_license_residence_id INT NOT NULL, " +
			" hco_service_id INT NOT NULL)")

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' created", table_name)
	} else {
		err = fmt.Errorf("unable to create table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v\n", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return
}

func CreateHcoService(db *Database) (err error) {
	table_name := "hco_service"
	var rawSQL string = fmt.Sprintf(
		"CREATE TABLE " + table_name +
			" (id SERIAL PRIMARY KEY, " +
			"  code VARCHAR(15) NOT NULL, " +
			"  name VARCHAR(200) NOT NULL)")

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' created", table_name)
	} else {
		err = fmt.Errorf("unable to create table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return
}

func CreateHcoProfession(db *Database) (err error) {
	table_name := "hco_profession"
	var rawSQL string = fmt.Sprintf(
		"CREATE TABLE " + table_name +
			" (id SERIAL PRIMARY KEY, " +
			"  code VARCHAR(15) NOT NULL, " +
			"  name VARCHAR(200) NOT NULL)")

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' created", table_name)
	} else {
		err = fmt.Errorf("unable to create table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return
}

func DropHcoProfession(db *Database) (err error) {
	table_name := "hco_profession"
	var rawSQL string = fmt.Sprintf("DROP TABLE " + table_name)

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' dropped", table_name)
	} else {
		err = fmt.Errorf("unable to drop table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func DropHcoService(db *Database) (err error) {
	table_name := "hco_service"
	var rawSQL string = fmt.Sprintf("DROP TABLE " + table_name)

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' dropped", table_name)
	} else {
		err = fmt.Errorf("unable to drop table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func DropHcoLicenseResidenceService(db *Database) (err error) {
	table_name := "hco_license_residence_service"
	var rawSQL string = fmt.Sprintf("DROP TABLE " + table_name)

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' dropped", table_name)
	} else {
		err = fmt.Errorf("unable to drop table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func DropHcoLicenseResidence(db *Database) (err error) {
	table_name := "hco_license_residence"
	var rawSQL string = fmt.Sprintf("DROP TABLE " + table_name)

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' dropped", table_name)
	} else {
		err = fmt.Errorf("unable to drop table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}

func DropHealthcareOrganisationLicense(db *Database) (err error) {
	table_name := "hco_license"
	var rawSQL string = fmt.Sprintf("DROP TABLE " + table_name)

	err = db.ExecuteDDL(rawSQL)

	if err == nil {
		log.Printf("Database table '%s' dropped", table_name)
	} else {
		err = fmt.Errorf("unable to drop table '%s' => %s", table_name, err.Error())
		fmt.Printf("DB ERROR => %v", err)
	}

	time.Sleep(time.Duration(ddlSleepMilliseconds) * time.Millisecond)

	return err
}
