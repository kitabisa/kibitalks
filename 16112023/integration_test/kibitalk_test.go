package integration_test

import (
	"context"
	"fmt"
	"github.com/kitabisa/perkakas/database/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
	"gopkg.in/resty.v1"
	"os"
	"testing"
	"time"
)

// Your test suite struct
type KibiTestSuite struct {
	suite.Suite
	// Add any fields or dependencies needed for your tests
}

// Setup the suite (runs once before any tests)
func (suite *KibiTestSuite) SetupSuite() {
	identifier := tc.StackIdentifier("kibitalk-api")
	compose, err := tc.NewDockerComposeWith(tc.WithStackFiles("docker-compose.yml"), identifier)
	assert.NoError(suite.T(), err, "NewDockerComposeAPIWith()")

	suite.T().Cleanup(func() {
		assert.NoError(suite.T(), compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	suite.T().Cleanup(cancel)

	err = compose.
		WaitForService("kibitalk-api", wait.ForHTTP("/health_check/db").WithPort("45001")).
		WaitForService("mockoon", wait.ForHTTP("/").WithPort("8083")).
		Up(ctx, tc.Wait(true))

	assert.NoError(suite.T(), err, "compose.Up()")

	time.Sleep(10 * time.Second)
}

// Tear down the suite (runs once after all tests)
func (suite *KibiTestSuite) TearDownSuite() {
	// Add tear down logic here
}

// Your unit tests go here
func (suite *KibiTestSuite) TestMigrationSuccess() {
	mysqlBuilder, err := mysql.NewMySQLConfigBuilder().WithHost("localhost").
		WithPort(3306).
		WithPassword("pass").
		WithUsername("root").
		WithDBName("kibitalk").Build()

	if err != nil {
		fmt.Println(err)
		assert.Fail(suite.T(), "failed to build mysql config")
	}

	db, err := mysqlBuilder.InitMysqlDB()
	if err != nil {
		fmt.Println(err)
		assert.Fail(suite.T(), "failed to init mysql db")
		return
	}

	d, e := os.ReadDir("../migrations/sql")
	if e != nil {
		fmt.Println(err)
		assert.Fail(suite.T(), "failed to count migrations file")
		return
	}

	numberOfFiles := len(d)

	fmt.Printf("Number of files are %d\n", numberOfFiles)
	var numberOfMigrations int

	err = db.QueryRow("select count(*) from gorp_migrations").Scan(&numberOfMigrations)
	switch {
	case err != nil:
		fmt.Println(err)
		assert.Fail(suite.T(), "failed to query mysql db")
		return
	default:
		fmt.Printf("Number of rows are %d\n", numberOfMigrations)
	}

	assert.Equal(suite.T(), numberOfFiles, numberOfMigrations)
}

func (suite *KibiTestSuite) TestCreateDonation() {
	// Create a Resty Client
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"amount": 50000, "payment_method_id": 1, "campaign_id": 1}`).
		Post("http://localhost:45001/v1/donation")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.StatusCode())
}

func (suite *KibiTestSuite) TestCreateDonationPayloadAmountError() {
	// Create a Resty Client
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"amount": "50000", "payment_method_id": 1, "campaign_id": 1}`).
		Post("http://localhost:45001/v1/donation")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 400, resp.StatusCode())
}

func (suite *KibiTestSuite) TestGetDonationSuccess() {
	// Create a Resty Client
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get("http://localhost:45001/v1/donation/1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.StatusCode())
	assert.Equal(suite.T(), `{"id":1,"amount":10000,"payment_method":"Kantong Donasi","campaign":"Bantuan Kemanusiaan untuk Gaza"}`, string(resp.Body()))
}

func TestSuite(t *testing.T) {
	// Create a new TestSuite
	suite.Run(t, new(KibiTestSuite))
}
