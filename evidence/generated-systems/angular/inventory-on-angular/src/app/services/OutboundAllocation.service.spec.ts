import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { OutboundAllocationService } from './OutboundAllocation.service';

describe('OutboundAllocationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [OutboundAllocationService] });
	});

  it('should be created', () => {
    const service: OutboundAllocationService = TestBed.get(OutboundAllocationService);
    expect(service).toBeTruthy();
  });
});
