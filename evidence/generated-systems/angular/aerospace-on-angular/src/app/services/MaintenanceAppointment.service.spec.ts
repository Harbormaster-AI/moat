import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MaintenanceAppointmentService } from './MaintenanceAppointment.service';

describe('MaintenanceAppointmentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MaintenanceAppointmentService] });
	});

  it('should be created', () => {
    const service: MaintenanceAppointmentService = TestBed.get(MaintenanceAppointmentService);
    expect(service).toBeTruthy();
  });
});
