import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RatingService {

  apiHost: string = 'http://localhost:8000/api/rating';
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient) {}

  createRating(newRating: any): Observable<any> {
    const url = this.apiHost;
    return this.http.post(url, JSON.stringify(newRating), { headers: this.headers });
  }

}
