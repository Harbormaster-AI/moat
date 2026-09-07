import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DistributorService } from './Distributor.service';

describe('DistributorService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DistributorService] });
	});

  it('should be created', () => {
    const service: DistributorService = TestBed.get(DistributorService);
    expect(service).toBeTruthy();
  });
});
