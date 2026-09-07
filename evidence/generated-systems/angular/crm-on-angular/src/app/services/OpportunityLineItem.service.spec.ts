import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { OpportunityLineItemService } from './OpportunityLineItem.service';

describe('OpportunityLineItemService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [OpportunityLineItemService] });
	});

  it('should be created', () => {
    const service: OpportunityLineItemService = TestBed.get(OpportunityLineItemService);
    expect(service).toBeTruthy();
  });
});
