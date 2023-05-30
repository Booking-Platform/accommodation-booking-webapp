import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HostServiceService {
  
  apiHost: string = 'http://localhost:8000/host/';
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient) {}
  
  getHostsForRatingByUserID(id: string): Observable<any[]> {
    const url = `${this.apiHost}getHostsForRating/${id}`;
    return this.http.get<any[]>(url, { headers: this.headers });
  }
}
