import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ReservationService {
  apiHost: string = 'http://localhost:8000/';
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient) {}

  createReservation(newReservation: any) {
    return this.http.post<any>(
      this.apiHost + 'api/reservation/create',
      JSON.stringify(newReservation),
      {
        headers: this.headers,
      }
    );
  }
}
