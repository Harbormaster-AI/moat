import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CycleCountEntryService } from './CycleCountEntry.service';

describe('CycleCountEntryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CycleCountEntryService] });
	});

  it('should be created', () => {
    const service: CycleCountEntryService = TestBed.get(CycleCountEntryService);
    expect(service).toBeTruthy();
  });
});
