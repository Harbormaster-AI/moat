import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AppliedFeeService } from '../../../services/AppliedFee.service';
import { SubBaseComponent } from '../../AppliedFee/sub.base.component';


@Component({
    selector: 'app-edit-appliedFee',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAppliedFeeComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AppliedFee';

    appliedFeeForm: FormGroup;
    appliedFee: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AppliedFeeService,
        private fb: FormBuilder
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

    
    updateAppliedFee(amount, description, PaymentOrder, Transaction, FeeType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAppliedFee(amount, description, PaymentOrder, Transaction, FeeType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAppliedFee']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAppliedFee(params['id']).subscribe(res => {
                this.appliedFee = res;
            });
        });
    }
}