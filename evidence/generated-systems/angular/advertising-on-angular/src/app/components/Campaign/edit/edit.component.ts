import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CampaignService } from '../../../services/Campaign.service';
import { SubBaseComponent } from '../../Campaign/sub.base.component';


@Component({
    selector: 'app-edit-campaign',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCampaignComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Campaign';

    campaignForm: FormGroup;
    campaign: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CampaignService,
        private fb: FormBuilder
) {
        super(http);
        this.campaignForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      totalBudget: ['', Validators.required],
      flight: ['', Validators.required],
      AdAccount: ['', ],
      LineItems: ['', ],
      Kpis: ['', ],
      TrackingPixels: ['', ],
      Audiences: ['', ],
      Reports: ['', ],
      InsertionOrder: ['', ],
      Objective: ['', ],
      Status: ['', ]
        });
    }

    
    updateCampaign(name, totalBudget, flight, AdAccount, LineItems, Kpis, TrackingPixels, Audiences, Reports, InsertionOrder, Objective, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCampaign(name, totalBudget, flight, AdAccount, LineItems, Kpis, TrackingPixels, Audiences, Reports, InsertionOrder, Objective, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCampaign']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCampaign(params['id']).subscribe(res => {
                this.campaign = res;
            });
        });
    }
}