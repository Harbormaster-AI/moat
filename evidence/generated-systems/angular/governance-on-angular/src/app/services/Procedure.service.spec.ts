import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ProcedureService } from './Procedure.service';

describe('ProcedureService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ProcedureService] });
	});

  it('should be created', () => {
    const service: ProcedureService = TestBed.get(ProcedureService);
    expect(service).toBeTruthy();
  });
});
