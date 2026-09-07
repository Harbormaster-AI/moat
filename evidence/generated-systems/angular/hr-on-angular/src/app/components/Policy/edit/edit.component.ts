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
                  policyNumber: ['', Validators.required],
      name: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      description: ['', Validators.required],
      Organization: ['', ],
      Acknowledgements: ['', ]
        });
    }

    
    updatePolicy(policyNumber, name, effectiveDate, description, Organization, Acknowledgements): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePolicy(policyNumber, name, effectiveDate, description, Organization, Acknowledgements, params['id'])
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