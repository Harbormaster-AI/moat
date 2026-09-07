import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InspectionCharacteristicService } from './InspectionCharacteristic.service';

describe('InspectionCharacteristicService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InspectionCharacteristicService] });
	});

  it('should be created', () => {
    const service: InspectionCharacteristicService = TestBed.get(InspectionCharacteristicService);
    expect(service).toBeTruthy();
  });
});
