import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SalesRegionService } from './SalesRegion.service';

describe('SalesRegionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SalesRegionService] });
	});

  it('should be created', () => {
    const service: SalesRegionService = TestBed.get(SalesRegionService);
    expect(service).toBeTruthy();
  });
});
