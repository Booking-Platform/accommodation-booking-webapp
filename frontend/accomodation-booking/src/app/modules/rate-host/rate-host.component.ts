import { Component, OnInit } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { HostServiceService } from 'src/app/services/host/host-service.service';

@Component({
  selector: 'app-rate-host',
  templateUrl: './rate-host.component.html',
  styleUrls: ['./rate-host.component.css']
})
export class RateHostComponent implements OnInit{
  
  public hosts: any[] = [];

  constructor(private hostService: HostServiceService, private jwtHelper: JwtHelperService) {}
  
  ngOnInit(): void {
    var userID = this.jwtHelper.decodeToken().userId;

  
    this.hostService.getHostsForRatingByUserID(userID).subscribe((res: any) => {
      this.hosts = res.hosts;
    });

  
  }

}
