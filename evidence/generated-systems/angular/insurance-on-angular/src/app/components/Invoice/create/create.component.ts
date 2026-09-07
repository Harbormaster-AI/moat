import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InvoiceService } from '../../../services/Invoice.service';
import { Invoice } from '../../../models/Invoice';
import { SubBaseComponent } from '../../Invoice/sub.base.component';

@Component({
    selector: 'app-create-invoice',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInvoiceComponent extends SubBaseComponent implements OnInit {

    title = 'Add Invoice';

    invoiceForm: FormGroup;
    invoice: Invoice;

    constructor( http: HttpClient,
        private invoiceService: InvoiceService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.invoiceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  invoiceNumber: ['', Validators.required],
      dueDate: ['', Validators.required],
      totalDue: ['', Validators.required],
      BillingAccount: ['', ],
      Policy: ['', ],
      Payments: ['', ],
      Status: ['', ]
        });
    }

    
    addInvoice(invoiceNumber, dueDate, totalDue, BillingAccount, Policy, Payments, Status): void {
        this.invoiceService
        .addInvoice(invoiceNumber, dueDate, totalDue, BillingAccount, Policy, Payments, Status)
            .subscribe(() => {
                this.router.navigate(['/indexInvoice']);
            });
    }

    ngOnInit(): void {
    }
}