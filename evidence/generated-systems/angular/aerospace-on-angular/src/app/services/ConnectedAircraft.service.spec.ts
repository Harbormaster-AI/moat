import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ConnectedAircraftService } from './ConnectedAircraft.service';

describe('ConnectedAircraftService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ConnectedAircraftService] });
	});

  it('should be created', () => {
    const service: ConnectedAircraftService = TestBed.get(ConnectedAircraftService);
    expect(service).toBeTruthy();
  });
});
