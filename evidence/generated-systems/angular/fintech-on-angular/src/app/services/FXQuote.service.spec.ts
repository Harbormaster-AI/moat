import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { FXQuoteService } from './FXQuote.service';

describe('FXQuoteService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [FXQuoteService] });
	});

  it('should be created', () => {
    const service: FXQuoteService = TestBed.get(FXQuoteService);
    expect(service).toBeTruthy();
  });
});
