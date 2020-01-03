package common /* import "github.com/psych0d0g/anirip/common" */

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

var illegalChars = []string{"\\", "/", ":", "*", "?", "\"", "<", ">", "|"}

const pathSep = string(os.PathSeparator)

// Delete removes a file from the system
func Delete(a ...string) error { return os.Remove(strings.Join(a, pathSep)) }

// Rename renames the source to the desired destination file name and
// recursively retries i times if there are any issues
 func Rename(prevPath, newPath string, mode os.FileMode) error {
     err := os.Rename(prevPath, newPath)
     if err != nil {
         byteArr, err2 := ioutil.ReadFile(prevPath)
              if err2 != nil {
                        return err2
             }

                err2 = ioutil.WriteFile(newPath, byteArr, mode)
         if err2 == nil {
                        // Remove the file iff it was able to be written
                        _ = os.Remove(prevPath)
         } else {
                        return fmt.Errorf("unable to write the file out to the new path. Previous: %s --> New: %s\n", prevPath, newPath)
                       // Remove any partial file data that may have been written in the case of unfulfilled writes.
                   _ = os.Remove(newPath)
          }

                return err2
     }
       return err
}

// GenerateEpisodeFilename constructs an episode filename and returns the
// filename fully sanitized
func GenerateEpisodeFilename(show string, season int, episode float64, desc string) string {
	ep := fmt.Sprintf("%g", episode)
	if episode < 10 {
		ep = "0" + ep // Prefix a zero to episode
	}
	return CleanFilename(fmt.Sprintf("%s - S%sE%s - %s", show,
		fmt.Sprintf("%02d", season), ep, desc))
}

// CleanFilename cleans the filename of any illegal file characters to prevent
// write errors
func CleanFilename(name string) string {
	for _, bad := range illegalChars {
		name = strings.Replace(name, bad, "", -1)
	}
	return strings.Replace(name, "  ", " ", -1)
}
