import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from 'src/app/model/user';
import { AuthService } from '../auth/auth.service';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  apiHost: string = 'http://localhost:8000/auth/';
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient, private authService: AuthService) {}

  createUser(newUser: User): Observable<any> {
    return this.http.post<User>(this.apiHost + 'Register', newUser, {
      headers: this.headers,
    });
  }

  loginUser(user: User): Observable<any> {
    return this.http.post<User>(this.apiHost + 'login', user, {
      headers: this.headers,
    });
  }

  deleteUser(id: string): Observable<any> {
    const header = new HttpHeaders({
      'Content-Type': 'application/json',
      Authorization: this.authService.getToken()
        ? 'Bearer ' + this.authService.getToken()!
        : '',
    });
    return this.http.delete<boolean>(
      `http://localhost:8000/host/deleteUser/${id}`,
      { headers: header }
    );
  }
}
