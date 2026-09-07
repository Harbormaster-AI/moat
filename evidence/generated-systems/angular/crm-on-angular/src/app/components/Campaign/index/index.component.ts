
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CampaignService } from '../../../services/Campaign.service';
import { Campaign } from '../../../models/Campaign';

@Component({
    selector: 'app-index-campaign',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCampaignComponent implements OnInit {

    campaigns: Campaign[] = [];

    constructor(
        private router: Router,
        private service: CampaignService
) {}

    ngOnInit(): void {
        this.getCampaigns();
}

    getCampaigns(): void {
        this.service.getCampaigns().subscribe((res) => {
        this.campaigns = res;
    });
}

    deleteCampaign(id: any): void {
        this.service.deleteCampaign(id)
            .subscribe(() => {
                this.getCampaigns();
            });
    }
}