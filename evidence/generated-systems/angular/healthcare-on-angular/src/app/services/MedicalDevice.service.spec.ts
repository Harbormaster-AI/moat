import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MedicalDeviceService } from './MedicalDevice.service';

describe('MedicalDeviceService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MedicalDeviceService] });
	});

  it('should be created', () => {
    const service: MedicalDeviceService = TestBed.get(MedicalDeviceService);
    expect(service).toBeTruthy();
  });
});
