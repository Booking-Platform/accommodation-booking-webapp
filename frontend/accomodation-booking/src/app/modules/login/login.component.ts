import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { User } from 'src/app/model/user';
import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent {
  user: User = {
    id: null,
    name: '',
    surname: '',
    email: '',
    password: '',
  };

  constructor(
    private userService: UserService,
    private toastr: ToastrService,
    private router: Router,
    private authService: AuthService
  ) {}

  login() {
    this.userService.loginUser(this.user).subscribe((res) => {
      if (res.tokenString != 'no match') {
        this.authService.setToken(res.tokenString);
        this.toastr.success('Logged in!', 'Success');
        this.router.navigate(['/']);
      } else {
        this.toastr.error('Username or password invalid!', 'Error');
      }
    });
  }
}
