import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Accommodation } from 'src/app/model/accommodation';

@Injectable({
  providedIn: 'root',
})
export class AccommodationService {  

  apiHost: string = 'http://localhost:8000/';
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient) {}

  getAccommodations(): Observable<Accommodation[]> {
    return this.http.get<Accommodation[]>(this.apiHost + 'accommodations', {
      headers: this.headers,
    });
  }

  search(searchParams: any): Observable<Accommodation[]> {
    let params = new URLSearchParams();
    params.set('from', searchParams.from);
    params.set('to', searchParams.to);
    params.set('numOfGuests', searchParams.numOfGuests);
    params.set('city', searchParams.city);

    return this.http.get<Accommodation[]>(
      `${this.apiHost}reservation/getAllByParams?${params.toString()}`
    );
  }

  createAccommodation(accommodation: any) {
    const { id, ...newAccommodation } = accommodation;
    return this.http.post<any>(
      this.apiHost + 'accommodation/create',
      JSON.stringify(newAccommodation),
      {
        headers: this.headers,
      }
    );
  }

  getAccommodationsByHostID(id: string) {
    window.alert(id)
    const url = `${this.apiHost}accommodation/getAllAccommodationsByHostID/${id}`;
    return this.http.get<any[]>(url, { headers: this.headers });
   }

  getAccommodationByID(id: string): Observable<any[]> {
    const url = `${this.apiHost}accommodations/${id}`;
    return this.http.get<any[]>(url, { headers: this.headers });
  }

  addAppointment(appointment: any) {
    return this.http.post<any>(
      this.apiHost + 'accommodation/appointment',
      JSON.stringify(appointment),
      {
        headers: this.headers,
      }
    );
  }

  changeAutomaticConfirmation(accommodationID: any) {
    return this.http.post<any>(
      this.apiHost + 'accommodation/confirmationStatus',
      JSON.stringify(accommodationID),
      {
        headers: this.headers,
      }
    );
  }
}
