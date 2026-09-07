import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PolicyAcknowledgementService } from '../../../services/PolicyAcknowledgement.service';
import { SubBaseComponent } from '../../PolicyAcknowledgement/sub.base.component';


@Component({
    selector: 'app-edit-policyAcknowledgement',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPolicyAcknowledgementComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PolicyAcknowledgement';

    policyAcknowledgementForm: FormGroup;
    policyAcknowledgement: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PolicyAcknowledgementService,
        private fb: FormBuilder
) {
        super(http);
        this.policyAcknowledgementForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  acknowledgementDate: ['', Validators.required],
      Policy: ['', ],
      Employee: ['', ],
      Status: ['', ]
        });
    }

    
    updatePolicyAcknowledgement(acknowledgementDate, Policy, Employee, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePolicyAcknowledgement(acknowledgementDate, Policy, Employee, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPolicyAcknowledgement']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPolicyAcknowledgement(params['id']).subscribe(res => {
                this.policyAcknowledgement = res;
            });
        });
    }
}