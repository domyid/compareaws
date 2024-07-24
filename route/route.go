package route

import (
	"fmt"
	"net/http"

	config "github.com/domyid/domyapi/config"
	controller "github.com/domyid/domyapi/controller"
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // If it's a preflight request, return early.
	}

	var method, path string = r.Method, r.URL.Path

	// Logging request details
	fmt.Printf("Received request - Method: %s, Path: %s\n", method, path)

	switch {
	case method == "POST" && path == "/login":
		fmt.Println("Routing to LoginSiakad")
		controller.LoginSiakad(w, r)
	case method == "POST" && path == "/refresh-token":
		fmt.Println("Routing to RefreshTokens")
		controller.RefreshTokens(w, r)
	case method == "GET" && path == "/data/mahasiswa":
		fmt.Println("Routing to GetMahasiswa")
		controller.GetMahasiswa(w, r)
	case method == "GET" && path == "/data/bimbingan/mahasiswa":
		fmt.Println("Routing to GetListBimbinganMahasiswa")
		controller.GetListBimbinganMahasiswa(w, r)
	case method == "POST" && path == "/data/bimbingan/mahasiswa":
		fmt.Println("Routing to PostBimbinganMahasiswa")
		controller.PostBimbinganMahasiswa(w, r)
	case method == "GET" && path == "/data/dosen":
		fmt.Println("Routing to GetDosen")
		controller.GetDosen(w, r)
	case method == "POST" && path == "/jadwalmengajar":
		fmt.Println("Routing to GetJadwalMengajar")
		controller.GetJadwalMengajar(w, r)
	case method == "POST" && path == "/riwayatmengajar":
		fmt.Println("Routing to GetRiwayatPerkuliahan")
		controller.GetRiwayatPerkuliahan(w, r)
	case method == "POST" && path == "/absensi":
		fmt.Println("Routing to GetAbsensiKelas")
		controller.GetAbsensiKelas(w, r)
	case method == "POST" && path == "/nilai":
		fmt.Println("Routing to GetNilaiMahasiswa")
		controller.GetNilaiMahasiswa(w, r)
	case method == "POST" && path == "/BAP":
		fmt.Println("Routing to GetBAP")
		controller.GetBAP(w, r)
	case method == "GET" && path == "/data/list/ta":
		fmt.Println("Routing to GetListTugasAkhirMahasiswa")
		controller.GetListTugasAkhirMahasiswa(w, r)
	case method == "POST" && path == "/data/list/bimbingan":
		fmt.Println("Routing to GetListBimbinganMahasiswa")
		controller.GetListBimbinganMahasiswa(w, r)
	case method == "POST" && path == "/approve/bimbingan":
		fmt.Println("Routing to ApproveBimbingan")
		controller.ApproveBimbingan(w, r)
	default:
		fmt.Println("Routing to NotFound")
		controller.NotFound(w, r)
	}
}
