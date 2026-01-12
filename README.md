# PlantPulse Go 🌿

[![Go Version](https://img.shields.io/github/go-mod/go-version/rahulyadav51/plant-pulse-go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?logo=amazon-aws&logoColor=white)](https://aws.amazon.com/)

**PlantPulse Go** is a serverless application designed for instant plant identification and health diagnosis. Built using **Go (Golang)** and deployed on **AWS**, this project leverages the **Google Gemini API** to analyze plant images and provide actionable care recommendations.

---

## 🚀 The Vision
I am a Full-Stack Developer with 2 years of experience in the **Laravel** ecosystem. As part of my 2026 growth plan, I am building PlantPulse to master **Go** and **AWS Cloud-Native** architectures. This repository serves as a "Learning in Public" project, documenting the transition from traditional PHP MVC frameworks to high-concurrency, serverless Go services.

## 🛠 Tech Stack
- **Language:** Go (Golang)
- **Framework:** Go Fiber (API)
- **Cloud Infrastructure:** - **AWS Lambda:** Serverless compute for Go binaries.
  - **Amazon API Gateway:** RESTful interface for image uploads.
  - **Amazon S3:** Scalable storage for diagnostic history.
- **AI Engine:** Google Gemini API (Multimodal Vision)

## 📐 System Architecture
1. **Request:** A user uploads a plant photo via the API Gateway.
2. **Compute:** An AWS Lambda function written in Go processes the multipart/form-data.
3. **Analysis:** The Go binary sends the image buffer to the Gemini API with a specialized prompt.
4. **Response:** Gemini identifies the species and diagnoses any visible diseases, returning a JSON response to the user.

## 📂 Project Structure
```text
.
├── cmd/
│   └── main.go          # Entry point for the application
├── internal/
│   ├── api/             # API routing and handlers
│   ├── ai/              # Gemini API integration logic
│   └── aws/             # S3 and Lambda specific helpers
├── scripts/             # Deployment and setup scripts
├── go.mod               # Dependencies
└── README.md
