<!DOCTYPE html>
<html>
<style>
input[type=text],input[type=date], select {
  width: 100%;
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

input[type=submit] {
  width: 100%;
  background-color: #4CAF50;
  color: white;
  padding: 14px 20px;
  margin: 8px 0;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

input[type=submit]:hover {
  background-color: #45a049;
}

div {
  border-radius: 5px;
  background-color: #f2f2f2;
  padding: 20px;
}
</style>
<body>

<h3>Beego  form Api </h3>

<div>
  <form action="/" method="POST">
    <label for="fname">First Name</label>
    <input type="text"  name="firstname" placeholder="Your first name.."required>

    <label for="lastname">Last Name</label>
    <input type="text"  name="lastname" placeholder="Your last name.."required>

     <label for="phonenumber">Phone Number</label>
    <input type="text"  name="phonenumber" placeholder="Your phone number."required>
    
    <br>
     <label for="email">Email</label>
    <input type="text"  name="email" placeholder="Your email.." required>
     <label for="password">Password</label>
    <input type="text"  name="password" placeholder="Your password."required>
     <label for="birthdate">Birthdate</label>
    <input type="date"  name="birthdate" placeholder="Your birthdate.." required>
    <span style="color:red">{{.flash.error}}</span>
     
    <input type="submit" value="Submit">
  </form>
</div>

</body>
</html>


