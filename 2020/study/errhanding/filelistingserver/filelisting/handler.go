package filelisting
import(
	"net/http"
	"fmt"
	_"errors"
	"strings"
	"io/ioutil"
	"os"

)
const prefix = "/list/"
type myError string

func (e myError) Error() string {
	return e.Message()
}
func (e myError) Message() string {
	return string(e)
}

// HandleFileList list页面.
func HandleFileList(w http.ResponseWriter, r *http.Request) error {
	if strings.Index(r.URL.Path, prefix) !=0 {
		//return errors.New("path must start with"+prefix)
		return myError(
			fmt.Sprintf("path %s must start with %s", r.URL.Path, prefix),
		)
	}
	path := r.URL.Path[len("/list/"):]
	 file, err := os.Open(path)
	 if err != nil {
		 // panic(err) // 直接panic 网页会down 掉，不优雅
		 // 优雅的报错
		 // http.Error(w, err.Error(), http.StatusInternalServerError)
		 return err
	 }
	 contents, err := ioutil.ReadAll(file)
	 if err != nil {
		 // panic(err) // 直接panic 网页会down 掉，不优雅
		 // 优雅的报错
		 // http.Error(w, err.Error(), http.StatusInternalServerError)
		 return err
	 }
	 w.Write(contents)
	 return nil
}
