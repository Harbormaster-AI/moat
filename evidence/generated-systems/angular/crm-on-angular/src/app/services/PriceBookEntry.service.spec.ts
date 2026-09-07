import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PriceBookEntryService } from './PriceBookEntry.service';

describe('PriceBookEntryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PriceBookEntryService] });
	});

  it('should be created', () => {
    const service: PriceBookEntryService = TestBed.get(PriceBookEntryService);
    expect(service).toBeTruthy();
  });
});
