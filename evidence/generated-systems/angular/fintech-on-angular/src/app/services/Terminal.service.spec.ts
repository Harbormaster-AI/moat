import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TerminalService } from './Terminal.service';

describe('TerminalService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TerminalService] });
	});

  it('should be created', () => {
    const service: TerminalService = TestBed.get(TerminalService);
    expect(service).toBeTruthy();
  });
});
