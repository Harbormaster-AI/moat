import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CampaignMemberService } from '../../../services/CampaignMember.service';
import { SubBaseComponent } from '../../CampaignMember/sub.base.component';


@Component({
    selector: 'app-edit-campaignMember',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCampaignMemberComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CampaignMember';

    campaignMemberForm: FormGroup;
    campaignMember: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CampaignMemberService,
        private fb: FormBuilder
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

    
    updateCampaignMember(responded, Campaign, Lead, Contact, Status, MemberType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCampaignMember(responded, Campaign, Lead, Contact, Status, MemberType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCampaignMember']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCampaignMember(params['id']).subscribe(res => {
                this.campaignMember = res;
            });
        });
    }
}