import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthService } from '../auth/auth.service';

@Injectable({
  providedIn: 'root',
})
export class RatingService {
  apiHost: string = 'http://localhost:8000/api/rating';
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
    Authorization: this.authService.getToken()
      ? 'Bearer ' + this.authService.getToken()!
      : '',
  });

  constructor(private http: HttpClient, private authService: AuthService) {}

  createRating(newRating: any): Observable<any> {
    const url = this.apiHost;
    return this.http.post(url, JSON.stringify(newRating), {
      headers: this.headers,
    });
  }
}
