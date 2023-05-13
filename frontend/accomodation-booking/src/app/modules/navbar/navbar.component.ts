import { Component } from '@angular/core';
import { AuthService } from 'src/app/services/auth/auth.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css'],
})
export class NavbarComponent {
  isLoggedIn: boolean = false;
  role: string = '';

  constructor(private authService: AuthService) {}

  ngOnInit(): void {
    this.isLoggedIn = this.authService.isLoggedIn();
    this.role = this.authService.getUserRole();
    this.authService.onLogout.subscribe(() => {
      this.isLoggedIn = false;
      this.role = this.authService.getUserRole();
    });
    this.authService.onLogin.subscribe(() => {
      this.isLoggedIn = true;
      this.role = this.authService.getUserRole();
    });
  }

  logout(): void {
    this.authService.removeToken();
  }
}
