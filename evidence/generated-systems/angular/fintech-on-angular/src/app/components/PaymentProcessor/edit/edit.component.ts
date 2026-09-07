import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PaymentProcessorService } from '../../../services/PaymentProcessor.service';
import { SubBaseComponent } from '../../PaymentProcessor/sub.base.component';


@Component({
    selector: 'app-edit-paymentProcessor',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPaymentProcessorComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PaymentProcessor';

    paymentProcessorForm: FormGroup;
    paymentProcessor: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PaymentProcessorService,
        private fb: FormBuilder
) {
        super(http);
        this.paymentProcessorForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      processorCode: ['', Validators.required],
      networkSupport: ['', Validators.required],
      Institutions: ['', ],
      Contracts: ['', ],
      Settlements: ['', ]
        });
    }

    
    updatePaymentProcessor(name, processorCode, networkSupport, Institutions, Contracts, Settlements): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePaymentProcessor(name, processorCode, networkSupport, Institutions, Contracts, Settlements, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPaymentProcessor']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPaymentProcessor(params['id']).subscribe(res => {
                this.paymentProcessor = res;
            });
        });
    }
}