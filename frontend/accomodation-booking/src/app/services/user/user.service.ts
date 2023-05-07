import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from 'src/app/model/user';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  apiHost: string = 'http://localhost:8000/api/user/';
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient) {}

  createUser(newUser: User): Observable<any> {
    return this.http.post<User>(this.apiHost + 'create', newUser, {
      headers: this.headers,
    });
  }
}
