import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CampaignMemberService } from '../../../services/CampaignMember.service';
import { CampaignMember } from '../../../models/CampaignMember';
import { SubBaseComponent } from '../../CampaignMember/sub.base.component';

@Component({
    selector: 'app-create-campaignMember',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCampaignMemberComponent extends SubBaseComponent implements OnInit {

    title = 'Add CampaignMember';

    campaignMemberForm: FormGroup;
    campaignMember: CampaignMember;

    constructor( http: HttpClient,
        private campaignMemberService: CampaignMemberService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.campaignMemberForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  responded: ['', Validators.required],
      Campaign: ['', ],
      Lead: ['', ],
      Contact: ['', ],
      Status: ['', ],
      MemberType: ['', ]
        });
    }

    
    addCampaignMember(responded, Campaign, Lead, Contact, Status, MemberType): void {
        this.campaignMemberService
        .addCampaignMember(responded, Campaign, Lead, Contact, Status, MemberType)
            .subscribe(() => {
                this.router.navigate(['/indexCampaignMember']);
            });
    }

    ngOnInit(): void {
    }
}