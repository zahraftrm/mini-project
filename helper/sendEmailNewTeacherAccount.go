package helper

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

func SendEmailNewTeacherAccount(email, name, password string) error {
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
			<title>Akun Untuk Pengajuan Pelatihan Guru Berhasil Dibuat</title>
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
					<h2>Sekarang kamu memiliki akun!</h2>
				</div>				
				
				<p>Halo, ` + name + `!</p>

				<p>Kamu berhasil terdaftar di sistem pengajuan pelatihan guru <b>EduTrainerHub</b> dengan detail akun sebagai berikut:</p>

				<ul>
					<li>Email: `+ email + ` </li>
					<li>Password: `+ password +`</li>
				</ul>

				<p>Diwajibkan untuk pengguna merubah password yang diberikan untuk menjaga keamanan data.</p>

				<div class="footer">
					<p>EduTrainerHub - Sistem Pengajuan Pelatihan Guru</p>
				</div>
			</div>
		</body>
		</html>
	`
	// data email dikirimkan
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUsername)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Pembuatan Akun Guru untuk Pengajuan Pelatihan")
	m.SetBody("text/html", message)

	d := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

    fmt.Println("email telah berhasil dikirim.")

    return nil
}