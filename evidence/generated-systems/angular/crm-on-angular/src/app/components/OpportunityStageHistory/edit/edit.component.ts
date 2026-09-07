import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OpportunityStageHistoryService } from '../../../services/OpportunityStageHistory.service';
import { SubBaseComponent } from '../../OpportunityStageHistory/sub.base.component';


@Component({
    selector: 'app-edit-opportunityStageHistory',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOpportunityStageHistoryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit OpportunityStageHistory';

    opportunityStageHistoryForm: FormGroup;
    opportunityStageHistory: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OpportunityStageHistoryService,
        private fb: FormBuilder
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

    
    updateOpportunityStageHistory(changedAt, comment, Opportunity, ChangedBy, FromStage, ToStage): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOpportunityStageHistory(changedAt, comment, Opportunity, ChangedBy, FromStage, ToStage, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOpportunityStageHistory']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOpportunityStageHistory(params['id']).subscribe(res => {
                this.opportunityStageHistory = res;
            });
        });
    }
}