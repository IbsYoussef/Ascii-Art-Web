# Ascii-Art-Web

A web-based version of the original CLI tool that renders styled ASCII art from user input using banner templates such as `standard`, `shadow`, and `thinkertoy`. Built with Go for the backend and a minimal web frontend, this project brings your ASCII art generation directly to the browser.


## 📑 Table of Contents

1. [📝 About](#-about)  
2. [📁 File Structure](#-file-structure)  
3. [✨ Features](#-features)  
4. [🚀 Usage Instructions](#-usage-instructions)  
   - [📦 Clone the Repository](#-clone-the-repository)  
   - [▶️ Run the Web App](#️-run-the-web-app)  
5. [🌐 Live Demo](#-live-demo)  
6. [🔭 Future Plans](#-future-plans)  
7. [🤝 Contributions](#-contributions)  
8. [🙏 Acknowledgements](#-acknowledgements)  
9. [📄 License](#-license)  


## 📝 About

This project is the web adaptation of my CLI-based ASCII Art tool, designed to be simple and accessible. It taught me how to implement a basic Go backend and set up a web server from scratch. The project features a web UI that allows users to enter text, select a banner style, and view the generated ASCII art directly in their browser. The backend processes the input and renders it using font templates, while the frontend provides a minimal and clean user experience using HTML, CSS, and JavaScript.

## 📁 File structure
```
.
├── README.md
├── assets # # Static frontend assets 
│   ├── css # Styling for the web page
│   │   └── styles.css
│   ├── img # Image assets for frontend
│   │   ├── Go-Computing.gif
│   │   ├── Image.png
│   │   ├── gopher_lifting_weights.gif
│   │   └── gophercises_jumping.gif
│   └── js # Client-side scripting
│       └── script.js
├── banners # Banners for font styles
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
├── cmd # Web server entry point
│   └── main.go
├── go.mod
├── internal # Core ascii rendering logic
│   ├── ascii
│   │   ├── input.go
│   │   ├── input_test.go
│   │   ├── loadBanner.go
│   │   ├── loadBanner_test.go
│   │   ├── renderAscii.go
│   │   └── renderAscii_test.go
│   ├── files # File reading and helpers
│   │   ├── readFile.go
│   │   └── readFile_test.go
│   └── web # Web logic (HTTP handlers & routes)
│       ├── handler.go
│       └── router.go
└── templates # HTML template served to users
    └── index.html

12 directories, 24 files
```

## ✨ Features
- 🌐 Web interface for generating ASCII art from text
- 🎨 Banner style selection: `standard`, `shadow`, `thinkertoy`
- 💬 Multi-line input with support for `\n` as line breaks
- 🔀 Modular backend architecture using Go’s `internal` packages

## 🚀 Usage Instructions

- ### 📦 Clone the Repository
Clone the repository to your local machine:
```
git clone https:#github.com/IbsYoussef/ascii-art-web.git
cd ascii-art-web
```

- ### ▶️ Run the web app 
```
go run ./cmd
```
Visit http:#localhost:9800 to access the web interface.

## 🌐 Live Demo

Coming soon

## 🔭 Future Plans
Here are a few enhancements I plan to add in future updates:

- **🖼️ Responsive Design**: Make the UI mobile/tablet friendly

- **🎨 UI Redesign**: Introduce themes, animations, and cleaner layout

- **📀 Download Option**: Let users download ASCII output as .txt

##  🤝 Contributions
Contributions are welcome! If you'd like to help improve **ascii-art-web**, please follow these steps:

1. **Fork the Repository:**  
   Click the "Fork" button at the top-right of the repository page to create your own copy of the project.

2. **Create a New Branch:**  
   Create a new branch for your feature or bug fix:
   ```bash
    git checkout -b feature-or-bugfix-description
   ```
3. **Make your Changes:**
Implement your changes and ensure that your code adheres to the project's style guidelines.
Tip: Write or update tests as needed.

4. **Commit and Push your Changes**:
Commit your changes with a clear, descriptive message and push your branch to your forked repository:
    ```bash
    git commit -m "Add: description of your changes"
    git push origin feature-or-bugfix-description
    ```
5. **Open a Pull Request**:
Open a pull request (PR) from your branch to the main repository. Please include a clear description of your changes and the motivation behind them.
If you're not sure about a major change, open an issue first to discuss your ideas.

Thank you for helping make ascii-art-web even better!


## 🙏 Acknowledgements
- Created as part of my Go learning journey at 01 Founders


## 📄 License
This project is licensed under the [MIT License](LICENSE).

Acknowledgements
Special Thanks:
Thanks to all contributors, mentors, and peers who provided feedback and support during the development of go-reloaded.

Inspiration:
This project was inspired by best practices in Go development and the need for automated text formatting solutions.

Resources:

The MIT License
Various open-source projects and communities that encourage collaboration and learning.
Thank you for checking out go-reloaded! We hope this tool helps streamline your text processing tasks and that you find it both useful and easy to contribute to.