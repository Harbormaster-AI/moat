import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InspectionLotService } from './InspectionLot.service';

describe('InspectionLotService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InspectionLotService] });
	});

  it('should be created', () => {
    const service: InspectionLotService = TestBed.get(InspectionLotService);
    expect(service).toBeTruthy();
  });
});
