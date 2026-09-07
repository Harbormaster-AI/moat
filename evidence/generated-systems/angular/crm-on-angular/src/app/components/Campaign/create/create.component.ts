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
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      budget: ['', Validators.required],
      actualCost: ['', Validators.required],
      expectedRevenue: ['', Validators.required],
      Organization: ['', ],
      ParentCampaign: ['', ],
      ChildCampaigns: ['', ],
      Members: ['', ],
      Opportunities: ['', ],
      Accounts: ['', ],
      Leads: ['', ],
      Contacts: ['', ],
      Teams: ['', ],
      Activities: ['', ],
      Status: ['', ],
      Type: ['', ]
        });
    }

    
    addCampaign(name, startDate, endDate, budget, actualCost, expectedRevenue, Organization, ParentCampaign, ChildCampaigns, Members, Opportunities, Accounts, Leads, Contacts, Teams, Activities, Status, Type): void {
        this.campaignService
        .addCampaign(name, startDate, endDate, budget, actualCost, expectedRevenue, Organization, ParentCampaign, ChildCampaigns, Members, Opportunities, Accounts, Leads, Contacts, Teams, Activities, Status, Type)
            .subscribe(() => {
                this.router.navigate(['/indexCampaign']);
            });
    }

    ngOnInit(): void {
    }
}