import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SalesCampaignService } from '../../../services/SalesCampaign.service';
import { SubBaseComponent } from '../../SalesCampaign/sub.base.component';


@Component({
    selector: 'app-edit-salesCampaign',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSalesCampaignComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SalesCampaign';

    salesCampaignForm: FormGroup;
    salesCampaign: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SalesCampaignService,
        private fb: FormBuilder
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

    
    updateSalesCampaign(campaignCode, Region, Operator, Quotes, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSalesCampaign(campaignCode, Region, Operator, Quotes, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSalesCampaign']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSalesCampaign(params['id']).subscribe(res => {
                this.salesCampaign = res;
            });
        });
    }
}