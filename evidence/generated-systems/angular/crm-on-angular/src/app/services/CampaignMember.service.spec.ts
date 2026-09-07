import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CampaignMemberService } from './CampaignMember.service';

describe('CampaignMemberService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CampaignMemberService] });
	});

  it('should be created', () => {
    const service: CampaignMemberService = TestBed.get(CampaignMemberService);
    expect(service).toBeTruthy();
  });
});
