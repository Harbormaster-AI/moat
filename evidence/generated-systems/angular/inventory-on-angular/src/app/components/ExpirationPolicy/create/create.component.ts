import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ExpirationPolicyService } from '../../../services/ExpirationPolicy.service';
import { ExpirationPolicy } from '../../../models/ExpirationPolicy';
import { SubBaseComponent } from '../../ExpirationPolicy/sub.base.component';

@Component({
    selector: 'app-create-expirationPolicy',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateExpirationPolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Add ExpirationPolicy';

    expirationPolicyForm: FormGroup;
    expirationPolicy: ExpirationPolicy;

    constructor( http: HttpClient,
        private expirationPolicyService: ExpirationPolicyService,
        private fb: FormBuilder,
        private router: Router
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

    
    addExpirationPolicy(rejectIfDaysToExpireLessThan, autoQuarantineDaysToExpire, Sku, Warehouse, RotationMethod): void {
        this.expirationPolicyService
        .addExpirationPolicy(rejectIfDaysToExpireLessThan, autoQuarantineDaysToExpire, Sku, Warehouse, RotationMethod)
            .subscribe(() => {
                this.router.navigate(['/indexExpirationPolicy']);
            });
    }

    ngOnInit(): void {
    }
}