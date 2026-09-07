import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { IncidentService } from './Incident.service';

describe('IncidentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [IncidentService] });
	});

  it('should be created', () => {
    const service: IncidentService = TestBed.get(IncidentService);
    expect(service).toBeTruthy();
  });
});
