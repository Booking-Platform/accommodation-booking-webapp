import { EventEmitter, Injectable } from '@angular/core';
import jwtDecode from 'jwt-decode';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor() {}

  private readonly tokenKey = 'jwt-token';
  onLogout: EventEmitter<void> = new EventEmitter<void>();
  onLogin: EventEmitter<void> = new EventEmitter<void>();

  setToken(token: string): void {
    localStorage.setItem(this.tokenKey, token);
    this.onLogin.emit();
  }

  getUserRole() {
    var jwt;
    if (this.getToken()) {
      jwt = jwtDecode(this.getToken()!) as any;
    } else {
      return '';
    }
    return jwt.role;
  }

  getToken(): string | null {
    return localStorage.getItem(this.tokenKey);
  }

  removeToken(): void {
    localStorage.removeItem(this.tokenKey);
    this.onLogout.emit();
  }

  isLoggedIn(): boolean {
    const token = this.getToken();
    return !!token;
  }
}
