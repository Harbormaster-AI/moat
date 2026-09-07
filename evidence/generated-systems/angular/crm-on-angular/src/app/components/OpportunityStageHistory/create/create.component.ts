import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OpportunityStageHistoryService } from '../../../services/OpportunityStageHistory.service';
import { OpportunityStageHistory } from '../../../models/OpportunityStageHistory';
import { SubBaseComponent } from '../../OpportunityStageHistory/sub.base.component';

@Component({
    selector: 'app-create-opportunityStageHistory',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOpportunityStageHistoryComponent extends SubBaseComponent implements OnInit {

    title = 'Add OpportunityStageHistory';

    opportunityStageHistoryForm: FormGroup;
    opportunityStageHistory: OpportunityStageHistory;

    constructor( http: HttpClient,
        private opportunityStageHistoryService: OpportunityStageHistoryService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.opportunityStageHistoryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  changedAt: ['', Validators.required],
      comment: ['', Validators.required],
      Opportunity: ['', ],
      ChangedBy: ['', ],
      FromStage: ['', ],
      ToStage: ['', ]
        });
    }

    
    addOpportunityStageHistory(changedAt, comment, Opportunity, ChangedBy, FromStage, ToStage): void {
        this.opportunityStageHistoryService
        .addOpportunityStageHistory(changedAt, comment, Opportunity, ChangedBy, FromStage, ToStage)
            .subscribe(() => {
                this.router.navigate(['/indexOpportunityStageHistory']);
            });
    }

    ngOnInit(): void {
    }
}