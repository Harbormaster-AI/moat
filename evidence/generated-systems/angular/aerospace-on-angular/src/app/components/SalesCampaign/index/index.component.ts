
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SalesCampaignService } from '../../../services/SalesCampaign.service';
import { SalesCampaign } from '../../../models/SalesCampaign';

@Component({
    selector: 'app-index-salesCampaign',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSalesCampaignComponent implements OnInit {

    salesCampaigns: SalesCampaign[] = [];

    constructor(
        private router: Router,
        private service: SalesCampaignService
) {}

    ngOnInit(): void {
        this.getSalesCampaigns();
}

    getSalesCampaigns(): void {
        this.service.getSalesCampaigns().subscribe((res) => {
        this.salesCampaigns = res;
    });
}

    deleteSalesCampaign(id: any): void {
        this.service.deleteSalesCampaign(id)
            .subscribe(() => {
                this.getSalesCampaigns();
            });
    }
}