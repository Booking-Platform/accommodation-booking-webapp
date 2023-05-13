import { Component } from '@angular/core';
import { UserService } from 'src/app/services/user/user.service';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
})
export class RegisterComponent {
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
    private router: Router
  ) {}

  register() {
    this.userService.createUser(this.user).subscribe((res) => {
      if (res.Email === 'error') {
        this.toastr.error('Email address already registered!');
      } else {
        this.toastr.success('Successfully registered! Please login.');
        this.router.navigate(['/login']);
      }
    });
  }
}
