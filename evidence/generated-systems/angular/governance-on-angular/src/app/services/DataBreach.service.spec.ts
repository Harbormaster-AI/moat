import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DataBreachService } from './DataBreach.service';

describe('DataBreachService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DataBreachService] });
	});

  it('should be created', () => {
    const service: DataBreachService = TestBed.get(DataBreachService);
    expect(service).toBeTruthy();
  });
});
