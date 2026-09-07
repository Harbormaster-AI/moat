
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OpportunityLineItemService } from '../../../services/OpportunityLineItem.service';
import { OpportunityLineItem } from '../../../models/OpportunityLineItem';

@Component({
    selector: 'app-index-opportunityLineItem',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOpportunityLineItemComponent implements OnInit {

    opportunityLineItems: OpportunityLineItem[] = [];

    constructor(
        private router: Router,
        private service: OpportunityLineItemService
) {}

    ngOnInit(): void {
        this.getOpportunityLineItems();
}

    getOpportunityLineItems(): void {
        this.service.getOpportunityLineItems().subscribe((res) => {
        this.opportunityLineItems = res;
    });
}

    deleteOpportunityLineItem(id: any): void {
        this.service.deleteOpportunityLineItem(id)
            .subscribe(() => {
                this.getOpportunityLineItems();
            });
    }
}