package main

func getStatus() string {
	return "OK <script>alert('vuln!');</script>"
}
