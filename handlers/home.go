package handlers

import (
	"fmt"
	"net/http"
)

// Home handles '/' endpoint
func Home(rw http.ResponseWriter, r *http.Request) {
	html := `<html>
			<header>
			<style>
				.loader {
					border: 16px solid #f3f3f3; /* Light grey */
					border-top: 16px solid #3498db; /* Blue */
					border-radius: 50%;
					width: 120px;
					height: 120px;
					animation: spin 2s linear infinite;
				}
				
				@keyframes spin {
					0% { transform: rotate(0deg); }
					100% { transform: rotate(360deg); }
				}		
			</style>
			</header>
			<body>
				<div style="width: 600px; margin:0 auto; padding-top:20px">
				<h2>Only ".png" files accepted </h2>
					<form enctype="multipart/form-data">
						<input type="file" id="image" style="margin-top:20px;">
						<label style="margin-top:20px; display: block;">How many shapes?</label>
						<input type="text" id="shapes" placeholder="shapes: 1-1000" style="margin-top:20px; display: block;">
						<label style="margin-top:20px; display: block;">Type of transformation</label>
						<select name="mode" id="mode" style="margin-top: 20px;">
							<option value="">--Please choose an option--</option>
							<option value="0">combo</option>
							<option value="1">triangle</option>
							<option value="2">rect</option>
							<option value="3">ellipse</option>
							<option value="4">circle</option>
							<option value="5">rotatedrect</option>
							<option value="6">beziers</option>
							<option value="7">rotatedellipse</option>
							<option value="8">polygon</option>
						</select>
						<button type="submit" style="margin-top:20px; display: block;">Upload Image</button>
					</form>
					<div id="imgContainer" style="padding-top:20px">
						<div class="loader" style="display: none;"></div>
						<img id="outImage" style="width: 500px;" />
					</div>
				</div>
				<script>
					const form = document.querySelector('form')
					form.addEventListener('submit', event => {
						event.preventDefault();

						const formData = new FormData();
						const fileInputElement = document.getElementById("image");
						const shapes = document.getElementById("shapes").value;
						const mode = document.getElementById("mode").value;
						formData.append("image", fileInputElement.files[0]);
						formData.append("shapes", shapes);
						formData.append("mode", mode);
						hideShowLoader()
						fetch("http://localhost:9090/upload",
							{
									body: formData,
									method: "post"
							})
							.then(res => res.blob())
							.then(blob => {
								const objectURL = URL.createObjectURL(blob);
								const outImage = document.getElementById("outImage")
  							outImage.src = objectURL;
								hideShowLoader()
							})
							.catch(err => {
								hideShowLoader()
							});
					})
					function hideShowLoader() {
						const x = document.querySelector(".loader");
						if (x.style.display === "none") {
							x.style.display = "block";
						} else {
							x.style.display = "none";
						}
					}
				</script>
			</body>
		</html>
		`
	fmt.Fprint(rw, html)
}
