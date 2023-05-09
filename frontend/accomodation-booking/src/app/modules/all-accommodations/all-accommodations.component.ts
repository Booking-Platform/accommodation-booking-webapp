import { Component, OnInit } from '@angular/core';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';

@Component({
  selector: 'app-all-accommodations',
  templateUrl: './all-accommodations.component.html',
  styleUrls: ['./all-accommodations.component.css']
})
export class AllAccommodationsComponent implements OnInit {
  
  public accommodations: any[] = [];
  
  constructor(private accommodationService: AccommodationService) {}

  ngOnInit(): void {
    this.accommodationService.getAccommodations().subscribe((res: any) => {
      this.accommodations = res.accommodations;
  
    });
  }

  changedAutomaticConfirmation(accommodationID: any) {
    this.accommodationService.changeAutomaticConfirmation(accommodationID).subscribe()
  }
}