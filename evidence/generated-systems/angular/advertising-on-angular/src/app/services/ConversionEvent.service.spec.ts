import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ConversionEventService } from './ConversionEvent.service';

describe('ConversionEventService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ConversionEventService] });
	});

  it('should be created', () => {
    const service: ConversionEventService = TestBed.get(ConversionEventService);
    expect(service).toBeTruthy();
  });
});
