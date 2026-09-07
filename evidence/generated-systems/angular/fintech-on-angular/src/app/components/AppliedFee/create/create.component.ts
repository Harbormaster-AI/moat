import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AppliedFeeService } from '../../../services/AppliedFee.service';
import { AppliedFee } from '../../../models/AppliedFee';
import { SubBaseComponent } from '../../AppliedFee/sub.base.component';

@Component({
    selector: 'app-create-appliedFee',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAppliedFeeComponent extends SubBaseComponent implements OnInit {

    title = 'Add AppliedFee';

    appliedFeeForm: FormGroup;
    appliedFee: AppliedFee;

    constructor( http: HttpClient,
        private appliedFeeService: AppliedFeeService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.appliedFeeForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  amount: ['', Validators.required],
      description: ['', Validators.required],
      PaymentOrder: ['', ],
      Transaction: ['', ],
      FeeType: ['', ]
        });
    }

    
    addAppliedFee(amount, description, PaymentOrder, Transaction, FeeType): void {
        this.appliedFeeService
        .addAppliedFee(amount, description, PaymentOrder, Transaction, FeeType)
            .subscribe(() => {
                this.router.navigate(['/indexAppliedFee']);
            });
    }

    ngOnInit(): void {
    }
}