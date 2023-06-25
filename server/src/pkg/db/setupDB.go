package andMapMemoriesDB

import (
	"database/sql"
	"os"
	"path"
)

func SetupDB(db *sql.DB) error {
	var err error;
	var initializeSQLDBScript []byte;

	{
		var workingDirectory string;

		{
			workingDirectory, err = os.Getwd()
			if err != nil {
				return err
			}
		}

		{
			initializeSQLDBScript, err = os.ReadFile(path.Join(workingDirectory, "public", "init.sql"))
			if err != nil {
				return err
			}
		}
	}

	_, err = db.Query(string(initializeSQLDBScript))
	if err != nil {
		return err
	}

	return nil
}