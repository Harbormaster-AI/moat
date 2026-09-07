import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PurchaseAgreementService } from '../../../services/PurchaseAgreement.service';
import { PurchaseAgreement } from '../../../models/PurchaseAgreement';
import { SubBaseComponent } from '../../PurchaseAgreement/sub.base.component';

@Component({
    selector: 'app-create-purchaseAgreement',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePurchaseAgreementComponent extends SubBaseComponent implements OnInit {

    title = 'Add PurchaseAgreement';

    purchaseAgreementForm: FormGroup;
    purchaseAgreement: PurchaseAgreement;

    constructor( http: HttpClient,
        private purchaseAgreementService: PurchaseAgreementService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.purchaseAgreementForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  agreementNumber: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      AircraftOrder: ['', ]
        });
    }

    
    addPurchaseAgreement(agreementNumber, effectiveDate, AircraftOrder): void {
        this.purchaseAgreementService
        .addPurchaseAgreement(agreementNumber, effectiveDate, AircraftOrder)
            .subscribe(() => {
                this.router.navigate(['/indexPurchaseAgreement']);
            });
    }

    ngOnInit(): void {
    }
}