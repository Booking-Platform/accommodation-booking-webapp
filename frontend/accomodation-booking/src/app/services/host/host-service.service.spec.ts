import { TestBed } from '@angular/core/testing';

import { HostServiceService } from './host-service.service';

describe('HostServiceService', () => {
  let service: HostServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HostServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
