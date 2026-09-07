import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MatterService } from './Matter.service';

describe('MatterService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MatterService] });
	});

  it('should be created', () => {
    const service: MatterService = TestBed.get(MatterService);
    expect(service).toBeTruthy();
  });
});
