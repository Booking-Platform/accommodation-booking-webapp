import { Component, OnInit } from '@angular/core';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';

@Component({
  selector: 'app-view',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.css'],
})
export class ViewComponent implements OnInit {
  public accommodations: Accommodation[] = [];
  public from: string = '';
  public to: string = '';
  public numOfGuests: string = '';
  public city: string = '';

  constructor(
    private accommodationService: AccommodationService,
    public dialog: MatDialog,
    private router: Router
  ) {}

  ngOnInit() {}

  reserveAccommodation(acc: Accommodation) {
    this.router.navigate(['/accommodation-details'], {
      queryParams: {
        accommodationID: acc.id,
        startDate: this.from,
        endDate: this.to,
      },
    });
  }

  search() {
    var searchParams = {
      from: this.from,
      to: this.to,
      numOfGuests: this.numOfGuests,
      city: this.city,
    };
    this.accommodationService.search(searchParams).subscribe((res: any) => {
      this.accommodations = res.accommodations;
    });
  }
}
