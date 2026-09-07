import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OpportunityService } from '../../../services/Opportunity.service';
import { SubBaseComponent } from '../../Opportunity/sub.base.component';


@Component({
    selector: 'app-edit-opportunity',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOpportunityComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Opportunity';

    opportunityForm: FormGroup;
    opportunity: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OpportunityService,
        private fb: FormBuilder
) {
        super(http);
        this.opportunityForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      amount: ['', Validators.required],
      closeDate: ['', Validators.required],
      probability: ['', Validators.required],
      description: ['', Validators.required],
      Organization: ['', ],
      Account: ['', ],
      Owner: ['', ],
      Contacts: ['', ],
      LineItems: ['', ],
      StageHistory: ['', ],
      Quotes: ['', ],
      Orders: ['', ],
      Campaigns: ['', ],
      Activities: ['', ],
      Teams: ['', ],
      Stage: ['', ],
      Type: ['', ],
      ForecastCategory: ['', ]
        });
    }

    
    updateOpportunity(name, amount, closeDate, probability, description, Organization, Account, Owner, Contacts, LineItems, StageHistory, Quotes, Orders, Campaigns, Activities, Teams, Stage, Type, ForecastCategory): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOpportunity(name, amount, closeDate, probability, description, Organization, Account, Owner, Contacts, LineItems, StageHistory, Quotes, Orders, Campaigns, Activities, Teams, Stage, Type, ForecastCategory, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOpportunity']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOpportunity(params['id']).subscribe(res => {
                this.opportunity = res;
            });
        });
    }
}