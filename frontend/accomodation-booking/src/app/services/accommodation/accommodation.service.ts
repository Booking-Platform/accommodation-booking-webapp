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
    return this.http.post<any>(
      this.apiHost + 'accommodations/search',
      JSON.stringify(searchParams),
      {
        headers: this.headers,
      }
    );
  }

  createAccommodation(accommodation: any) {
    const { id, ...newAccommodation } = accommodation;
    console.log(newAccommodation);

    return this.http.post<any>(
      this.apiHost + 'accommodation/create',
      JSON.stringify(newAccommodation),
      {
        headers: this.headers,
      }
    );
  }
}
