import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BusinessUnitService } from './BusinessUnit.service';

describe('BusinessUnitService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BusinessUnitService] });
	});

  it('should be created', () => {
    const service: BusinessUnitService = TestBed.get(BusinessUnitService);
    expect(service).toBeTruthy();
  });
});
