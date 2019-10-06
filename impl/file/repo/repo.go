package repo

import (
	"os"
	"path"

	dgrzerr "github.com/datacequia/go-dogg3rz/errors"
	"github.com/datacequia/go-dogg3rz/impl/file"
)

type FileRepositoryResource struct {
}

func (repo *FileRepositoryResource) InitRepo(name string) error {

	repoDir := path.Join(file.RepositoriesDirPath(), name)
	// CREATE 'refs/heads' SUBDIR
	refsDir := path.Join(repoDir, "refs")
	headsDir := path.Join(refsDir, "heads")

	dirsList := []string{repoDir, refsDir, headsDir}

	for _, d := range dirsList {

		err := os.Mkdir(d, os.FileMode(0700))

		if err != nil {
			if os.IsNotExist(err) {
				// BASE REPO DIR DOES NOT EXIST
				return dgrzerr.NotFound.Wrapf(err, file.RepositoriesDirPath())
			}
			if os.IsExist(err) {

				return dgrzerr.AlreadyExists.Wrapf(err, name)
			}

			return err

		}
		// WRITE THE HEAD FILE WITH A POINTER TO DEFAULT MASTER BRANCH
		err = file.WriteHeadFile(name, file.MasterBranchName)
		if err != nil {
			return err
		}

	}

	return nil

}
