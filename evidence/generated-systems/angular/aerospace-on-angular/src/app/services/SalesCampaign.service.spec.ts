import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SalesCampaignService } from './SalesCampaign.service';

describe('SalesCampaignService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SalesCampaignService] });
	});

  it('should be created', () => {
    const service: SalesCampaignService = TestBed.get(SalesCampaignService);
    expect(service).toBeTruthy();
  });
});
