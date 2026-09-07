import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DemandSignalService } from './DemandSignal.service';

describe('DemandSignalService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DemandSignalService] });
	});

  it('should be created', () => {
    const service: DemandSignalService = TestBed.get(DemandSignalService);
    expect(service).toBeTruthy();
  });
});
