import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CampaignService } from '../../../services/Campaign.service';
import { Campaign } from '../../../models/Campaign';
import { SubBaseComponent } from '../../Campaign/sub.base.component';

@Component({
    selector: 'app-create-campaign',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCampaignComponent extends SubBaseComponent implements OnInit {

    title = 'Add Campaign';

    campaignForm: FormGroup;
    campaign: Campaign;

    constructor( http: HttpClient,
        private campaignService: CampaignService,
        private fb: FormBuilder,
        private router: Router
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

    
    addCampaign(name, totalBudget, flight, AdAccount, LineItems, Kpis, TrackingPixels, Audiences, Reports, InsertionOrder, Objective, Status): void {
        this.campaignService
        .addCampaign(name, totalBudget, flight, AdAccount, LineItems, Kpis, TrackingPixels, Audiences, Reports, InsertionOrder, Objective, Status)
            .subscribe(() => {
                this.router.navigate(['/indexCampaign']);
            });
    }

    ngOnInit(): void {
    }
}