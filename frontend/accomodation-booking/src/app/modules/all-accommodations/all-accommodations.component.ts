import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';

@Component({
  selector: 'app-all-accommodations',
  templateUrl: './all-accommodations.component.html',
  styleUrls: ['./all-accommodations.component.css'],
})
export class AllAccommodationsComponent implements OnInit {
  public accommodations: any[] = [];
  public benefits: string[] = [];

  constructor(
    private accommodationService: AccommodationService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.accommodationService.getAccommodations().subscribe((res: any) => {
      this.accommodations = res.accommodations;
    });
  }

  changedAutomaticConfirmation(accommodationID: any) {
    this.accommodationService
      .changeAutomaticConfirmation(accommodationID)
      .subscribe();
  }

  addAppointment(accId: string): void {
    this.router.navigate(['/addAppointment'], {
      queryParams: {
        accommodationID: accId,
      },
    });
  }
}
