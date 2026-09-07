import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PolicyAcknowledgementService } from '../../../services/PolicyAcknowledgement.service';
import { PolicyAcknowledgement } from '../../../models/PolicyAcknowledgement';
import { SubBaseComponent } from '../../PolicyAcknowledgement/sub.base.component';

@Component({
    selector: 'app-create-policyAcknowledgement',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePolicyAcknowledgementComponent extends SubBaseComponent implements OnInit {

    title = 'Add PolicyAcknowledgement';

    policyAcknowledgementForm: FormGroup;
    policyAcknowledgement: PolicyAcknowledgement;

    constructor( http: HttpClient,
        private policyAcknowledgementService: PolicyAcknowledgementService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPolicyAcknowledgement(acknowledgementDate, Policy, Employee, Status): void {
        this.policyAcknowledgementService
        .addPolicyAcknowledgement(acknowledgementDate, Policy, Employee, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPolicyAcknowledgement']);
            });
    }

    ngOnInit(): void {
    }
}