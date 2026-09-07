
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CampaignMemberService } from '../../../services/CampaignMember.service';
import { CampaignMember } from '../../../models/CampaignMember';

@Component({
    selector: 'app-index-campaignMember',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCampaignMemberComponent implements OnInit {

    campaignMembers: CampaignMember[] = [];

    constructor(
        private router: Router,
        private service: CampaignMemberService
) {}

    ngOnInit(): void {
        this.getCampaignMembers();
}

    getCampaignMembers(): void {
        this.service.getCampaignMembers().subscribe((res) => {
        this.campaignMembers = res;
    });
}

    deleteCampaignMember(id: any): void {
        this.service.deleteCampaignMember(id)
            .subscribe(() => {
                this.getCampaignMembers();
            });
    }
}