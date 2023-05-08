import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ReservationService {

  
  apiHost: string = 'http://localhost:8000/api/reservation/';
  apiGateway: string = 'http://localhost:8000/reservation/';

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
    const url = `${this.apiGateway}getReservationsByUserID/${userID}`;
    return this.http.get<any[]>(url, { headers: this.headers });
  }

  changeReservationStatus(reservationID: any, status: string) {
    const headers = new HttpHeaders().set('Content-Type', 'application/json');
    const body = { id: reservationID, status: status };

    return this.http.post(this.apiHost + "changeStatus", body, { headers: headers });
  }

  getReservationsForConfirmation(): Observable<any[]> {
    const url = `${this.apiGateway}getAllForConfirmation`;
    return this.http.get<any[]>(url, { headers: this.headers });
  }

}
