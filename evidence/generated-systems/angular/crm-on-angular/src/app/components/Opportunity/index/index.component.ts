
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OpportunityService } from '../../../services/Opportunity.service';
import { Opportunity } from '../../../models/Opportunity';

@Component({
    selector: 'app-index-opportunity',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOpportunityComponent implements OnInit {

    opportunitys: Opportunity[] = [];

    constructor(
        private router: Router,
        private service: OpportunityService
) {}

    ngOnInit(): void {
        this.getOpportunitys();
}

    getOpportunitys(): void {
        this.service.getOpportunitys().subscribe((res) => {
        this.opportunitys = res;
    });
}

    deleteOpportunity(id: any): void {
        this.service.deleteOpportunity(id)
            .subscribe(() => {
                this.getOpportunitys();
            });
    }
}