import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ReservationService {
  apiHost: string = 'http://localhost:8000/api/reservation/';
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient) {}

  createReservation(newReservation: any) {
    return this.http.post<any>(
      this.apiHost + 'create',
      JSON.stringify(newReservation),
      {
        headers: this.headers,
      }
    );
  }


  getAllReservationsByUserID(userID: string): Observable<any[]> {
    const url = `${this.apiHost}getReservationsByUserID/${userID}`;
    return this.http.get<any[]>(url, { headers: this.headers });
  }
}
