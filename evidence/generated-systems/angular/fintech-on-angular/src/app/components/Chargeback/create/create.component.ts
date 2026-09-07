import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ChargebackService } from '../../../services/Chargeback.service';
import { Chargeback } from '../../../models/Chargeback';
import { SubBaseComponent } from '../../Chargeback/sub.base.component';

@Component({
    selector: 'app-create-chargeback',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateChargebackComponent extends SubBaseComponent implements OnInit {

    title = 'Add Chargeback';

    chargebackForm: FormGroup;
    chargeback: Chargeback;

    constructor( http: HttpClient,
        private chargebackService: ChargebackService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.chargebackForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  chargebackReference: ['', Validators.required],
      amount: ['', Validators.required],
      postedAt: ['', Validators.required],
      Dispute: ['', ],
      Transaction: ['', ],
      Stage: ['', ],
      Status: ['', ]
        });
    }

    
    addChargeback(chargebackReference, amount, postedAt, Dispute, Transaction, Stage, Status): void {
        this.chargebackService
        .addChargeback(chargebackReference, amount, postedAt, Dispute, Transaction, Stage, Status)
            .subscribe(() => {
                this.router.navigate(['/indexChargeback']);
            });
    }

    ngOnInit(): void {
    }
}