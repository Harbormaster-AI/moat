import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OpportunityService } from '../../../services/Opportunity.service';
import { Opportunity } from '../../../models/Opportunity';
import { SubBaseComponent } from '../../Opportunity/sub.base.component';

@Component({
    selector: 'app-create-opportunity',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOpportunityComponent extends SubBaseComponent implements OnInit {

    title = 'Add Opportunity';

    opportunityForm: FormGroup;
    opportunity: Opportunity;

    constructor( http: HttpClient,
        private opportunityService: OpportunityService,
        private fb: FormBuilder,
        private router: Router
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

    
    addOpportunity(name, amount, closeDate, probability, description, Organization, Account, Owner, Contacts, LineItems, StageHistory, Quotes, Orders, Campaigns, Activities, Teams, Stage, Type, ForecastCategory): void {
        this.opportunityService
        .addOpportunity(name, amount, closeDate, probability, description, Organization, Account, Owner, Contacts, LineItems, StageHistory, Quotes, Orders, Campaigns, Activities, Teams, Stage, Type, ForecastCategory)
            .subscribe(() => {
                this.router.navigate(['/indexOpportunity']);
            });
    }

    ngOnInit(): void {
    }
}