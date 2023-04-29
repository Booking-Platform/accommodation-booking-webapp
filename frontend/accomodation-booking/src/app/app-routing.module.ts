import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ViewComponent } from './modules/view/view.component';
import { ReserveAccommodationComponent } from './modules/reserve-accommodation/reserve-accommodation.component';
import { ReservationDetailsComponent } from './modules/reservation-details/reservation-details.component';

const routes: Routes = [
  { path: 'view', component: ViewComponent },
  { path: 'reserve', component: ReserveAccommodationComponent },
  { path: 'accommodation-details', component: ReservationDetailsComponent },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
