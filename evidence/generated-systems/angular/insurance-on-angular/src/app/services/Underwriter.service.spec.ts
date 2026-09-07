import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { UnderwriterService } from './Underwriter.service';

describe('UnderwriterService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [UnderwriterService] });
	});

  it('should be created', () => {
    const service: UnderwriterService = TestBed.get(UnderwriterService);
    expect(service).toBeTruthy();
  });
});
