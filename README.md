<p align="center">
  <img src="logo.svg" alt="CEFETDB" width="250">
</p>

# 📝 CEFETDB2
CEFETDB 2 is the platform where students of the Bachelor in Information Systems course at CEFET/RJ campus Nova Friburgo share study materials for exams of all disciplines.

The first version of the project was a front-end to the author's own Google Drive running in Apps Script, and was made in two afternoons after a request in the students' group chat. It featured a bunch of security issues, a catastrophically bad mobile version, and a _not so well thought_ design, to say the least. Yet, it was widely adopted by the community and saved many students cramming for their exams. All while displaying the permanent Beta badge, amassing hundreds of files contributed along the year it served.

This 2.0 version consists of a RESTful API back-end made using the [Go Drive package](https://pkg.go.dev/google.golang.org/api/drive/v3) to access the [Google Drive API](https://developers.google.com/drive/api), coupled with a responsive, beautifully made React front-end, thanks to [@NyviDev](https://github.com/NyviDev). The response times are 50% faster on average, students can finally access it over their phones, and hopefully the maintenance requirements will be less frequent.

# 🚀 Running Locally
If you really want to run the project in your own environment, you'll need a Google Cloud Project with the Google Drive API enabled. This way you can get the API credentials needed to run the project, since it uses _Google Drive as a Database_™. You can follow [the official guide](https://developers.google.com/drive/api/quickstart/go) to do so.

As soon as you get your hands on your precious Drive API credentials file, you can toss it in the `backend` folder, and run the `backend/cmd/auth/get_oauth_token.go` script. A URL will be printed in the terminal. Just access it and log in with a Google Account, then click on authorize. You will be redirect to "this site can't be reached" page, but don't worry, just copy the token portion from the URL. It is the beautiful string between `state=state-token&code=` and `&scope`. Now that you have your token, just paste it in the terminal, and the script should do the rest for you.

Now you should able to `go run backend/cmd/main/main.go` and _voilà_, your API is running. To run the front-end, please refer to the readme at [the front-end repository](https://github.com/NyviDev/cefetdb-frontend).

Also, you can use the production Dockerfile in this repository to run it, you'll just have to do the token step manually, as in the current production environment the credentials are being handled by the CI pipeline.

Looking back, I really shoud've made a more streamlined process to run the application. To the ever-growing-never-shrinking To Do list it goes.

# ✨ Contributing
If you'd like to contribute to the project, you can clone this repository and make a new branch. It would be nice if you follow the \<type>/\<description> format, but I'm not really in a position to demand anything here. After committing your changes with nice messages (or no, you do you), just open a pull request. If the improvement you want to do is for the front-end, please apply these steps to the front-end repository.

Oh, by the way, for front-end changes, please use [the front-end repository](https://github.com/NyviDev/cefetdb-frontend). Yeah, I know it's a weird way to structure the project source code, but believe me, there are good reasons for this.

Keep in mind that a contribution doesn't have to be a feature or another code-related change. Documentation, translations, or any other thing you think would be nice for the project to have are welcome.
