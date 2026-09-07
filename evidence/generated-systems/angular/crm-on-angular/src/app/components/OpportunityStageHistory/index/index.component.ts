
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OpportunityStageHistoryService } from '../../../services/OpportunityStageHistory.service';
import { OpportunityStageHistory } from '../../../models/OpportunityStageHistory';

@Component({
    selector: 'app-index-opportunityStageHistory',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOpportunityStageHistoryComponent implements OnInit {

    opportunityStageHistorys: OpportunityStageHistory[] = [];

    constructor(
        private router: Router,
        private service: OpportunityStageHistoryService
) {}

    ngOnInit(): void {
        this.getOpportunityStageHistorys();
}

    getOpportunityStageHistorys(): void {
        this.service.getOpportunityStageHistorys().subscribe((res) => {
        this.opportunityStageHistorys = res;
    });
}

    deleteOpportunityStageHistory(id: any): void {
        this.service.deleteOpportunityStageHistory(id)
            .subscribe(() => {
                this.getOpportunityStageHistorys();
            });
    }
}