package helper

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmailTrainingStatus(email, name, status, title string) error {
	fmt.Println(email, name, status, title)

	smtpServer := os.Getenv("SMTPSERVER")
	smtpPortStr := os.Getenv("SMTPPORT")
	smtpUsername := os.Getenv("SMTPUSERNAME")
	smtpPassword := os.Getenv("SMTPPASSWORD")

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return err
	}

	// pesan email
	message := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Perubahan Status Pengajuan Pelatihan</title>
			<style>
				body {
					 font-family: Tahoma, Verdana, sans-serif;
					margin: 0;
					padding: 0;
				}
				.container {
					max-width: 600px;
					margin: 0 auto;
					padding: 20px;
					background-color: #ffffff;
					box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
				}
				h2 {
					color: #000;
				}
				p {
					color: #000;
					font-size: 16px;
				}
				ul {
					list-style-type: disc;
				}
				li {
					font-size: 16px;
					color: #000;
				}
				.footer {
					background-color: #D3D3D3;
					color: #000;
					padding: 5px 0;
					text-align: center;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div "header">
					<h2>Status Pengajuanmu Telah Diperbaharui!</h2>
				</div>
				
				<p>Halo, ` + name + `!</p>

				<p>Status pengajuan pelatihan milikmu dengan judul <b> `+ title +` </b> telah mengalami perubahan status sebagai berikut:</p>

				<ul>
					<li>Email: `+ email + ` </li>
					<li>Status: `+ status +`</li>
				</ul>

				<p>Jika status berubah menjadi <b>'Revisi'</b> kamu wajib melakukan perubahan sedangkan jika status berubah menjadi <b> 'Disetujui' </b> kamu dapat memulai kegiatan pelatihanmu.</p>

				<div class="footer">
					<p>EduTrainerHub - Sistem Pengajuan Pelatihan Guru</p>
				</div>
			</div>
		</body>
		</html>
	`
	d := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)

	// alamat email
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUsername)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Perubahan Status Pengajuan Pelatihan")
	m.SetBody("text/html", message)

	
	fmt.Println(smtpServer, smtpPort, smtpUsername, smtpPassword)
	
	// kirim email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("gagal mengirim email:", err)
	}

	fmt.Println("Email telah berhasil dikirim.")

	return nil
}