import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthService } from '../auth/auth.service';

@Injectable({
  providedIn: 'root',
})
export class HostServiceService {
  apiHost: string = 'http://localhost:8000/host/';
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
    Authorization: this.authService.getToken()
      ? 'Bearer ' + this.authService.getToken()!
      : '',
  });

  constructor(private http: HttpClient, private authService: AuthService) {}

  getHostsForRatingByUserID(id: string): Observable<any[]> {
    const url = `${this.apiHost}getHostsForRating/${id}`;
    return this.http.get<any[]>(url, { headers: this.headers });
  }
}
