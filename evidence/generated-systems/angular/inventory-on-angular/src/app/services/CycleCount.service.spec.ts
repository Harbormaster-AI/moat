import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CycleCountService } from './CycleCount.service';

describe('CycleCountService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CycleCountService] });
	});

  it('should be created', () => {
    const service: CycleCountService = TestBed.get(CycleCountService);
    expect(service).toBeTruthy();
  });
});
