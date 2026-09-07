import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LotService } from './Lot.service';

describe('LotService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LotService] });
	});

  it('should be created', () => {
    const service: LotService = TestBed.get(LotService);
    expect(service).toBeTruthy();
  });
});
