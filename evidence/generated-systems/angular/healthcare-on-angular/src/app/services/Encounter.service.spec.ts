import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { EncounterService } from './Encounter.service';

describe('EncounterService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [EncounterService] });
	});

  it('should be created', () => {
    const service: EncounterService = TestBed.get(EncounterService);
    expect(service).toBeTruthy();
  });
});
