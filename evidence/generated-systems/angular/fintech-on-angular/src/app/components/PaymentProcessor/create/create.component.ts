import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PaymentProcessorService } from '../../../services/PaymentProcessor.service';
import { PaymentProcessor } from '../../../models/PaymentProcessor';
import { SubBaseComponent } from '../../PaymentProcessor/sub.base.component';

@Component({
    selector: 'app-create-paymentProcessor',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePaymentProcessorComponent extends SubBaseComponent implements OnInit {

    title = 'Add PaymentProcessor';

    paymentProcessorForm: FormGroup;
    paymentProcessor: PaymentProcessor;

    constructor( http: HttpClient,
        private paymentProcessorService: PaymentProcessorService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPaymentProcessor(name, processorCode, networkSupport, Institutions, Contracts, Settlements): void {
        this.paymentProcessorService
        .addPaymentProcessor(name, processorCode, networkSupport, Institutions, Contracts, Settlements)
            .subscribe(() => {
                this.router.navigate(['/indexPaymentProcessor']);
            });
    }

    ngOnInit(): void {
    }
}