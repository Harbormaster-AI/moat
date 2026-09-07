import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SalesCampaignService } from '../../../services/SalesCampaign.service';
import { SalesCampaign } from '../../../models/SalesCampaign';
import { SubBaseComponent } from '../../SalesCampaign/sub.base.component';

@Component({
    selector: 'app-create-salesCampaign',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSalesCampaignComponent extends SubBaseComponent implements OnInit {

    title = 'Add SalesCampaign';

    salesCampaignForm: FormGroup;
    salesCampaign: SalesCampaign;

    constructor( http: HttpClient,
        private salesCampaignService: SalesCampaignService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.salesCampaignForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  campaignCode: ['', Validators.required],
      Region: ['', ],
      Operator: ['', ],
      Quotes: ['', ],
      Status: ['', ]
        });
    }

    
    addSalesCampaign(campaignCode, Region, Operator, Quotes, Status): void {
        this.salesCampaignService
        .addSalesCampaign(campaignCode, Region, Operator, Quotes, Status)
            .subscribe(() => {
                this.router.navigate(['/indexSalesCampaign']);
            });
    }

    ngOnInit(): void {
    }
}