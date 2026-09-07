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

    
    updateCampaign(name, startDate, endDate, budget, actualCost, expectedRevenue, Organization, ParentCampaign, ChildCampaigns, Members, Opportunities, Accounts, Leads, Contacts, Teams, Activities, Status, Type): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCampaign(name, startDate, endDate, budget, actualCost, expectedRevenue, Organization, ParentCampaign, ChildCampaigns, Members, Opportunities, Accounts, Leads, Contacts, Teams, Activities, Status, Type, params['id'])
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