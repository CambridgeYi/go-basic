/*
@Time : 2019/6/7 13:15 
@Author : dong.liu
@File : gopath
@Software: GoLand
@Description:
*/
package path

import (
	"os"
	"os/exec"
	"path/filepath"
)

// execPath returns the executable path.
func ExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)

}
