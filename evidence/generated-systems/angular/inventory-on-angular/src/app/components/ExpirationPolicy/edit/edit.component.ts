import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ExpirationPolicyService } from '../../../services/ExpirationPolicy.service';
import { SubBaseComponent } from '../../ExpirationPolicy/sub.base.component';


@Component({
    selector: 'app-edit-expirationPolicy',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditExpirationPolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ExpirationPolicy';

    expirationPolicyForm: FormGroup;
    expirationPolicy: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ExpirationPolicyService,
        private fb: FormBuilder
) {
        super(http);
        this.expirationPolicyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  rejectIfDaysToExpireLessThan: ['', Validators.required],
      autoQuarantineDaysToExpire: ['', Validators.required],
      Sku: ['', ],
      Warehouse: ['', ],
      RotationMethod: ['', ]
        });
    }

    
    updateExpirationPolicy(rejectIfDaysToExpireLessThan, autoQuarantineDaysToExpire, Sku, Warehouse, RotationMethod): void {
        this.route.params.subscribe((params) => {

                        this.service.updateExpirationPolicy(rejectIfDaysToExpireLessThan, autoQuarantineDaysToExpire, Sku, Warehouse, RotationMethod, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexExpirationPolicy']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getExpirationPolicy(params['id']).subscribe(res => {
                this.expirationPolicy = res;
            });
        });
    }
}