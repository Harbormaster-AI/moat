import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { QuoteLineItemService } from './QuoteLineItem.service';

describe('QuoteLineItemService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [QuoteLineItemService] });
	});

  it('should be created', () => {
    const service: QuoteLineItemService = TestBed.get(QuoteLineItemService);
    expect(service).toBeTruthy();
  });
});
