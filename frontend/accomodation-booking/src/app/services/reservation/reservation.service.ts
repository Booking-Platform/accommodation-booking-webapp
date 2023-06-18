import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthService } from '../auth/auth.service';

@Injectable({
  providedIn: 'root',
})
export class ReservationService {
  apiHost: string = 'http://localhost:8000/api/reservation';
  apiGateway: string = 'http://localhost:8000/reservation/';

  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
    Authorization: this.authService.getToken()
      ? 'Bearer ' + this.authService.getToken()!
      : '',
  });

  constructor(private http: HttpClient, private authService: AuthService) {}

  createReservation(newReservation: any): Observable<any> {
    const url = this.apiHost;
    return this.http.post(url, newReservation, { headers: this.headers });
  }

  getAllReservationsByUserID(userID: string): Observable<any[]> {
    const url = `${this.apiGateway}getReservationsByUserID/${userID}`;
    return this.http.get<any[]>(url, { headers: this.headers });
  }

  changeReservationStatus(reservationID: any, status: string) {
    const body = { id: reservationID, status: status };

    return this.http.post(this.apiHost + '/changeStatus', body, {
      headers: this.headers,
    });
  }

  getReservationsForConfirmation(): Observable<any[]> {
    const url = `${this.apiGateway}getAllForConfirmation`;
    return this.http.get<any[]>(url, { headers: this.headers });
  }
}
