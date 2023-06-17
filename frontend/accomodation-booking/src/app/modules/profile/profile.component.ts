import { Component } from '@angular/core';
import { Router } from '@angular/router';
import jwtDecode from 'jwt-decode';
import { ToastrService } from 'ngx-toastr';
import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css'],
})
export class ProfileComponent {
  name: string = '';
  surname: string = '';
  email: string = '';
  address: string = '';

  ngOnInit(): void {
    localStorage.setItem('name', 'Petar');
    localStorage.setItem('surname', 'Popovic');
    localStorage.setItem('email', 'perap@gmail.com');
    localStorage.setItem('address', 'Skojevska 23, Novi Sad, Serbia');

    this.name = localStorage.getItem('name')!;
    this.surname = localStorage.getItem('surname')!;
    this.email = localStorage.getItem('email')!;
    this.address = localStorage.getItem('address')!;
  }

  constructor(
    private authService: AuthService,
    private userService: UserService,
    private toastr: ToastrService,
    private router: Router
  ) {}

  enableInputs() {
    var editButton = document.getElementById('edit-button');
    editButton!.style.display = 'none';

    var confirmButton = document.getElementById('confirm-button');
    confirmButton!.style.display = 'inline-block';

    var cancelButton = document.getElementById('cancel-button');
    cancelButton!.style.display = 'inline-block';

    var inputs = document.querySelectorAll('.user-profile input');
    for (var i = 0; i < inputs.length; i++) {
      inputs[i].removeAttribute('disabled');
    }
  }

  disableInputs() {
    var inputs = document.querySelectorAll('.user-profile input');
    for (var i = 0; i < inputs.length; i++) {
      inputs[i].setAttribute('disabled', 'disabled');
    }

    var editButton = document.getElementById('edit-button');
    editButton!.style.display = 'inline-block';

    var confirmButton = document.getElementById('confirm-button');
    confirmButton!.style.display = 'none';

    var cancelButton = document.getElementById('cancel-button');
    cancelButton!.style.display = 'none';
  }

  editProfile() {
    localStorage.setItem('name', this.name);
    localStorage.setItem('surname', this.surname);
    localStorage.setItem('email', this.email);
    localStorage.setItem('address', this.address);

    this.name = localStorage.getItem('name')!;
    this.surname = localStorage.getItem('surname')!;
    this.email = localStorage.getItem('email')!;
    this.address = localStorage.getItem('address')!;

    this.disableInputs();
  }

  deleteUser() {
    const tokenString = this.authService.getToken();
    if (tokenString) var decoded = jwtDecode(tokenString) as any;
    this.userService.deleteUser(decoded['userId']).subscribe((res) => {
      if (res) {
        this.toastr.success('Deleted!');
        this.authService.removeToken();
        this.router.navigate(['/']);
      } else {
        this.toastr.error("Can't delete!");
      }
    });
  }
}
