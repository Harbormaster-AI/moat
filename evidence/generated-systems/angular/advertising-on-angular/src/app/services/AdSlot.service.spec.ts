import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AdSlotService } from './AdSlot.service';

describe('AdSlotService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AdSlotService] });
	});

  it('should be created', () => {
    const service: AdSlotService = TestBed.get(AdSlotService);
    expect(service).toBeTruthy();
  });
});
