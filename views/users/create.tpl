<div align = "middle">
<div class="row">
    <div class="col-md-12" role="main">
      <div class="post-container">     
         <!-- content --> <!-- query_form -->
         <div class="post-content">		
           <form id="query_form" class="form-horizontal form-well" role="form" action="/users/create" method="post">			
              <!--<div>Id: <input type="text" class="form-control" id="Id" name="Id"></div>-->
              <div>Name: <input type="text" class="form-control" id="Name" name="Name"></div>
              <div class="mb-3">
                <label for="Email" class="form-label">Email</label>
                <input type="text" class="form-control {{if .errors.Email}}is-invalid{{end}}" id="Email" name="Email" value="{{.model.Email}}" required>
                {{if .errors.Email}}<div class="invalid-feedback">{{.errors.Email}}</div>{{end}}
              </div>
              <div class="mb-3">
                <label for="Phone" class="form-label">Phone</label>
                <input type="text" class="form-control {{if .errors.Phone}}is-invalid{{end}}" id="Phone" name="Phone" value="{{.model.Phone}}" required>
                {{if .errors.Phone}}<div class="invalid-feedback">{{.errors.Phone}}</div>{{end}}
              </div>
              <!-- <div>Email: <input type="text" class="form-control" id="Email" name="Email" required></div> -->
              <!-- <div>Phone: <input type="text" class="form-control" id="Phone" name="Phone" required></div> -->
              <div>Password: <input type="text" class="form-control" id="Password" name="Password"></div>
              <div> <button type="submit" class="btn btn-primary">Submit</button></div>
           </form>
         </div>
      </div>
    </div>   
  </div>
</div>
