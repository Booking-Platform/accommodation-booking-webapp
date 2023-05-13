import { EventEmitter, Injectable } from '@angular/core';

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
