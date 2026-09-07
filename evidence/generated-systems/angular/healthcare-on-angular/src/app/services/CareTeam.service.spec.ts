import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CareTeamService } from './CareTeam.service';

describe('CareTeamService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CareTeamService] });
	});

  it('should be created', () => {
    const service: CareTeamService = TestBed.get(CareTeamService);
    expect(service).toBeTruthy();
  });
});
