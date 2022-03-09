package clienterr

import "fmt"

type Code string

const (
	AppointmentPresentStatusModification = Code("AppointmentPresentStatusModification")
	DuplicateVaccineAppointment          = Code("DuplicateVaccineAppointment")
	InvalidImageExtension                = Code("InvalidImageExtension")
	InvalidImageSize                     = Code("InvalidImageSize")
	DuplicatePatient                     = Code("DuplicatePatient")
	ActiveVaccinationAttemptExists       = Code("ActiveVaccinationAttemptExists")
)

var codeMessageMap = map[Code]string{
	AppointmentPresentStatusModification: "Tidak bisa mengubah status dari hadir ke status lainnya",
	DuplicateVaccineAppointment:          "Tidak bisa membuat jadwal vaksinasi; ada jadwal vaksinasi dengan tipe yang sama yang sudah ditandai hadir atau terjadwal",
	InvalidImageExtension:                "Tipe gambar harus PNG atau JPG",
	InvalidImageSize:                     "Besar gambar harus dibawah 5MB",
	DuplicatePatient:                     "Sudah ada pasien dengan ID yang sama",
	ActiveVaccinationAttemptExists:       "Tidak bisa membuat upaya vaksinasi; hanya satu aktif upaya vaksinasi per pasien",
}

func GenerateMessage(code Code, values ...interface{}) string {
	message, ok := codeMessageMap[code]
	if !ok {
		return ""
	}

	if len(values) == 0 {
		return message
	}
	return fmt.Sprintf(message, values...)
}
