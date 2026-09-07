import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PolicyService } from '../../../services/Policy.service';
import { SubBaseComponent } from '../../Policy/sub.base.component';


@Component({
    selector: 'app-edit-policy',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Policy';

    policyForm: FormGroup;
    policy: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PolicyService,
        private fb: FormBuilder
) {
        super(http);
        this.policyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      versionLabel: ['', Validators.required],
      approvalDate: ['', Validators.required],
      nextReviewDate: ['', Validators.required],
      documentUrl: ['', Validators.required],
      Organization: ['', ],
      Owners: ['', ],
      RelatedRequirements: ['', ],
      Controls: ['', ],
      Procedures: ['', ],
      Exceptions: ['', ],
      Attestations: ['', ],
      PolicyType: ['', ],
      Status: ['', ]
        });
    }

    
    updatePolicy(title, versionLabel, approvalDate, nextReviewDate, documentUrl, Organization, Owners, RelatedRequirements, Controls, Procedures, Exceptions, Attestations, PolicyType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePolicy(title, versionLabel, approvalDate, nextReviewDate, documentUrl, Organization, Owners, RelatedRequirements, Controls, Procedures, Exceptions, Attestations, PolicyType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPolicy']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPolicy(params['id']).subscribe(res => {
                this.policy = res;
            });
        });
    }
}