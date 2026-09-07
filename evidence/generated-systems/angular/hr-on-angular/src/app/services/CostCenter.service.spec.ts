import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CostCenterService } from './CostCenter.service';

describe('CostCenterService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CostCenterService] });
	});

  it('should be created', () => {
    const service: CostCenterService = TestBed.get(CostCenterService);
    expect(service).toBeTruthy();
  });
});
