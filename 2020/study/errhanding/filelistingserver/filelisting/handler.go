package filelisting
import(
	"net/http"
	"io/ioutil"
	"os"

)

// HandleFileList list页面.
func HandleFileList(w http.ResponseWriter, r *http.Request) error {
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
